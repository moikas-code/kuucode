import { invoke } from "@tauri-apps/api/core"
import { listen } from "@tauri-apps/api/event"
import { Terminal } from "xterm"
import { FitAddon } from "xterm-addon-fit"

class KuucodeWrapper {
  private terminal: Terminal
  private fitAddon: FitAddon
  private terminalElement: HTMLElement

  constructor() {
    this.terminalElement = document.getElementById("terminal")!

    // Create xterm terminal that looks exactly like kuuzuki TUI
    this.terminal = new Terminal({
      theme: {
        background: "#000000",
        foreground: "#ffffff",
        cursor: "#ffffff",
        black: "#000000",
        red: "#ff5555",
        green: "#50fa7b",
        yellow: "#f1fa8c",
        blue: "#bd93f9",
        magenta: "#ff79c6",
        cyan: "#8be9fd",
        white: "#f8f8f2",
        brightBlack: "#44475a",
        brightRed: "#ff5555",
        brightGreen: "#50fa7b",
        brightYellow: "#f1fa8c",
        brightBlue: "#bd93f9",
        brightMagenta: "#ff79c6",
        brightCyan: "#8be9fd",
        brightWhite: "#ffffff",
      },
      fontFamily: "SF Mono, Monaco, Inconsolata, Fira Code, monospace",
      fontSize: 14,
      lineHeight: 1.2,
      cursorBlink: true,
      allowTransparency: false,
      convertEol: true,
    })

    // Add fit addon for responsive sizing
    this.fitAddon = new FitAddon()
    this.terminal.loadAddon(this.fitAddon)

    this.init()
  }

  private async init() {
    // Clear loading message
    this.terminalElement.innerHTML = ""

    // Open terminal in the DOM
    this.terminal.open(this.terminalElement)

    // Fit terminal to container
    this.fitAddon.fit()

    // Handle window resize
    window.addEventListener("resize", () => {
      this.fitAddon.fit()
    })

    // Listen for kuuzuki output
    await listen("kuuzuki-output", (event: any) => {
      this.terminal.write(event.payload + "\r\n")
    })

    // Listen for kuuzuki errors
    await listen("kuuzuki-error", (event: any) => {
      this.terminal.write("\x1b[31m" + event.payload + "\x1b[0m\r\n") // Red text for errors
    })

    // Handle terminal input
    this.terminal.onData((data) => {
      this.sendInputToKuucode(data)
    })

    // Listen for app ready event
    await listen("app-ready", () => {
      this.startKuucode()
    })

    // Start kuuzuki immediately if app is already ready
    this.startKuucode()
  }

  private async startKuucode() {
    try {
      this.terminal.writeln("\x1b[36m🚀 Starting Kuucode TUI...\x1b[0m\r\n")

      // Start kuuzuki TUI process
      const result = await invoke("start_kuucode_tui")
      console.log("Kuucode started:", result)
    } catch (error) {
      this.terminal.writeln(`\x1b[31m❌ Failed to start Kuucode: ${error}\x1b[0m\r\n`)
      this.terminal.writeln("\x1b[33mMake sure kuuzuki is installed and in your PATH.\x1b[0m\r\n")
      this.terminal.writeln("\x1b[33mInstall with: npm install -g kuuzuki\x1b[0m\r\n")
    }
  }

  private async sendInputToKuucode(input: string) {
    try {
      await invoke("send_input_to_kuucode", { input })
    } catch (error) {
      console.error("Failed to send input to kuuzuki:", error)
    }
  }
}

// Initialize the wrapper
new KuucodeWrapper()
