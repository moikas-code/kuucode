import type { Argv } from "yargs"
import { Session } from "../../session"
import { UI } from "../ui"
import { cmd } from "./cmd"
import { bootstrap } from "../bootstrap"

export const SessionsCommand = cmd({
  command: "sessions",
  describe: "list all sessions",
  builder: (yargs: Argv) => {
    return yargs
      .option("current", {
        alias: ["c"],
        describe: "show current session ID only",
        type: "boolean",
      })
      .option("short", {
        alias: ["s"],
        describe: "show short session IDs (last 8 characters)",
        type: "boolean",
      })
  },
  handler: async (args) => {
    await bootstrap({ cwd: process.cwd() }, async () => {
      if (args.current) {
        // For current session, we'd need to track the last used session
        // This is a simplified implementation
        const list = Session.list()
        const first = await list.next()
        await list.return()

        if (first.done) {
          UI.error("No sessions found")
          return
        }

        const sessionID = args.short && first.value.id.length > 8 ? first.value.id.slice(-8) : first.value.id

        UI.println(sessionID)
        return
      }

      // List all sessions
      UI.empty()
      UI.println(UI.Style.TEXT_NORMAL_BOLD + "Sessions:")
      UI.empty()

      let count = 0
      for await (const session of Session.list()) {
        count++
        const sessionID = args.short && session.id.length > 8 ? session.id.slice(-8) : session.id

        const title = session.title.length > 50 ? session.title.substring(0, 47) + "..." : session.title

        const date = new Date(session.time.updated).toLocaleDateString()

        UI.println(
          UI.Style.TEXT_INFO_BOLD + sessionID,
          UI.Style.TEXT_NORMAL + " - " + title,
          UI.Style.TEXT_DIM + " (" + date + ")",
        )
      }

      if (count === 0) {
        UI.println(UI.Style.TEXT_DIM + "No sessions found")
      } else {
        UI.empty()
        UI.println(UI.Style.TEXT_DIM + `Total: ${count} sessions`)
        UI.empty()
        UI.println(UI.Style.TEXT_DIM + "Use 'kuuzuki run --session <id>' to continue a session")
      }
    })
  },
})
