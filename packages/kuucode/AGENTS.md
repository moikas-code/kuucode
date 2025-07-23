# kuucode agent guidelines

## About Kuucode

Kuucode is an AI-powered CLI tool for software development. It's also available as a **desktop application** that embeds kuucode in a custom terminal interface.

### CLI Tool

- **Primary Interface**: Terminal-based AI assistant
- **Commands**: Natural language queries, file operations, code generation
- **Usage**: `kuucode` command in any terminal

### Desktop App

- **Embedded Terminal**: Runs kuucode CLI in a custom terminal wrapper
- **Native Interface**: Desktop window with kuucode terminal embedded
- **Same Functionality**: Identical to CLI version, just in a desktop window

## Build/Test Commands

- **Install**: `bun install`
- **Run**: `bun run index.ts`
- **Desktop App**: `cd packages/desktop && bun run tauri dev`
- **Typecheck**: `bun run typecheck` (npm run typecheck)
- **Test**: `bun test` (runs all tests)
- **Single test**: `bun test test/tool/tool.test.ts` (specific test file)

## Code Style

- **Runtime**: Bun with TypeScript ESM modules
- **Imports**: Use relative imports for local modules, named imports preferred
- **Types**: Zod schemas for validation, TypeScript interfaces for structure
- **Naming**: camelCase for variables/functions, PascalCase for classes/namespaces
- **Error handling**: Use Result patterns, avoid throwing exceptions in tools
- **File structure**: Namespace-based organization (e.g., `Tool.define()`, `Session.create()`)

## IMPORTANT

- Try to keep things in one function unless composable or reusable
- DO NOT do unnecessary destructuring of variables
- DO NOT use `else` statements unless necessary
- DO NOT use `try`/`catch` if it can be avoided
- AVOID `try`/`catch` where possible
- AVOID `else` statements
- AVOID using `any` type
- AVOID `let` statements
- PREFER single word variable names where possible
- Use as many bun apis as possible like Bun.file()

## Architecture

### Core Components

- **Desktop App**: Tauri-based native application (`packages/desktop/`)
- **Terminal Interface**: Primary kuucode command execution environment
- **Tools**: Implement `Tool.Info` interface with `execute()` method
- **Context**: Pass `sessionID` in tool context, use `App.provide()` for DI
- **Validation**: All inputs validated with Zod schemas
- **Logging**: Use `Log.create({ service: "name" })` pattern
- **Storage**: Use `Storage` namespace for persistence

### Desktop App Structure

- **Terminal Wrapper**: Custom terminal that runs the kuucode CLI
- **Native Window**: Tauri-based desktop application
- **Embedded Process**: Spawns and manages kuucode CLI process
- **Same Commands**: All kuucode CLI functionality available

### Command Flow

```
Desktop App → Custom Terminal → kuucode CLI Process
                              ↓
                    Same as running `kuucode` in terminal
```

User Input → Kuucode Terminal → Built-in Commands OR Kuucode API
↓
┌─ help, clear, ls, cd
├─ editor <file> → kuucode /editor
└─ AI queries → kuucode chat API

```

```
