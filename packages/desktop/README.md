# Kuuzuki Desktop - TUI Wrapper

A simple desktop wrapper around the kuuzuki TUI. This gives you the exact same kuuzuki experience but in a standalone desktop window.

## 🎯 What This Is

This is **NOT** a custom UI. It's a simple wrapper that:

- Runs the actual kuuzuki TUI inside a desktop window
- Provides the exact same interface as `kuuzuki` in terminal
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

- **Pure kuuzuki TUI** - Exactly the same as running `kuuzuki` in terminal
- **Desktop window** - Standalone app, no need for terminal
- **Basic menubar** - Restart kuuzuki, copy/paste, fullscreen
- **Auto-start** - Kuuzuki starts automatically when you open the app

## 📋 Requirements

- **Kuuzuki installed** - The app runs the actual `kuuzuki` command
- **Bun** - For running the wrapper
- **Rust** (optional) - For native desktop app

## 🔧 How It Works

1. Opens a terminal emulator (xterm.js) in a desktop window
2. Spawns the actual `kuuzuki` process
3. Pipes all input/output between the terminal and kuuzuki
4. You get the real kuuzuki TUI, just in a desktop window

## 🎨 Interface

```
┌─────────────────────────────────────────┐
│ Kuuzuki | Edit | View                   │ ← Basic menubar
├─────────────────────────────────────────┤
│                                         │
│  [Actual Kuuzuki TUI runs here]         │ ← Real kuuzuki interface
│                                         │
│  Same as running `kuuzuki` in terminal  │
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

- ✅ Pure kuuzuki TUI wrapper
- ✅ Desktop window for kuuzuki
- ✅ Exact same experience as terminal
- ✅ Simple and minimal

## 🛠️ Development

The app has only 3 files:

- `index.html` - Basic HTML container
- `src/main.ts` - Terminal wrapper logic
- `src-tauri/src/main.rs` - Process management

## 🎯 Perfect For

- Running kuuzuki without opening a terminal
- Having kuuzuki in its own dedicated window
- Desktop integration while keeping the pure TUI experience
- Users who love the kuuzuki TUI but want it as a desktop app

This is kuuzuki, unchanged, just wrapped in a desktop window. 🎯
