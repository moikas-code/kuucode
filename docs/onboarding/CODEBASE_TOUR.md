# Codebase Tour

Welcome to the guided tour of the Kuucode codebase! This document will walk you through the key files and directories, explaining what each does and how they fit together.

## 🗺️ Project Structure Overview

```
kuucode/
├── 📁 packages/           # Main application packages
│   ├── 📁 kuucode/        # TypeScript server (core logic)
│   ├── 📁 tui/            # Go terminal interface
│   ├── 📁 web/            # Documentation website
│   └── 📁 function/       # Serverless functions
├── 📁 sdks/               # Generated client SDKs
│   ├── 📁 github/         # GitHub Action
│   └── 📁 vscode/         # VS Code extension
├── 📁 scripts/            # Build and development scripts
├── 📁 config/             # SDK generation configuration
├── 📁 docs/               # Documentation
└── 📄 Configuration files # Package.json, configs, etc.
```

## 🎯 Entry Points - Start Here

### **Main Server Entry Point**
📄 `packages/kuucode/src/index.ts`

```typescript
#!/usr/bin/env bun

import { parseArgs } from "util"
import { serve } from "./server/server.js"
import { auth } from "./cli/auth.js"
import { run } from "./cli/run.js"

// Main CLI entry point
// Handles: serve, auth, run commands
// This is where everything starts!
```

**What it does:**
- Parses command line arguments
- Routes to different commands (serve, auth, run)
- Sets up the main application context

### **TUI Entry Point**
📄 `packages/tui/main.go`

```go
package main

import (
    "github.com/moikas-code/kuucode/internal/tui"
)

func main() {
    // Initialize and run the terminal UI
    tui.Run()
}
```

**What it does:**
- Initializes the Go TUI application
- Sets up Bubble Tea framework
- Connects to the TypeScript server

## 🖥️ Server Package Deep Dive

### **Core Server Logic**
📄 `packages/kuucode/src/server/server.ts`

```typescript
import { Hono } from "hono"
import { streamText } from "ai"

const app = new Hono()

// Key endpoints:
app.get("/app", getApp)                    // App info
app.get("/session", listSessions)         // List sessions
app.post("/session", createSession)       // Create session
app.post("/session/:id/message", chat)    // Send message
app.get("/event", streamEvents)           // Server-sent events

export { app }
```

**Key Routes:**
- `/app` - Application metadata and status
- `/session/*` - Session management and chat
- `/config/*` - Configuration endpoints
- `/file/*` - File operations
- `/find/*` - Search functionality
- `/event` - Real-time event streaming

### **Provider System**
📁 `packages/kuucode/src/provider/`

```
provider/
├── 📄 anthropic.ts       # Claude integration
├── 📄 openai.ts          # GPT integration
├── 📄 base.ts            # Provider interface
├── 📄 manager.ts         # Provider selection
└── 📄 types.ts           # Shared types
```

**Key File:** `base.ts`
```typescript
export interface Provider {
  name: string
  models: Model[]
  authenticate(credentials: Credentials): Promise<void>
  generateText(request: GenerateRequest): AsyncIterable<TextChunk>
}
```

### **Session Management**
📁 `packages/kuucode/src/session/`

```
session/
├── 📄 manager.ts         # Session lifecycle
├── 📄 storage.ts         # Persistence
├── 📄 context.ts         # Context management
└── 📄 types.ts           # Session types
```

**Key File:** `manager.ts`
```typescript
export class SessionManager {
  async create(config: SessionConfig): Promise<Session>
  async get(id: string): Promise<Session>
  async update(id: string, updates: Partial<Session>): Promise<void>
  async delete(id: string): Promise<void>
  async list(): Promise<Session[]>
}
```

### **Tool System**
📁 `packages/kuucode/src/tool/`

```
tool/
├── 📄 bash.ts            # Shell command execution
├── 📄 edit.ts            # File editing
├── 📄 read.ts            # File reading
├── 📄 glob.ts            # File pattern matching
├── 📄 grep.ts            # Text searching
└── 📄 ...                # More tools
```

**Example Tool:** `edit.ts`
```typescript
export const editTool = {
  name: "edit",
  description: "Edit files with precise replacements",
  parameters: {
    filePath: { type: "string", description: "Path to file" },
    oldString: { type: "string", description: "Text to replace" },
    newString: { type: "string", description: "Replacement text" }
  },
  execute: async ({ filePath, oldString, newString }) => {
    // File editing logic
  }
}
```

### **Configuration System**
📁 `packages/kuucode/src/config/`

```
config/
├── 📄 loader.ts          # Config file loading
├── 📄 schema.ts          # Validation schemas
├── 📄 defaults.ts        # Default values
└── 📄 types.ts           # Config types
```

## 🖼️ TUI Package Deep Dive

### **Main TUI Application**
📄 `packages/tui/internal/tui/tui.go`

```go
type Model struct {
    state      AppState
    chat       chat.Model
    fileViewer fileviewer.Model
    status     status.Model
    // ... other components
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    // Handle keyboard input, API responses, etc.
}

func (m Model) View() string {
    // Render the current state to terminal
}
```

### **Component Architecture**
📁 `packages/tui/internal/components/`

```
components/
├── 📁 chat/              # Chat interface
│   ├── 📄 chat.go        # Main chat component
│   ├── 📄 input.go       # Message input
│   └── 📄 messages.go    # Message display
├── 📁 fileviewer/        # File browser
│   ├── 📄 fileviewer.go  # File tree
│   └── 📄 preview.go     # File preview
├── 📁 status/            # Status bar
├── 📁 modal/             # Modal dialogs
└── 📁 ...                # Other components
```

### **API Client**
📄 `packages/tui/internal/api/client.go`

```go
type Client struct {
    baseURL    string
    httpClient *http.Client
}

func (c *Client) SendMessage(sessionID, message string) error {
    // Send chat message to server
}

func (c *Client) StreamEvents(sessionID string) (<-chan Event, error) {
    // Listen for server-sent events
}
```

### **Generated SDK**
📁 `packages/tui/sdk/`

This directory contains the auto-generated Go SDK from OpenAPI Generator:

```
sdk/
├── 📄 client.go          # Main API client
├── 📄 session.go         # Session operations
├── 📄 app.go             # App operations
├── 📄 config.go          # Config operations
└── 📄 ...                # Other API endpoints
```

## 🎨 Styling and Themes

### **TUI Styling**
📁 `packages/tui/internal/styles/`

```
styles/
├── 📄 theme.go           # Color themes
├── 📄 markdown.go        # Markdown rendering
└── 📄 syntax.go          # Syntax highlighting
```

**Example:** `theme.go`
```go
type Theme struct {
    Primary   lipgloss.Color
    Secondary lipgloss.Color
    Accent    lipgloss.Color
    // ... more colors
}

var DefaultTheme = Theme{
    Primary:   lipgloss.Color("#7C3AED"),
    Secondary: lipgloss.Color("#A855F7"),
    // ...
}
```

## 🔧 Build and Development

### **Scripts Directory**
📁 `scripts/`

```
scripts/
├── 📄 generate-sdks      # SDK generation (replaces Stainless)
├── 📄 stainless          # Legacy compatibility script
├── 📄 hooks              # Git hooks setup
├── 📄 release            # Release automation
└── 📄 stats.ts           # Download statistics
```

### **SDK Generation Config**
📁 `config/`

```
config/
├── 📄 typescript-config.json  # TypeScript SDK settings
├── 📄 go-config.json          # Go SDK settings
└── 📄 python-config.json      # Python SDK settings
```

## 📚 Documentation

### **Web Documentation**
📁 `packages/web/`

```
web/
├── 📁 src/
│   ├── 📁 content/docs/   # Documentation content
│   ├── 📁 components/     # Astro components
│   └── 📁 pages/          # Website pages
├── 📄 astro.config.mjs    # Astro configuration
└── 📄 package.json        # Dependencies
```

### **Onboarding Docs**
📁 `docs/onboarding/` (this directory!)

```
onboarding/
├── 📄 README.md           # Start here
├── 📄 PROJECT_OVERVIEW.md # What Kuucode is
├── 📄 ARCHITECTURE.md     # System design
├── 📄 CODEBASE_TOUR.md    # This file!
└── 📄 ...                 # More guides
```

## 🔌 Extensions and SDKs

### **VS Code Extension**
📁 `sdks/vscode/`

```
vscode/
├── 📁 src/
│   └── 📄 extension.ts    # Main extension logic
├── 📄 package.json        # Extension manifest
└── 📄 README.md           # Extension docs
```

### **GitHub Action**
📁 `sdks/github/`

```
github/
├── 📁 src/
│   ├── 📄 index.ts        # Action entry point
│   └── 📄 types.ts        # GitHub API types
├── 📄 action.yml          # Action definition
└── 📄 package.json        # Dependencies
```

## 🔍 Key Files to Understand

### **Configuration Files**

| File | Purpose | Key Contents |
|------|---------|--------------|
| `package.json` | Root package config | Scripts, dependencies, workspaces |
| `kuucode.json` | App configuration | Provider settings, preferences |
| `tsconfig.json` | TypeScript config | Compilation settings |
| `go.mod` | Go module config | Dependencies, module path |

### **Important Constants and Types**

📄 `packages/kuucode/src/types.ts`
```typescript
export interface Session {
  id: string
  name: string
  messages: Message[]
  config: SessionConfig
  createdAt: Date
  updatedAt: Date
}

export interface Message {
  id: string
  role: 'user' | 'assistant' | 'system'
  content: string
  timestamp: Date
}
```

📄 `packages/tui/internal/types/types.go`
```go
type AppState struct {
    CurrentSession *Session
    Sessions       []Session
    Config         *Config
    // ...
}
```

## 🚀 Data Flow Through Key Files

### **Chat Message Flow**

1. **User Input** → `packages/tui/internal/components/chat/input.go`
2. **HTTP Request** → `packages/tui/internal/api/client.go`
3. **Server Routing** → `packages/kuucode/src/server/server.ts`
4. **Provider Call** → `packages/kuucode/src/provider/anthropic.ts`
5. **Streaming Response** → `packages/kuucode/src/server/streaming.ts`
6. **TUI Update** → `packages/tui/internal/components/chat/messages.go`

### **File Operation Flow**

1. **File Request** → `packages/tui/internal/components/fileviewer/fileviewer.go`
2. **API Call** → `packages/kuucode/src/server/server.ts` (`/file` endpoint)
3. **Tool Execution** → `packages/kuucode/src/tool/read.ts`
4. **File System** → Native file operations
5. **Response** → Back through the chain

## 🎯 Where to Start Coding

### **Adding a New Tool**
1. Create `packages/kuucode/src/tool/your-tool.ts`
2. Implement the tool interface
3. Register in `packages/kuucode/src/tool/index.ts`
4. Regenerate SDKs: `./scripts/generate-sdks --dev`

### **Adding a New Provider**
1. Create `packages/kuucode/src/provider/your-provider.ts`
2. Implement the Provider interface
3. Register in `packages/kuucode/src/provider/manager.ts`
4. Add configuration schema

### **Adding TUI Features**
1. Create component in `packages/tui/internal/components/`
2. Add to main TUI model in `packages/tui/internal/tui/tui.go`
3. Handle events in the Update function
4. Render in the View function

### **Modifying the API**
1. Edit routes in `packages/kuucode/src/server/server.ts`
2. Update types if needed
3. Regenerate SDKs: `./scripts/generate-sdks --dev`
4. Update TUI client code if needed

## 🔍 Debugging Entry Points

### **Server Issues**
- Check `packages/kuucode/src/index.ts` for startup issues
- Look at `packages/kuucode/src/server/server.ts` for API problems
- Examine provider files for AI integration issues

### **TUI Issues**
- Start with `packages/tui/main.go` for initialization problems
- Check `packages/tui/internal/tui/tui.go` for UI issues
- Look at `packages/tui/internal/api/client.go` for communication problems

### **SDK Issues**
- Check `./scripts/generate-sdks` for generation problems
- Look at config files in `config/` for configuration issues
- Examine generated SDKs in `packages/tui/sdk/` for Go issues

---

**Next Steps:**
- **[Key Concepts](KEY_CONCEPTS.md)** - Understand the important patterns
- **[API Flow](API_FLOW.md)** - See how requests flow through the system
- **[Common Tasks](COMMON_TASKS.md)** - Learn frequent development patterns

**Need to find something specific?** Use your editor's search functionality to find files, functions, or patterns across the codebase!