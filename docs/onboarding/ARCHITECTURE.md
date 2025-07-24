# Architecture Deep Dive

This document explains how Kuuzuki is architected, why certain decisions were made, and how the components work together.

## 🏗️ High-Level Architecture

```
┌─────────────────────────────────────────────────────────────────┐
│                        User Interface Layer                     │
├─────────────────────────────────────────────────────────────────┤
│  ┌─────────────────┐  ┌─────────────────┐  ┌─────────────────┐  │
│  │   Terminal UI   │  │   VS Code Ext   │  │   Web Client    │  │
│  │     (Go)        │  │  (TypeScript)   │  │  (TypeScript)   │  │
│  └─────────────────┘  └─────────────────┘  └─────────────────┘  │
└─────────────────────────────────────────────────────────────────┘
                                │
                                ▼
┌─────────────────────────────────────────────────────────────────┐
│                      API Gateway Layer                         │
├─────────────────────────────────────────────────────────────────┤
│  ┌─────────────────────────────────────────────────────────────┐ │
│  │              Kuuzuki Server (TypeScript)                   │ │
│  │  ┌─────────────┐ ┌─────────────┐ ┌─────────────────────┐  │ │
│  │  │   Router    │ │ Middleware  │ │   Authentication    │  │ │
│  │  └─────────────┘ └─────────────┘ └─────────────────────┘  │ │
│  └─────────────────────────────────────────────────────────────┘ │
└─────────────────────────────────────────────────────────────────┘
                                │
                                ▼
┌─────────────────────────────────────────────────────────────────┐
│                     Business Logic Layer                       │
├─────────────────────────────────────────────────────────────────┤
│  ┌─────────────┐ ┌─────────────┐ ┌─────────────┐ ┌───────────┐  │
│  │   Session   │ │   Config    │ │    File     │ │   Tool    │  │
│  │ Management  │ │ Management  │ │  Operations │ │ Execution │  │
│  └─────────────┘ └─────────────┘ └─────────────┘ └───────────┘  │
└─────────────────────────────────────────────────────────────────┘
                                │
                                ▼
┌─────────────────────────────────────────────────────────────────┐
│                    Provider Integration Layer                   │
├─────────────────────────────────────────────────────────────────┤
│  ┌─────────────┐ ┌─────────────┐ ┌─────────────┐ ┌───────────┐  │
│  │  Anthropic  │ │   OpenAI    │ │   Google    │ │   Local   │  │
│  │   Claude    │ │     GPT     │ │   Gemini    │ │  Models   │  │
│  └─────────────┘ └─────────────┘ └─────────────┘ └───────────┘  │
└─────────────────────────────────────────────────────────────────┘
```

## 🎯 Core Design Principles

### **1. Separation of Concerns**
- **UI Layer**: Handles user interaction and display
- **API Layer**: Manages HTTP requests and responses
- **Business Logic**: Core functionality and data processing
- **Provider Layer**: AI model integration and communication

### **2. Client-Server Architecture**
- **Stateless Server**: Business logic in TypeScript server
- **Thin Clients**: UI components communicate via REST API
- **Real-time Updates**: Server-sent events for streaming responses
- **Multiple Clients**: TUI, VS Code extension, web clients

### **3. Provider Agnostic**
- **Unified Interface**: Same API for all AI providers
- **Easy Switching**: Change providers without code changes
- **Fallback Support**: Automatic failover between providers
- **Local Support**: Works with local models (Ollama, etc.)

## 🔧 Component Architecture

### **TypeScript Server (`packages/kuuzuki/`)**

```
src/
├── index.ts              # Main entry point
├── server/
│   ├── server.ts         # HTTP server and routes
│   ├── middleware.ts     # Request/response processing
│   └── streaming.ts      # Server-sent events
├── provider/
│   ├── anthropic.ts      # Claude integration
│   ├── openai.ts         # GPT integration
│   ├── base.ts           # Provider interface
│   └── manager.ts        # Provider selection logic
├── session/
│   ├── manager.ts        # Session lifecycle
│   ├── storage.ts        # Persistence layer
│   └── context.ts        # Context management
├── config/
│   ├── loader.ts         # Configuration loading
│   ├── schema.ts         # Config validation
│   └── defaults.ts       # Default settings
├── auth/
│   ├── providers.ts      # OAuth integrations
│   ├── tokens.ts         # Token management
│   └── middleware.ts     # Auth middleware
└── tool/
    ├── bash.ts           # Shell command execution
    ├── edit.ts           # File editing
    ├── read.ts           # File reading
    └── ...               # Other tools
```

**Key Responsibilities:**
- **HTTP API**: RESTful endpoints for all operations
- **Provider Management**: AI model integration and switching
- **Session Handling**: Conversation persistence and context
- **Tool Execution**: File operations, shell commands, etc.
- **Authentication**: User auth and API key management

### **Go TUI (`packages/tui/`)**

```
internal/
├── tui/
│   ├── tui.go            # Main TUI application
│   ├── state.go          # Application state
│   └── events.go         # Event handling
├── components/
│   ├── chat/             # Chat interface
│   ├── fileviewer/       # File browser
│   ├── status/           # Status bar
│   ├── modal/            # Modal dialogs
│   └── ...               # Other UI components
├── api/
│   ├── client.go         # HTTP client wrapper
│   ├── streaming.go      # SSE handling
│   └── types.go          # API types
├── styles/
│   ├── theme.go          # Color themes
│   ├── markdown.go       # Markdown rendering
│   └── syntax.go         # Syntax highlighting
└── util/
    ├── filesystem.go     # File operations
    ├── git.go            # Git integration
    └── lsp.go            # Language server protocol
```

**Key Responsibilities:**
- **User Interface**: Terminal-based chat and file browsing
- **Real-time Updates**: Streaming response handling
- **Keyboard Navigation**: Efficient keyboard shortcuts
- **File Integration**: Project-aware file operations
- **Theme Support**: Customizable appearance

## 🔄 Data Flow Architecture

### **Request Lifecycle**

```
1. User Input
   │
   ▼
2. TUI Event Handler
   │
   ▼
3. HTTP Request to Server
   │
   ▼
4. Server Router & Middleware
   │
   ▼
5. Business Logic Processing
   │
   ▼
6. Provider API Call
   │
   ▼
7. Streaming Response
   │
   ▼
8. Server-Sent Events
   │
   ▼
9. TUI Response Handler
   │
   ▼
10. UI Update & Display
```

### **Detailed Flow Example: Chat Message**

```typescript
// 1. User types message in TUI
// TUI sends POST request to /session/{id}/message

// 2. Server receives request
app.post('/session/:id/message', async (c) => {
  const { message } = await c.req.json()
  const sessionId = c.req.param('id')
  
  // 3. Load session context
  const session = await sessionManager.get(sessionId)
  
  // 4. Select appropriate provider
  const provider = await providerManager.select(session.config)
  
  // 5. Stream response
  return streamText({
    model: provider.model,
    messages: [...session.messages, { role: 'user', content: message }],
    onChunk: (chunk) => {
      // 6. Send chunk via SSE
      eventStream.send(sessionId, chunk)
    }
  })
})

// 7. TUI receives SSE events and updates UI
```

## 🗄️ Data Architecture

### **Session Management**

```
Session {
  id: string
  name: string
  messages: Message[]
  config: SessionConfig
  context: ProjectContext
  createdAt: Date
  updatedAt: Date
}

Message {
  id: string
  role: 'user' | 'assistant' | 'system'
  content: string | ToolCall[]
  timestamp: Date
  metadata: MessageMetadata
}
```

### **Configuration Hierarchy**

```
1. Default Config (built-in)
   ↓
2. Global Config (~/.config/kuuzuki/config.json)
   ↓
3. Project Config (./kuuzuki.json)
   ↓
4. Environment Variables
   ↓
5. CLI Arguments
```

### **Provider Interface**

```typescript
interface Provider {
  name: string
  models: Model[]
  authenticate(credentials: Credentials): Promise<void>
  generateText(request: GenerateRequest): AsyncIterable<TextChunk>
  generateObject(request: ObjectRequest): Promise<Object>
}
```

## 🔌 Integration Architecture

### **Language Server Protocol (LSP)**

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Kuuzuki TUI   │───▶│  LSP Manager    │───▶│  Language       │
│                 │    │                 │    │  Servers        │
│ • Code context  │    │ • Auto-detect   │    │                 │
│ • Diagnostics   │    │ • Lifecycle     │    │ • TypeScript    │
│ • Hover info    │    │ • Communication │    │ • Go            │
│                 │    │                 │    │ • Python        │
└─────────────────┘    └─────────────────┘    └─────────────────┘
```

**Benefits:**
- **Better Code Understanding**: AI gets type information, diagnostics
- **Accurate Suggestions**: Context-aware recommendations
- **Multi-Language Support**: Works with any LSP-compatible language

### **Git Integration**

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   File Changes  │───▶│  Git Watcher    │───▶│  Context        │
│                 │    │                 │    │  Updates        │
│ • Modified      │    │ • Detect changes│    │                 │
│ • Staged        │    │ • Track status  │    │ • Diff context  │
│ • Untracked     │    │ • Branch info   │    │ • Commit history│
└─────────────────┘    └─────────────────┘    └─────────────────┘
```

## 🚀 Performance Architecture

### **Streaming Strategy**

```
Provider Response → Server Processing → SSE Stream → TUI Updates
     (chunks)         (transform)        (real-time)    (incremental)
```

**Benefits:**
- **Immediate Feedback**: Users see responses as they're generated
- **Better UX**: No waiting for complete responses
- **Efficient**: Lower memory usage, faster perceived performance

### **Caching Strategy**

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   File System  │───▶│  Memory Cache   │───▶│  Provider API   │
│                 │    │                 │    │                 │
│ • File content  │    │ • Recent files  │    │ • AI responses  │
│ • LSP data      │    │ • Session data  │    │ • Model info    │
│ • Git status    │    │ • Config cache  │    │                 │
└─────────────────┘    └─────────────────┘    └─────────────────┘
```

## 🔒 Security Architecture

### **Authentication Flow**

```
1. User initiates auth → 2. OAuth redirect → 3. Provider auth
                                                      ↓
6. Store tokens ← 5. Exchange code ← 4. Auth callback
                                                      ↓
7. API requests with tokens → 8. Provider API calls
```

### **Security Layers**

- **Token Management**: Secure storage of API keys and OAuth tokens
- **Request Validation**: Input sanitization and validation
- **File System Isolation**: Restricted file access within project boundaries
- **Provider Isolation**: Separate credentials for each provider

## 🎨 UI Architecture (TUI)

### **Bubble Tea Pattern**

```go
type Model struct {
    state AppState
    components map[string]Component
    layout Layout
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    // Handle events and update state
}

func (m Model) View() string {
    // Render current state to terminal
}
```

### **Component Hierarchy**

```
App
├── Header (status, title)
├── Main
│   ├── Chat (messages, input)
│   ├── FileViewer (tree, content)
│   └── Modal (dialogs, overlays)
└── Footer (shortcuts, status)
```

## 🔧 Extension Architecture

### **Tool System**

```typescript
interface Tool {
  name: string
  description: string
  parameters: ToolParameters
  execute(params: any): Promise<ToolResult>
}

// Built-in tools
const tools = [
  new BashTool(),      // Shell command execution
  new EditTool(),      // File editing
  new ReadTool(),      // File reading
  new SearchTool(),    // Code search
  // ... more tools
]
```

### **Provider Plugin System**

```typescript
// Easy to add new providers
class CustomProvider implements Provider {
  async generateText(request: GenerateRequest) {
    // Custom implementation
  }
}

// Register with provider manager
providerManager.register('custom', new CustomProvider())
```

## 📊 Monitoring Architecture

### **Logging Strategy**

```
Application Logs → Structured JSON → File/Console Output
     ↓                    ↓                ↓
Error Tracking    Performance Metrics   Debug Info
```

### **Metrics Collection**

- **Request Latency**: API response times
- **Provider Performance**: AI model response times
- **Error Rates**: Failed requests and their causes
- **Usage Patterns**: Most used features and commands

---

**Next:** [Codebase Tour](CODEBASE_TOUR.md) - Explore the actual code structure and key files!