# Kuuzuki AGENTS.md Update - Desktop App Focus

## Update Summary

Updated `/packages/kuuzuki/AGENTS.md` to reflect that kuuzuki is now a standalone desktop application rather than just a CLI tool.

## Changes Made

### 1. Added Desktop App Description
- Emphasized kuuzuki as a **standalone desktop application**
- Highlighted AI-first development workflow
- Described terminal-based interface as primary
- Listed key features: native menubar, filesystem sidebar, AI integration

### 2. Updated Build Commands
- Added desktop app development command: `cd packages/desktop && bun run tauri dev`
- Maintained existing CLI build commands for compatibility

### 3. Enhanced Architecture Section
- Added **Core Components** section describing desktop app structure
- Added **Desktop App Structure** with frontend/backend details
- Added **Command Flow** diagram showing user interaction patterns
- Described Tauri-based architecture with Rust backend and TypeScript frontend

### 4. Key Features Highlighted
- **Primary Interface**: Kuuzuki terminal with natural language commands
- **Native Desktop**: Full menubar integration and filesystem sidebar
- **AI Integration**: Direct command execution with real-time AI assistance
- **File Operations**: Built-in editor accessible via `/editor` command
- **Cross-platform**: Available on macOS, Windows, and Linux

## Architecture Overview

```
┌─────────────────────────────────────────────────────┐
│ Native Menubar (Kuuzuki | Edit | View | Help)       │
├─────────────┬───────────────────────────────────────┤
│ Filesystem  │ 🚀 KUUZUKI TERMINAL (PRIMARY)        │
│ Sidebar     │                                       │
│             │ kuuzuki> help                         │
│ 📁 src/     │ kuuzuki> editor main.ts               │
│ 📄 main.ts  │ kuuzuki> explain this function        │
│ 📄 app.ts   │                                       │
│             │ Terminal Output Area                  │
│             │                                       │
├─────────────┴───────────────────────────────────────┤
│ Status: ● Server Running | 15 commands | 05:23     │
└─────────────────────────────────────────────────────┘
```

## Impact

This update ensures that developers working on kuuzuki understand:
1. The primary focus is now the desktop application
2. The terminal interface is the main interaction method
3. AI integration is central to the user experience
4. The architecture supports both CLI and desktop workflows

## Next Steps

- Consider updating other documentation files to reflect desktop app focus
- Ensure build scripts and deployment processes account for desktop app
- Update README files to highlight desktop app as primary interface