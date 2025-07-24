import { Log } from "../util/log"
import { tool, type Tool as AITool } from "ai"
import { z } from "zod"

const log = Log.create({ service: "tool-fallback" })

/**
 * Semantic similarity scoring for tool names and functionality
 */
function calculateSimilarity(requested: string, available: string): number {
  const req = requested.toLowerCase()
  const avail = available.toLowerCase()

  // Exact match
  if (req === avail) return 1.0

  // Contains match
  if (avail.includes(req) || req.includes(avail)) return 0.9

  // Prefix/suffix match
  if (avail.startsWith(req) || avail.endsWith(req) || req.startsWith(avail) || req.endsWith(avail)) return 0.8

  // Word boundary matches
  const reqWords = req.split(/[-_\s]+/)
  const availWords = avail.split(/[-_\s]+/)

  let matchCount = 0
  for (const reqWord of reqWords) {
    for (const availWord of availWords) {
      if (reqWord === availWord || availWord.includes(reqWord) || reqWord.includes(availWord)) {
        matchCount++
        break
      }
    }
  }

  if (matchCount > 0) {
    return 0.7 * (matchCount / Math.max(reqWords.length, availWords.length))
  }

  // Levenshtein distance for fuzzy matching
  const distance = levenshteinDistance(req, avail)
  const maxLength = Math.max(req.length, avail.length)
  const similarity = 1 - distance / maxLength

  return similarity > 0.5 ? similarity * 0.6 : 0
}

/**
 * Calculate Levenshtein distance between two strings
 */
function levenshteinDistance(a: string, b: string): number {
  const matrix = Array(b.length + 1)
    .fill(null)
    .map(() => Array(a.length + 1).fill(null))

  for (let i = 0; i <= a.length; i++) matrix[0][i] = i
  for (let j = 0; j <= b.length; j++) matrix[j][0] = j

  for (let j = 1; j <= b.length; j++) {
    for (let i = 1; i <= a.length; i++) {
      const indicator = a[i - 1] === b[j - 1] ? 0 : 1
      matrix[j][i] = Math.min(
        matrix[j][i - 1] + 1, // deletion
        matrix[j - 1][i] + 1, // insertion
        matrix[j - 1][i - 1] + indicator, // substitution
      )
    }
  }

  return matrix[b.length][a.length]
}

/**
 * Semantic keyword mapping for common tool functionality
 */
const SEMANTIC_KEYWORDS: Record<string, string[]> = {
  fork: ["fork-parity", "git", "repository", "upstream", "branch", "merge"],
  git: ["fork-parity", "repository", "commit", "branch", "diff"],
  repo: ["fork-parity", "git", "repository", "upstream"],
  upstream: ["fork-parity", "git", "repository", "branch"],
  file: ["read", "write", "edit", "glob", "list"],
  search: ["grep", "find", "glob", "semantic"],
  code: ["format", "lint", "analyze", "check"],
  security: ["scan", "vulnerability", "audit", "check"],
  test: ["analyze", "coverage", "quality"],
  performance: ["analyze", "optimize", "hotspot"],
  documentation: ["quality", "check", "analyze"],
}

/**
 * Find the best alternative tools for a requested tool name
 */
export function findAlternatives(
  requestedTool: string,
  availableTools: Record<string, AITool>,
  maxSuggestions = 3,
): Array<{ name: string; score: number; reason: string }> {
  const alternatives: Array<{ name: string; score: number; reason: string }> = []

  // Direct similarity matching
  for (const [toolName] of Object.entries(availableTools)) {
    const similarity = calculateSimilarity(requestedTool, toolName)
    if (similarity > 0.3) {
      alternatives.push({
        name: toolName,
        score: similarity,
        reason: similarity > 0.8 ? "name similarity" : "partial name match",
      })
    }
  }

  // Semantic keyword matching
  const requestedLower = requestedTool.toLowerCase()
  const semanticMatches = SEMANTIC_KEYWORDS[requestedLower] || []

  for (const keyword of semanticMatches) {
    for (const [toolName] of Object.entries(availableTools)) {
      const keywordSimilarity = calculateSimilarity(keyword, toolName)
      if (keywordSimilarity > 0.5) {
        const existing = alternatives.find((alt) => alt.name === toolName)
        if (existing) {
          existing.score = Math.max(existing.score, keywordSimilarity * 0.9)
          existing.reason = "semantic match"
        } else {
          alternatives.push({
            name: toolName,
            score: keywordSimilarity * 0.9,
            reason: "semantic match",
          })
        }
      }
    }
  }

  // Description-based matching
  for (const [toolName, toolDef] of Object.entries(availableTools)) {
    if (toolDef.description) {
      const descSimilarity = calculateSimilarity(requestedTool, toolDef.description)
      if (descSimilarity > 0.4) {
        const existing = alternatives.find((alt) => alt.name === toolName)
        if (existing) {
          existing.score = Math.max(existing.score, descSimilarity * 0.8)
          existing.reason = "description match"
        } else {
          alternatives.push({
            name: toolName,
            score: descSimilarity * 0.8,
            reason: "description match",
          })
        }
      }
    }
  }

  // Sort by score and return top suggestions
  return alternatives
    .sort((a, b) => b.score - a.score)
    .slice(0, maxSuggestions)
    .filter((alt) => alt.score > 0.3)
}

/**
 * Create a fallback tool that suggests alternatives
 */
export function createFallbackTool(
  requestedTool: string,
  alternatives: Array<{ name: string; score: number; reason: string }>,
): AITool {
  return tool({
    id: `fallback.${requestedTool}` as any,
    description: `Fallback for unavailable tool '${requestedTool}'`,
    inputSchema: z.object({}),
    async execute() {
      const message =
        alternatives.length > 0
          ? `Tool '${requestedTool}' not found. Did you mean one of these?\n\n` +
            alternatives
              .map((alt) => `• **${alt.name}** (${Math.round(alt.score * 100)}% match, ${alt.reason})`)
              .join("\n") +
            "\n\nPlease specify which tool you'd like to use."
          : `Tool '${requestedTool}' not found and no similar tools available.`

      log.info("tool fallback triggered", {
        requested: requestedTool,
        alternatives: alternatives.map((a) => a.name),
      })

      return {
        output: message,
        title: `Tool '${requestedTool}' not available`,
        metadata: {
          fallback: true,
          requested: requestedTool,
          alternatives: alternatives.map((a) => a.name),
        },
      }
    },
    toModelOutput(result) {
      return {
        type: "text",
        value: result.output,
      }
    },
  })
}

/**
 * Automatically select the best alternative if confidence is high enough
 */
export function autoSelectBestAlternative(
  alternatives: Array<{ name: string; score: number; reason: string }>,
  confidenceThreshold = 0.8,
): string | null {
  if (alternatives.length === 0) return null

  const best = alternatives[0]
  if (best.score >= confidenceThreshold) {
    log.info("auto-selecting best alternative", {
      selected: best.name,
      score: best.score,
      reason: best.reason,
    })
    return best.name
  }

  return null
}

/**
 * Enhanced tool wrapper that provides fallback functionality
 */
export function wrapToolsWithFallback(
  tools: Record<string, AITool>,
  enableAutoSelect = true,
  autoSelectThreshold = 0.8,
): Record<string, AITool> {
  return new Proxy(tools, {
    get(target, prop: string) {
      // If tool exists, return it
      if (prop in target) {
        return target[prop]
      }

      // Tool doesn't exist, find alternatives
      const alternatives = findAlternatives(prop, target as Record<string, AITool>)

      // Try auto-selection if enabled
      if (enableAutoSelect) {
        const autoSelected = autoSelectBestAlternative(alternatives, autoSelectThreshold)
        if (autoSelected && autoSelected in target) {
          log.info("auto-selected alternative tool", {
            requested: prop,
            selected: autoSelected,
          })
          return target[autoSelected]
        }
      }

      // Return fallback tool with suggestions
      return createFallbackTool(prop, alternatives)
    },

    has() {
      // Always return true to prevent "tool not found" errors
      return true
    },

    ownKeys(target) {
      return Reflect.ownKeys(target)
    },

    getOwnPropertyDescriptor(target, prop) {
      if (prop in target) {
        return Reflect.getOwnPropertyDescriptor(target, prop)
      }
      // Return a descriptor for fallback tools
      return {
        enumerable: false,
        configurable: true,
        value: createFallbackTool(prop as string, []),
      }
    },
  })
}
