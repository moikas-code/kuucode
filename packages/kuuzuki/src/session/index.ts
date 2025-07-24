import path from "path"
import { Decimal } from "decimal.js"
import { z, ZodSchema } from "zod"
import {
  generateText,
  LoadAPIKeyError,
  streamText,
  tool,
  wrapLanguageModel,
  type Tool as AITool,
  type LanguageModelUsage,
  type ProviderMetadata,
  type CoreMessage,
  // stepCountIs,
  type StreamTextResult,
} from "ai"

import PROMPT_INITIALIZE from "../session/prompt/initialize.txt"
import PROMPT_PLAN from "../session/prompt/plan.txt"
import PROMPT_ANTHROPIC_SPOOF from "../session/prompt/anthropic_spoof.txt"

import { App } from "../app/app"
import { Bus } from "../bus"
import { Config } from "../config/config"
import { Flag } from "../flag/flag"
import { Identifier } from "../id/id"
import { Installation } from "../installation"
import { MCP } from "../mcp"
import { Provider } from "../provider/provider"
import { ProviderTransform } from "../provider/transform"
import type { ModelsDev } from "../provider/models"
import { Share } from "../share/share"
import { Snapshot } from "../snapshot"
import { Storage } from "../storage/storage"
import { Log } from "../util/log"
import { NamedError } from "../util/error"
import { SystemPrompt } from "./system"
import { FileTime } from "../file/time"
import { MessageV2 } from "./message-v2"
import { Mode } from "./mode"
import { LSP } from "../lsp"
import { ReadTool } from "../tool/read"
import { wrapToolsWithFallback } from "../tool/fallback"

export namespace Session {
  const log = Log.create({ service: "session" })

  const OUTPUT_TOKEN_MAX = 32_000
  const HARD_CONTEXT_LIMIT = 180_000 // Conservative but reasonable limit

  export const Info = z
    .object({
      id: Identifier.schema("session"),
      parentID: Identifier.schema("session").optional(),
      share: z
        .object({
          url: z.string(),
        })
        .optional(),
      title: z.string(),
      version: z.string(),
      time: z.object({
        created: z.number(),
        updated: z.number(),
      }),
      revert: z
        .object({
          messageID: z.string(),
          part: z.number(),
          snapshot: z.string().optional(),
        })
        .optional(),
    })
    .openapi({
      ref: "Session",
    })
  export type Info = z.output<typeof Info>

  export const ShareInfo = z
    .object({
      secret: z.string(),
      url: z.string(),
    })
    .openapi({
      ref: "SessionShare",
    })
  export type ShareInfo = z.output<typeof ShareInfo>

  export const Event = {
    Updated: Bus.event(
      "session.updated",
      z.object({
        info: Info,
      }),
    ),
    Deleted: Bus.event(
      "session.deleted",
      z.object({
        info: Info,
      }),
    ),
    Idle: Bus.event(
      "session.idle",
      z.object({
        sessionID: z.string(),
      }),
    ),
    Error: Bus.event(
      "session.error",
      z.object({
        sessionID: z.string().optional(),
        error: MessageV2.Assistant.shape.error,
      }),
    ),
  }

  const state = App.state(
    "session",
    () => {
      const sessions = new Map<string, Info>()
      const messages = new Map<string, MessageV2.Info[]>()
      const pending = new Map<string, AbortController>()
      const queued = new Map<
        string,
        {
          input: ChatInput
          message: MessageV2.User
          parts: MessageV2.Part[]
          processed: boolean
          callback: (input: { info: MessageV2.Assistant; parts: MessageV2.Part[] }) => void
        }[]
      >()

      return {
        sessions,
        messages,
        pending,
        queued,
      }
    },
    async (state) => {
      for (const [_, controller] of state.pending) {
        controller.abort()
      }
    },
  )

  export async function create(parentID?: string) {
    const result: Info = {
      id: Identifier.descending("session"),
      version: Installation.VERSION,
      parentID,
      title: (parentID ? "Child session - " : "New Session - ") + new Date().toISOString(),
      time: {
        created: Date.now(),
        updated: Date.now(),
      },
    }
    log.info("created", result)
    state().sessions.set(result.id, result)
    await Storage.writeJSON("session/info/" + result.id, result)
    const cfg = await Config.get()
    if (!result.parentID && (Flag.KUUZUKI_AUTO_SHARE || cfg.share === "auto"))
      share(result.id)
        .then((share) => {
          update(result.id, (draft) => {
            draft.share = share
          })
        })
        .catch(() => {
          // Silently ignore sharing errors during session creation
        })
    Bus.publish(Event.Updated, {
      info: result,
    })
    return result
  }

  export async function get(id: string) {
    const result = state().sessions.get(id)
    if (result) {
      return result
    }
    const read = await Storage.readJSON<Info>("session/info/" + id)
    state().sessions.set(id, read)
    return read as Info
  }

  export async function getShare(id: string) {
    return Storage.readJSON<ShareInfo>("session/share/" + id)
  }

  export async function share(id: string) {
    const cfg = await Config.get()
    if (cfg.share === "disabled") {
      throw new Error("Sharing is disabled in configuration")
    }

    const session = await get(id)
    if (session.share) return session.share
    const share = await Share.create(id)
    await update(id, (draft) => {
      draft.share = {
        url: share.url,
      }
    })
    await Storage.writeJSON<ShareInfo>("session/share/" + id, share)
    await Share.sync("session/info/" + id, session)
    for (const msg of await messages(id)) {
      await Share.sync("session/message/" + id + "/" + msg.info.id, msg.info)
      for (const part of msg.parts) {
        await Share.sync("session/part/" + id + "/" + msg.info.id + "/" + part.id, part)
      }
    }
    return share
  }

  export async function unshare(id: string) {
    const share = await getShare(id)
    if (!share) return
    await Storage.remove("session/share/" + id)
    await update(id, (draft) => {
      draft.share = undefined
    })
    await Share.remove(id, share.secret)
  }

  export async function update(id: string, editor: (session: Info) => void) {
    const { sessions } = state()
    const session = await get(id)
    if (!session) return
    editor(session)
    session.time.updated = Date.now()
    sessions.set(id, session)
    await Storage.writeJSON("session/info/" + id, session)
    Bus.publish(Event.Updated, {
      info: session,
    })
    return session
  }

  export async function messages(sessionID: string) {
    const result = [] as {
      info: MessageV2.Info
      parts: MessageV2.Part[]
    }[]
    for (const p of await Storage.list("session/message/" + sessionID)) {
      const read = await Storage.readJSON<MessageV2.Info>(p)
      result.push({
        info: read,
        parts: await parts(sessionID, read.id),
      })
    }
    result.sort((a, b) => (a.info.id > b.info.id ? 1 : -1))
    return result
  }

  export async function getMessage(sessionID: string, messageID: string) {
    return Storage.readJSON<MessageV2.Info>("session/message/" + sessionID + "/" + messageID)
  }

  export async function parts(sessionID: string, messageID: string) {
    const result = [] as MessageV2.Part[]
    for (const item of await Storage.list("session/part/" + sessionID + "/" + messageID)) {
      const read = await Storage.readJSON<MessageV2.Part>(item)
      result.push(read)
    }
    result.sort((a, b) => (a.id > b.id ? 1 : -1))
    return result
  }

  export async function* list() {
    for (const item of await Storage.list("session/info")) {
      const sessionID = path.basename(item, ".json")
      yield get(sessionID)
    }
  }

  export async function children(parentID: string) {
    const result = [] as Session.Info[]
    for (const item of await Storage.list("session/info")) {
      const sessionID = path.basename(item, ".json")
      const session = await get(sessionID)
      if (session.parentID !== parentID) continue
      result.push(session)
    }
    return result
  }

  export function abort(sessionID: string) {
    const controller = state().pending.get(sessionID)
    if (!controller) return false
    controller.abort()
    state().pending.delete(sessionID)
    return true
  }

  export async function remove(sessionID: string, emitEvent = true) {
    try {
      abort(sessionID)
      const session = await get(sessionID)
      for (const child of await children(sessionID)) {
        await remove(child.id, false)
      }
      await unshare(sessionID).catch(() => {})
      await Storage.remove(`session/info/${sessionID}`).catch(() => {})
      await Storage.removeDir(`session/message/${sessionID}/`).catch(() => {})
      state().sessions.delete(sessionID)
      state().messages.delete(sessionID)
      if (emitEvent) {
        Bus.publish(Event.Deleted, {
          info: session,
        })
      }
    } catch (e) {
      log.error(e)
    }
  }

  async function updateMessage(msg: MessageV2.Info) {
    await Storage.writeJSON("session/message/" + msg.sessionID + "/" + msg.id, msg)
    Bus.publish(MessageV2.Event.Updated, {
      info: msg,
    })
  }

  async function updatePart(part: MessageV2.Part) {
    await Storage.writeJSON(["session", "part", part.sessionID, part.messageID, part.id].join("/"), part)
    Bus.publish(MessageV2.Event.PartUpdated, {
      part,
    })
    return part
  }

  export const ChatInput = z.object({
    sessionID: Identifier.schema("session"),
    messageID: Identifier.schema("message").optional(),
    providerID: z.string(),
    modelID: z.string(),
    mode: z.string().optional(),
    tools: z.record(z.boolean()).optional(),
    parts: z.array(
      z.discriminatedUnion("type", [
        MessageV2.TextPart.omit({
          messageID: true,
          sessionID: true,
        })
          .partial({
            id: true,
          })
          .openapi({
            ref: "TextPartInput",
          }),
        MessageV2.FilePart.omit({
          messageID: true,
          sessionID: true,
        })
          .partial({
            id: true,
          })
          .openapi({
            ref: "FilePartInput",
          }),
      ]),
    ),
  })
  export type ChatInput = z.infer<typeof ChatInput>

  export async function chat(
    input: z.infer<typeof ChatInput>,
    retryCount = 0,
  ): Promise<{ info: MessageV2.Assistant; parts: MessageV2.Part[] }> {
    const l = log.clone().tag("session", input.sessionID)
    l.info("chatting")

    const userMsg: MessageV2.Info = {
      id: input.messageID ?? Identifier.ascending("message"),
      role: "user",
      sessionID: input.sessionID,
      time: {
        created: Date.now(),
      },
    }

    const app = App.info()
    const userParts = await Promise.all(
      input.parts.map(async (part): Promise<MessageV2.Part[]> => {
        if (part.type === "file") {
          const url = new URL(part.url)
          switch (url.protocol) {
            case "file:":
              // have to normalize, symbol search returns absolute paths
              // Decode the pathname since URL constructor doesn't automatically decode it
              const pathname = decodeURIComponent(url.pathname)
              const relativePath = pathname.replace(app.path.cwd, ".")
              const filePath = path.join(app.path.cwd, relativePath)

              if (part.mime === "text/plain") {
                let offset: number | undefined = undefined
                let limit: number | undefined = undefined
                const range = {
                  start: url.searchParams.get("start"),
                  end: url.searchParams.get("end"),
                }
                if (range.start != null) {
                  const filePath = part.url.split("?")[0]
                  let start = parseInt(range.start)
                  let end = range.end ? parseInt(range.end) : undefined
                  // some LSP servers (eg, gopls) don't give full range in
                  // workspace/symbol searches, so we'll try to find the
                  // symbol in the document to get the full range
                  if (start === end) {
                    const symbols = await LSP.documentSymbol(filePath)
                    for (const symbol of symbols) {
                      let range: LSP.Range | undefined
                      if ("range" in symbol) {
                        range = symbol.range
                      } else if ("location" in symbol) {
                        range = symbol.location.range
                      }
                      if (range?.start?.line && range?.start?.line === start) {
                        start = range.start.line
                        end = range?.end?.line ?? start
                        break
                      }
                    }
                    offset = Math.max(start - 2, 0)
                    if (end) {
                      limit = end - offset + 2
                    }
                  }
                }
                const args = { filePath, offset, limit }
                const result = await ReadTool.execute(args, {
                  sessionID: input.sessionID,
                  abort: new AbortController().signal,
                  messageID: userMsg.id,
                  metadata: async () => {},
                })
                return [
                  {
                    id: Identifier.ascending("part"),
                    messageID: userMsg.id,
                    sessionID: input.sessionID,
                    type: "text",
                    synthetic: true,
                    text: `Called the Read tool with the following input: ${JSON.stringify(args)}`,
                  },
                  {
                    id: Identifier.ascending("part"),
                    messageID: userMsg.id,
                    sessionID: input.sessionID,
                    type: "text",
                    synthetic: true,
                    text: result.output,
                  },
                  {
                    ...part,
                    id: part.id ?? Identifier.ascending("part"),
                    messageID: userMsg.id,
                    sessionID: input.sessionID,
                  },
                ]
              }

              let file = Bun.file(filePath)
              FileTime.read(input.sessionID, filePath)
              return [
                {
                  id: Identifier.ascending("part"),
                  messageID: userMsg.id,
                  sessionID: input.sessionID,
                  type: "text",
                  text: `Called the Read tool with the following input: {\"filePath\":\"${pathname}\"}`,
                  synthetic: true,
                },
                {
                  id: part.id ?? Identifier.ascending("part"),
                  messageID: userMsg.id,
                  sessionID: input.sessionID,
                  type: "file",
                  url: `data:${part.mime};base64,` + Buffer.from(await file.bytes()).toString("base64"),
                  mime: part.mime,
                  filename: part.filename!,
                  source: part.source,
                },
              ]
          }
        }
        return [
          {
            id: Identifier.ascending("part"),
            ...part,
            messageID: userMsg.id,
            sessionID: input.sessionID,
          },
        ]
      }),
    ).then((x) => x.flat())
    if (input.mode === "plan")
      userParts.push({
        id: Identifier.ascending("part"),
        messageID: userMsg.id,
        sessionID: input.sessionID,
        type: "text",
        text: PROMPT_PLAN,
        synthetic: true,
      })

    await updateMessage(userMsg)
    for (const part of userParts) {
      await updatePart(part)
    }

    if (isLocked(input.sessionID)) {
      return new Promise((resolve) => {
        const queue = state().queued.get(input.sessionID) ?? []
        queue.push({
          input: input,
          message: userMsg,
          parts: userParts,
          processed: false,
          callback: resolve,
        })
        state().queued.set(input.sessionID, queue)
      })
    }

    const model = await Provider.getModel(input.providerID, input.modelID)
    let msgs = await messages(input.sessionID)
    const session = await get(input.sessionID)

    if (session.revert) {
      const trimmed = []
      for (const msg of msgs) {
        if (
          msg.info.id > session.revert.messageID ||
          (msg.info.id === session.revert.messageID && session.revert.part === 0)
        ) {
          await Storage.remove("session/message/" + input.sessionID + "/" + msg.info.id)
          await Bus.publish(MessageV2.Event.Removed, {
            sessionID: input.sessionID,
            messageID: msg.info.id,
          })
          continue
        }

        if (msg.info.id === session.revert.messageID) {
          if (session.revert.part === 0) break
          msg.parts = msg.parts.slice(0, session.revert.part)
        }
        trimmed.push(msg)
      }
      msgs = trimmed
      await update(input.sessionID, (draft) => {
        draft.revert = undefined
      })
    }

    const previous = msgs.filter((x) => x.info.role === "assistant").at(-1)?.info as MessageV2.Assistant
    const outputLimit = Math.min(model.info.limit.output, OUTPUT_TOKEN_MAX) || OUTPUT_TOKEN_MAX

    // auto summarize if too long - much more aggressive approach
    if (previous && previous.tokens) {
      const tokens =
        previous.tokens.input + previous.tokens.cache.read + previous.tokens.cache.write + previous.tokens.output
      const contextLimit = model.info.limit.context || HARD_CONTEXT_LIMIT
      const effectiveLimit = Math.min(contextLimit, HARD_CONTEXT_LIMIT)
      // Smart threshold: leave room for response + safety margin
      // 180k limit - 32k response - 20k safety = 128k threshold (71%)
      const threshold = Math.max((effectiveLimit - outputLimit) * 0.71, 0)

      if (tokens > threshold) {
        log.warn("Proactive summarization triggered", {
          tokens,
          threshold,
          effectiveLimit,
          sessionID: input.sessionID,
        })
        await summarize({
          sessionID: input.sessionID,
          providerID: input.providerID,
          modelID: input.modelID,
        })
        return chat(input, retryCount)
      }
    }

    // Additional safety check: if we have too many messages, summarize regardless
    const messageCount = msgs.length
    if (messageCount > 100) {
      // Reasonable message count limit
      log.warn("Message count limit reached, forcing summarization", {
        messageCount,
        sessionID: input.sessionID,
      })
      await summarize({
        sessionID: input.sessionID,
        providerID: input.providerID,
        modelID: input.modelID,
      })
      return chat(input, retryCount)
    }

    using abort = lock(input.sessionID)

    const lastSummary = msgs.findLast((msg) => msg.info.role === "assistant" && msg.info.summary === true)
    if (lastSummary) msgs = msgs.filter((msg) => msg.info.id >= lastSummary.info.id)

    if (msgs.length === 1 && !session.parentID) {
      const small = (await Provider.getSmallModel(input.providerID)) ?? model
      generateText({
        maxOutputTokens: small.info.reasoning ? 1024 : 20,
        providerOptions: {
          [input.providerID]: small.info.options,
        },
        messages: [
          ...SystemPrompt.title(input.providerID).map(
            (x): CoreMessage => ({
              role: "system",
              content: x,
            }),
          ),
          ...MessageV2.toModelMessage([
            {
              info: {
                id: Identifier.ascending("message"),
                role: "user",
                sessionID: input.sessionID,
                time: {
                  created: Date.now(),
                },
              },
              parts: userParts,
            },
          ]),
        ],
        model: small.language,
      })
        .then((result) => {
          if (result.text)
            return Session.update(input.sessionID, (draft) => {
              draft.title = result.text
            })
        })
        .catch(() => {})
    }

    const mode = await Mode.get(input.mode ?? "build")
    let system = input.providerID === "anthropic" ? [PROMPT_ANTHROPIC_SPOOF.trim()] : []
    system.push(...(mode.prompt ? [mode.prompt] : SystemPrompt.provider(input.modelID)))
    system.push(...(await SystemPrompt.environment()))
    system.push(...(await SystemPrompt.custom()))
    // max 2 system prompt messages for caching purposes
    const [first, ...rest] = system
    system = [first, rest.join("\n")]

    const assistantMsg: MessageV2.Info = {
      id: Identifier.ascending("message"),
      role: "assistant",
      system,
      path: {
        cwd: app.path.cwd,
        root: app.path.root,
      },
      cost: 0,
      tokens: {
        input: 0,
        output: 0,
        reasoning: 0,
        cache: { read: 0, write: 0 },
      },
      modelID: input.modelID,
      providerID: input.providerID,
      time: {
        created: Date.now(),
      },
      sessionID: input.sessionID,
    }
    await updateMessage(assistantMsg)
    const tools: Record<string, AITool> = {}

    const processor = createProcessor(assistantMsg, model.info)

    for (const item of await Provider.tools(input.providerID)) {
      if (mode.tools[item.id] === false) continue
      if (input.tools?.[item.id] === false) continue
      if (session.parentID && item.id === "task") continue
      tools[item.id] = tool({
        id: item.id as any,
        description: item.description,
        inputSchema: item.parameters as ZodSchema,
        async execute(args, options) {
          const result = await item.execute(args, {
            sessionID: input.sessionID,
            abort: abort.signal,
            messageID: assistantMsg.id,
            metadata: async (val) => {
              const match = processor.partFromToolCall(options.toolCallId)
              if (match && match.state.status === "running") {
                await updatePart({
                  ...match,
                  state: {
                    title: val.title,
                    metadata: val.metadata,
                    status: "running",
                    input: args,
                    time: {
                      start: Date.now(),
                    },
                  },
                })
              }
            },
          })
          return result
        },
        toModelOutput(result) {
          return {
            type: "text",
            value: result.output,
          }
        },
      })
    }

    for (const [key, item] of Object.entries(await MCP.tools())) {
      if (mode.tools[key] === false) continue
      const execute = item.execute
      if (!execute) continue
      item.execute = async (args, opts) => {
        const result = await execute(args, opts)
        const output = result.content
          .filter((x: any) => x.type === "text")
          .map((x: any) => x.text)
          .join("\n\n")

        return {
          output,
        }
      }
      item.toModelOutput = (result) => {
        return {
          type: "text",
          value: result.output,
        }
      }
      tools[key] = item
    }

    const stream = streamText({
      onError() {},
      async prepareStep({ messages }) {
        const queue = (state().queued.get(input.sessionID) ?? []).filter((x) => !x.processed)
        if (queue.length) {
          for (const item of queue) {
            if (item.processed) continue
            messages.push(
              ...MessageV2.toModelMessage([
                {
                  info: item.message,
                  parts: item.parts,
                },
              ]),
            )
            item.processed = true
          }
          assistantMsg.time.completed = Date.now()
          await updateMessage(assistantMsg)
          Object.assign(assistantMsg, {
            id: Identifier.ascending("message"),
            role: "assistant",
            system,
            path: {
              cwd: app.path.cwd,
              root: app.path.root,
            },
            cost: 0,
            tokens: {
              input: 0,
              output: 0,
              reasoning: 0,
              cache: { read: 0, write: 0 },
            },
            modelID: input.modelID,
            providerID: input.providerID,
            time: {
              created: Date.now(),
            },
            sessionID: input.sessionID,
          })
          await updateMessage(assistantMsg)
        }
        return {
          messages,
        }
      },
      maxRetries: 10,
      maxOutputTokens: outputLimit,
      abortSignal: abort.signal,
      // stopWhen: stepCountIs(1000),
      providerOptions: {
        [input.providerID]: model.info.options,
      },
      messages: [
        ...system.map(
          (x): CoreMessage => ({
            role: "system",
            content: x,
          }),
        ),
        ...MessageV2.toModelMessage(msgs),
      ],
      temperature: model.info.temperature ? 0 : undefined,
      tools: model.info.tool_call === false ? undefined : wrapToolsWithFallback(tools),
      model: wrapLanguageModel({
        model: model.language,
        middleware: [
          {
            async transformParams(args) {
              if (args.type === "stream") {
                // @ts-expect-error
                args.params.prompt = ProviderTransform.message(args.params.prompt, input.providerID, input.modelID)
              }
              return args.params
            },
          },
        ],
      }),
    })

    let result
    try {
      result = await processor.process(stream)
    } catch (error: any) {
      // Enhanced error handling with comprehensive regex patterns
      const errorMessage = error?.message || ""
      const errorType = error?.type || ""

      // Comprehensive token limit error detection
      const isTokenLimitError =
        errorMessage.includes("prompt is too long") ||
        errorMessage.includes("tokens >") ||
        errorMessage.includes("maximum") ||
        errorMessage.includes("context limit") ||
        errorMessage.includes("exceed") ||
        error?.name === "AI_APICallError" ||
        // Regex patterns for various token limit formats
        /\d+\s*tokens?\s*>\s*\d+/i.test(errorMessage) || // "208405 tokens > 200000"
        /\d+\s*\+\s*\d+\s*>\s*\d+/i.test(errorMessage) || // "173941 + 32000 > 200000"
        /input\s+length.*exceed/i.test(errorMessage) || // "input length and max_tokens exceed"
        /context\s+limit/i.test(errorMessage) // "context limit"

      // Handle Anthropic overloaded errors with exponential backoff
      const isOverloadedError =
        errorType === "overloaded_error" ||
        errorMessage.includes("Overloaded") ||
        errorMessage.includes("rate limit") ||
        errorMessage.includes("too many requests")

      // Handle empty messages error (can happen after summarization)
      const isEmptyMessagesError =
        errorMessage.includes("at least one message is required") || errorMessage.includes("messages: at least one")

      if (isOverloadedError) {
        if (retryCount >= 3) {
          log.error("API overloaded after multiple retry attempts", {
            error: errorMessage,
            sessionID: input.sessionID,
            retryCount,
          })
          throw new Error(`API overloaded after ${retryCount} attempts. Please try again later.`)
        }

        // Exponential backoff: 2s, 4s, 8s
        const backoffMs = Math.pow(2, retryCount + 1) * 1000
        log.warn("API overloaded, retrying with backoff", {
          error: errorMessage,
          sessionID: input.sessionID,
          retryCount: retryCount + 1,
          backoffMs,
        })

        await new Promise((resolve) => setTimeout(resolve, backoffMs))
        return chat(input, retryCount + 1)
      }

      if (isEmptyMessagesError) {
        log.error("Empty messages after summarization", {
          error: errorMessage,
          sessionID: input.sessionID,
          retryCount,
        })
        throw new Error("Conversation became empty after summarization. Please start a new conversation.")
      }

      if (isTokenLimitError) {
        if (retryCount >= 2) {
          log.error("Token limit exceeded after multiple summarization attempts", {
            error: errorMessage,
            sessionID: input.sessionID,
            retryCount,
          })
          throw new Error(`Unable to reduce context below token limit after ${retryCount} attempts: ${errorMessage}`)
        }

        log.warn("Token limit exceeded, forcing summarization and retry", {
          error: errorMessage,
          sessionID: input.sessionID,
          retryCount: retryCount + 1,
        })
        await summarize({
          sessionID: input.sessionID,
          providerID: input.providerID,
          modelID: input.modelID,
        })
        return chat(input, retryCount + 1)
      }
      throw error
    }
    const queued = state().queued.get(input.sessionID) ?? []
    const unprocessed = queued.find((x) => !x.processed)
    if (unprocessed) {
      unprocessed.processed = true
      return chat(unprocessed.input, retryCount)
    }
    for (const item of queued) {
      item.callback(result)
    }
    state().queued.delete(input.sessionID)
    return result
  }

  function createProcessor(assistantMsg: MessageV2.Assistant, model: ModelsDev.Model) {
    const toolCalls: Record<string, MessageV2.ToolPart> = {}
    return {
      partFromToolCall(toolCallID: string) {
        return toolCalls[toolCallID]
      },
      async process(stream: StreamTextResult<Record<string, AITool>, never>) {
        try {
          let currentText: MessageV2.TextPart | undefined

          for await (const value of stream.fullStream) {
            log.info("part", {
              type: value.type,
            })
            switch (value.type) {
              case "start":
                const snapshot = await Snapshot.create(assistantMsg.sessionID)
                if (snapshot)
                  await updatePart({
                    id: Identifier.ascending("part"),
                    messageID: assistantMsg.id,
                    sessionID: assistantMsg.sessionID,
                    type: "snapshot",
                    snapshot,
                  })
                break

              case "tool-input-start":
                const part = await updatePart({
                  id: Identifier.ascending("part"),
                  messageID: assistantMsg.id,
                  sessionID: assistantMsg.sessionID,
                  type: "tool",
                  tool: value.toolName,
                  callID: value.id,
                  state: {
                    status: "pending",
                  },
                })
                toolCalls[value.id] = part as MessageV2.ToolPart
                break

              case "tool-input-delta":
                break

              case "tool-call": {
                const match = toolCalls[value.toolCallId]
                if (match) {
                  const part = await updatePart({
                    ...match,
                    state: {
                      status: "running",
                      input: value.input,
                      time: {
                        start: Date.now(),
                      },
                    },
                  })
                  toolCalls[value.toolCallId] = part as MessageV2.ToolPart
                }
                break
              }
              case "tool-result": {
                const match = toolCalls[value.toolCallId]
                if (match && match.state.status === "running") {
                  await updatePart({
                    ...match,
                    state: {
                      status: "completed",
                      input: value.input,
                      output: value.output.output,
                      metadata: value.output.metadata,
                      title: value.output.title,
                      time: {
                        start: match.state.time.start,
                        end: Date.now(),
                      },
                    },
                  })
                  delete toolCalls[value.toolCallId]
                  const snapshot = await Snapshot.create(assistantMsg.sessionID)
                  if (snapshot)
                    await updatePart({
                      id: Identifier.ascending("part"),
                      messageID: assistantMsg.id,
                      sessionID: assistantMsg.sessionID,
                      type: "snapshot",
                      snapshot,
                    })
                }
                break
              }

              case "tool-error": {
                const match = toolCalls[value.toolCallId]
                if (match && match.state.status === "running") {
                  await updatePart({
                    ...match,
                    state: {
                      status: "error",
                      input: value.input,
                      error: (value.error as any).toString(),
                      time: {
                        start: match.state.time.start,
                        end: Date.now(),
                      },
                    },
                  })
                  delete toolCalls[value.toolCallId]
                  const snapshot = await Snapshot.create(assistantMsg.sessionID)
                  if (snapshot)
                    await updatePart({
                      id: Identifier.ascending("part"),
                      messageID: assistantMsg.id,
                      sessionID: assistantMsg.sessionID,
                      type: "snapshot",
                      snapshot,
                    })
                }
                break
              }

              case "error":
                throw value.error

              case "start-step":
                await updatePart({
                  id: Identifier.ascending("part"),
                  messageID: assistantMsg.id,
                  sessionID: assistantMsg.sessionID,
                  type: "step-start",
                })
                break

              case "finish-step":
                const usage = getUsage(model, value.usage, value.providerMetadata)
                assistantMsg.cost += usage.cost
                assistantMsg.tokens = usage.tokens
                await updatePart({
                  id: Identifier.ascending("part"),
                  messageID: assistantMsg.id,
                  sessionID: assistantMsg.sessionID,
                  type: "step-finish",
                  tokens: usage.tokens,
                  cost: usage.cost,
                })
                await updateMessage(assistantMsg)
                break

              case "text-start":
                currentText = {
                  id: Identifier.ascending("part"),
                  messageID: assistantMsg.id,
                  sessionID: assistantMsg.sessionID,
                  type: "text",
                  text: "",
                  time: {
                    start: Date.now(),
                  },
                }
                break

              case "text":
                if (currentText) {
                  currentText.text += value.text
                  await updatePart(currentText)
                }
                break

              case "text-end":
                if (currentText && currentText.text) {
                  currentText.time = {
                    start: Date.now(),
                    end: Date.now(),
                  }
                  await updatePart(currentText)
                }
                currentText = undefined
                break

              case "finish":
                assistantMsg.time.completed = Date.now()
                await updateMessage(assistantMsg)
                break

              default:
                log.info("unhandled", {
                  ...value,
                })
                continue
            }
          }
        } catch (e) {
          log.error("", {
            error: e,
          })
          switch (true) {
            case e instanceof DOMException && e.name === "AbortError":
              assistantMsg.error = new MessageV2.AbortedError(
                { message: e.message },
                {
                  cause: e,
                },
              ).toObject()
              break
            case MessageV2.OutputLengthError.isInstance(e):
              assistantMsg.error = e
              break
            case LoadAPIKeyError.isInstance(e):
              assistantMsg.error = new MessageV2.AuthError(
                {
                  providerID: model.id,
                  message: e.message,
                },
                { cause: e },
              ).toObject()
              break
            case e instanceof Error:
              assistantMsg.error = new NamedError.Unknown({ message: e.toString() }, { cause: e }).toObject()
              break
            default:
              assistantMsg.error = new NamedError.Unknown({ message: JSON.stringify(e) }, { cause: e })
          }
          Bus.publish(Event.Error, {
            sessionID: assistantMsg.sessionID,
            error: assistantMsg.error,
          })
        }
        const p = await parts(assistantMsg.sessionID, assistantMsg.id)
        for (const part of p) {
          if (part.type === "tool" && part.state.status !== "completed") {
            updatePart({
              ...part,
              state: {
                status: "error",
                error: "Tool execution aborted",
                time: {
                  start: Date.now(),
                  end: Date.now(),
                },
                input: {},
              },
            })
          }
        }
        assistantMsg.time.completed = Date.now()
        await updateMessage(assistantMsg)
        return { info: assistantMsg, parts: p }
      },
    }
  }

  export async function revert(_input: { sessionID: string; messageID: string; part: number }) {
    // TODO
    /*
    const message = await getMessage(input.sessionID, input.messageID)
    if (!message) return
    const part = message.parts[input.part]
    if (!part) return
    const session = await get(input.sessionID)
    const snapshot =
      session.revert?.snapshot ?? (await Snapshot.create(input.sessionID))
    const old = (() => {
      if (message.role === "assistant") {
        const lastTool = message.parts.findLast(
          (part, index) =>
            part.type === "tool-invocation" && index < input.part,
        )
        if (lastTool && lastTool.type === "tool-invocation")
          return message.metadata.tool[lastTool.toolInvocation.toolCallId]
            .snapshot
      }
      return message.metadata.snapshot
    })()
    if (old) await Snapshot.restore(input.sessionID, old)
    await update(input.sessionID, (draft) => {
      draft.revert = {
        messageID: input.messageID,
        part: input.part,
        snapshot,
      }
    })
    */
  }

  export async function unrevert(sessionID: string) {
    const session = await get(sessionID)
    if (!session) return
    if (!session.revert) return
    if (session.revert.snapshot) await Snapshot.restore(sessionID, session.revert.snapshot)
    update(sessionID, (draft) => {
      draft.revert = undefined
    })
  }

  export async function summarize(input: { sessionID: string; providerID: string; modelID: string }) {
    using abort = lock(input.sessionID)
    const msgs = await messages(input.sessionID)
    const lastSummary = msgs.findLast((msg) => msg.info.role === "assistant" && msg.info.summary === true)
    const filtered = msgs.filter((msg) => !lastSummary || msg.info.id >= lastSummary.info.id)
    const model = await Provider.getModel(input.providerID, input.modelID)
    const app = App.info()
    const system = SystemPrompt.summarize(input.providerID)

    const next: MessageV2.Info = {
      id: Identifier.ascending("message"),
      role: "assistant",
      sessionID: input.sessionID,
      system,
      path: {
        cwd: app.path.cwd,
        root: app.path.root,
      },
      summary: true,
      cost: 0,
      modelID: input.modelID,
      providerID: input.providerID,
      tokens: {
        input: 0,
        output: 0,
        reasoning: 0,
        cache: { read: 0, write: 0 },
      },
      time: {
        created: Date.now(),
      },
    }
    await updateMessage(next)

    const processor = createProcessor(next, model.info)
    const stream = streamText({
      maxRetries: 10,
      abortSignal: abort.signal,
      model: model.language,
      messages: [
        ...system.map(
          (x): CoreMessage => ({
            role: "system",
            content: x,
          }),
        ),
        ...MessageV2.toModelMessage(filtered),
        {
          role: "user",
          content: [
            {
              type: "text",
              text: "Provide a detailed but concise summary of our conversation above. Focus on information that would be helpful for continuing the conversation, including what we did, what we're doing, which files we're working on, and what we're going to do next.",
            },
          ],
        },
      ],
    })

    const result = await processor.process(stream)
    return result
  }

  function isLocked(sessionID: string) {
    return state().pending.has(sessionID)
  }

  function lock(sessionID: string) {
    log.info("locking", { sessionID })
    if (state().pending.has(sessionID)) throw new BusyError(sessionID)
    const controller = new AbortController()
    state().pending.set(sessionID, controller)
    return {
      signal: controller.signal,
      [Symbol.dispose]() {
        log.info("unlocking", { sessionID })
        state().pending.delete(sessionID)
        Bus.publish(Event.Idle, {
          sessionID,
        })
      },
    }
  }

  function getUsage(model: ModelsDev.Model, usage: LanguageModelUsage, metadata?: ProviderMetadata) {
    const tokens = {
      input: usage.inputTokens ?? 0,
      output: usage.outputTokens ?? 0,
      reasoning: 0,
      cache: {
        write: (metadata?.["anthropic"]?.["cacheCreationInputTokens"] ??
          // @ts-expect-error
          metadata?.["bedrock"]?.["usage"]?.["cacheWriteInputTokens"] ??
          0) as number,
        read: usage.cachedInputTokens ?? 0,
      },
    }
    return {
      cost: new Decimal(0)
        .add(new Decimal(tokens.input).mul(model.cost.input).div(1_000_000))
        .add(new Decimal(tokens.output).mul(model.cost.output).div(1_000_000))
        .add(new Decimal(tokens.cache.read).mul(model.cost.cache_read ?? 0).div(1_000_000))
        .add(new Decimal(tokens.cache.write).mul(model.cost.cache_write ?? 0).div(1_000_000))
        .toNumber(),
      tokens,
    }
  }

  export class BusyError extends Error {
    constructor(public readonly sessionID: string) {
      super(`Session ${sessionID} is busy`)
    }
  }

  export async function initialize(input: {
    sessionID: string
    modelID: string
    providerID: string
    messageID: string
  }) {
    const app = App.info()
    await Session.chat({
      sessionID: input.sessionID,
      messageID: input.messageID,
      providerID: input.providerID,
      modelID: input.modelID,
      parts: [
        {
          id: Identifier.ascending("part"),
          type: "text",
          text: PROMPT_INITIALIZE.replace("${path}", app.path.root),
        },
      ],
    })
    await App.initialize()
  }
}
