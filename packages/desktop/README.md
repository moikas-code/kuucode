# Kuucode Desktop - TUI Wrapper

A simple desktop wrapper around the kuucode TUI. This gives you the exact same kuucode experience but in a standalone desktop window.

## 🎯 What This Is

This is **NOT** a custom UI. It's a simple wrapper that:

- Runs the actual kuucode TUI inside a desktop window
- Provides the exact same interface as `kuucode` in terminal
- Adds basic desktop integration (menubar, window management)
- No custom components, no Monaco editor, no fancy UI

## 🚀 Quick Start

```bash
# From project root
bun run desktop

# Or directly
cd packages/desktop
bun run dev
```

## 🖥️ What You Get

- **Pure kuucode TUI** - Exactly the same as running `kuucode` in terminal
- **Desktop window** - Standalone app, no need for terminal
- **Basic menubar** - Restart kuucode, copy/paste, fullscreen
- **Auto-start** - Kuucode starts automatically when you open the app

## 📋 Requirements

- **Kuucode installed** - The app runs the actual `kuucode` command
- **Bun** - For running the wrapper
- **Rust** (optional) - For native desktop app

## 🔧 How It Works

1. Opens a terminal emulator (xterm.js) in a desktop window
2. Spawns the actual `kuucode` process
3. Pipes all input/output between the terminal and kuucode
4. You get the real kuucode TUI, just in a desktop window

## 🎨 Interface

```
┌─────────────────────────────────────────┐
│ Kuucode | Edit | View                   │ ← Basic menubar
├─────────────────────────────────────────┤
│                                         │
│  [Actual Kuucode TUI runs here]         │ ← Real kuucode interface
│                                         │
│  Same as running `kuucode` in terminal  │
│                                         │
└─────────────────────────────────────────┘
```

## 🚫 What This Is NOT

- ❌ Custom chat interface
- ❌ File explorer sidebar
- ❌ Monaco code editor
- ❌ Plugin management UI
- ❌ Custom terminal commands

## ✅ What This IS

- ✅ Pure kuucode TUI wrapper
- ✅ Desktop window for kuucode
- ✅ Exact same experience as terminal
- ✅ Simple and minimal

## 🛠️ Development

The app has only 3 files:

- `index.html` - Basic HTML container
- `src/main.ts` - Terminal wrapper logic
- `src-tauri/src/main.rs` - Process management

## 🎯 Perfect For

- Running kuucode without opening a terminal
- Having kuucode in its own dedicated window
- Desktop integration while keeping the pure TUI experience
- Users who love the kuucode TUI but want it as a desktop app

This is kuucode, unchanged, just wrapped in a desktop window. 🎯
