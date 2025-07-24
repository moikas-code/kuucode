// This method is called when your extension is deactivated
export function deactivate() {}

import * as vscode from "vscode"

export function activate(context: vscode.ExtensionContext) {
  const TERMINAL_NAME = "kuuzuki"

  // Register command to open terminal in split screen and run kuuzuki
  let openTerminalDisposable = vscode.commands.registerCommand("kuuzuki.openTerminal", async () => {
    // Create a new terminal in split screen
    const terminal = vscode.window.createTerminal({
      name: TERMINAL_NAME,
      location: {
        viewColumn: vscode.ViewColumn.Beside,
        preserveFocus: false,
      },
    })

    terminal.show()
    terminal.sendText("KUUZUKI_THEME=system KUUZUKI_CALLER=vscode kuuzuki")
  })

  // Register command to add filepath to terminal
  let addFilepathDisposable = vscode.commands.registerCommand("kuuzuki.addFilepathToTerminal", async () => {
    const activeEditor = vscode.window.activeTextEditor

    if (!activeEditor) {
      vscode.window.showInformationMessage("No active file to get path from")
      return
    }

    const document = activeEditor.document
    const workspaceFolder = vscode.workspace.getWorkspaceFolder(document.uri)

    if (!workspaceFolder) {
      vscode.window.showInformationMessage("File is not in a workspace")
      return
    }

    // Get the relative path from workspace root
    const relativePath = vscode.workspace.asRelativePath(document.uri)
    let filepathWithAt = `@${relativePath}`

    // Check if there's a selection and add line numbers
    const selection = activeEditor.selection
    if (!selection.isEmpty) {
      // Convert to 1-based line numbers
      const startLine = selection.start.line + 1
      const endLine = selection.end.line + 1

      if (startLine === endLine) {
        // Single line selection
        filepathWithAt += `#L${startLine}`
      } else {
        // Multi-line selection
        filepathWithAt += `#L${startLine}-${endLine}`
      }
    }

    // Get or create terminal
    const terminal = vscode.window.activeTerminal
    if (!terminal) return

    if (terminal.name === TERMINAL_NAME) {
      terminal.sendText(filepathWithAt)
      terminal.show()
    }
  })

  context.subscriptions.push(openTerminalDisposable, addFilepathDisposable)
}
