# Key Concepts

This document explains the core concepts and patterns you need to understand to work effectively with Kuuzuki.

## 🎯 Core Concepts

### **1. Sessions**

A **Session** is a conversation thread between the user and AI, similar to a chat conversation.

```typescript
interface Session {
  id: string           // Unique identifier
  name: string         // User-friendly name
  messages: Message[]  // Conversation history
  config: SessionConfig // Provider, model settings
  context: ProjectContext // File/project awareness
  createdAt: Date
  updatedAt: Date
}
```

**Key Properties:**
- **Persistent**: Sessions survive app restarts
- **Contextual**: Aware of your project files and structure
- **Configurable**: Each session can use different AI providers/models
- **Shareable**: Can be shared via web links

**Lifecycle:**
```
Create → Add Messages → Update Context → Share/Archive → Delete
```

### **2. Messages**

**Messages** are individual exchanges within a session.

```typescript
interface Message {
  id: string
  role: 'user' | 'assistant' | 'system'
  content: string | ToolCall[]
  timestamp: Date
  metadata: MessageMetadata
}
```

**Message Types:**
- **User**: Your input/questions
- **Assistant**: AI responses
- **System**: Internal messages (tool calls, context updates)

**Content Types:**
- **Text**: Regular conversation
- **Tool Calls**: When AI uses tools (file editing, shell commands)
- **Mixed**: Combination of text and tool usage

### **3. Providers**

**Providers** are AI service integrations (Claude, GPT, local models).

```typescript
interface Provider {
  name: string
  models: Model[]
  authenticate(credentials: Credentials): Promise<void>
  generateText(request: GenerateRequest): AsyncIterable<TextChunk>
}
```

**Built-in Providers:**
- **Anthropic**: Claude models (recommended)
- **OpenAI**: GPT models
- **Google**: Gemini models
- **Local**: Ollama and other local models

**Provider Selection:**
```typescript
// Automatic selection based on config
const provider = await providerManager.select(session.config)

// Manual selection
const provider = await providerManager.get('anthropic')
```

### **4. Tools**

**Tools** are capabilities the AI can use to interact with your system.

```typescript
interface Tool {
  name: string
  description: string
  parameters: ToolParameters
  execute(params: any): Promise<ToolResult>
}
```

**Built-in Tools:**
- **edit**: Modify files with precise replacements
- **read**: Read file contents
- **bash**: Execute shell commands
- **glob**: Find files by pattern
- **grep**: Search text in files

**Tool Execution Flow:**
```
AI decides to use tool → Tool parameters generated → Tool executed → Results returned to AI
```

### **5. Context**

**Context** is the project awareness that makes Kuuzuki intelligent about your codebase.

```typescript
interface ProjectContext {
  workingDirectory: string
  gitRepository?: GitInfo
  openFiles: string[]
  lspData: LSPContext
  recentChanges: FileChange[]
}
```

**Context Sources:**
- **File System**: Current directory structure
- **Git**: Repository status, recent commits
- **LSP**: Language server data (types, diagnostics)
- **User Activity**: Recently opened/edited files

## 🔄 Core Patterns

### **1. Streaming Responses**

Kuuzuki uses **streaming** for real-time AI responses.

```typescript
// Server: Stream AI response
return streamText({
  model: provider.model,
  messages: conversation,
  onChunk: (chunk) => {
    eventStream.send(sessionId, {
      type: 'text-chunk',
      content: chunk.text
    })
  }
})

// TUI: Handle streaming updates
func (m Model) handleStreamEvent(event StreamEvent) {
  switch event.Type {
  case "text-chunk":
    m.currentMessage += event.Content
    // Update UI immediately
  }
}
```

**Benefits:**
- **Immediate feedback**: See responses as they're generated
- **Better UX**: No waiting for complete responses
- **Cancellation**: Can stop generation mid-stream

### **2. Event-Driven Architecture**

The TUI uses **events** to handle user input and server responses.

```go
// Bubble Tea event handling pattern
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  switch msg := msg.(type) {
  case tea.KeyMsg:
    return m.handleKeyPress(msg)
  case StreamEvent:
    return m.handleStreamEvent(msg)
  case tea.WindowSizeMsg:
    return m.handleResize(msg)
  }
  return m, nil
}
```

**Event Types:**
- **Keyboard**: User input
- **Stream**: AI responses
- **Window**: Terminal resize
- **Timer**: Periodic updates

### **3. Configuration Hierarchy**

Configuration follows a **hierarchy** from general to specific.

```
1. Built-in defaults
   ↓
2. Global config (~/.config/kuuzuki/)
   ↓  
3. Project config (./kuuzuki.json)
   ↓
4. Session config (per-session settings)
   ↓
5. Runtime overrides (CLI args, env vars)
```

**Example:**
```json
// Global config
{
  "provider": {
    "anthropic": { "apiKey": "..." }
  }
}

// Project config  
{
  "provider": {
    "default": "anthropic",
    "anthropic": { "model": "claude-3-sonnet" }
  }
}

// Session config
{
  "provider": "openai",
  "model": "gpt-4"
}
```

### **4. Component Composition**

The TUI uses **component composition** for modular UI.

```go
type Model struct {
  // Composed components
  chat       chat.Model
  fileViewer fileviewer.Model
  status     status.Model
  modal      modal.Model
  
  // Layout state
  layout     Layout
  focus      Component
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  var cmds []tea.Cmd
  
  // Update focused component
  switch m.focus {
  case ChatComponent:
    m.chat, cmd = m.chat.Update(msg)
    cmds = append(cmds, cmd)
  case FileViewerComponent:
    m.fileViewer, cmd = m.fileViewer.Update(msg)
    cmds = append(cmds, cmd)
  }
  
  return m, tea.Batch(cmds...)
}
```

## 🧠 Mental Models

### **1. Client-Server Separation**

Think of Kuuzuki as **two separate applications** that communicate:

```
TUI (Client)                    Server
├── User Interface             ├── Business Logic
├── Event Handling             ├── AI Provider Management  
├── Display Logic              ├── Session Management
└── HTTP Client                └── Tool Execution
```

**Why This Matters:**
- **TUI can be restarted** without losing sessions
- **Multiple clients** can connect to same server
- **Server can run remotely** for team usage
- **Clear separation** of concerns

### **2. Stateless Server, Stateful Client**

```
Server (Stateless)             TUI (Stateful)
├── Processes requests         ├── Maintains UI state
├── Calls AI providers         ├── Tracks user input
├── Executes tools             ├── Manages display
└── Returns responses          └── Handles navigation
```

**Implications:**
- **Server restarts** don't affect TUI state
- **Horizontal scaling** possible for server
- **TUI handles** all user interaction state
- **Server focuses** on business logic

### **3. Tool-Augmented AI**

AI doesn't just generate text - it can **use tools** to interact with your system:

```
User: "Fix the bug in auth.ts"
  ↓
AI: I'll help you fix that. Let me first read the file.
  ↓
Tool Call: read(filePath: "auth.ts")
  ↓
Tool Result: [file contents]
  ↓
AI: I see the issue. Let me fix it.
  ↓
Tool Call: edit(filePath: "auth.ts", oldString: "...", newString: "...")
  ↓
Tool Result: File updated successfully
  ↓
AI: Fixed! The issue was...
```

### **4. Context-Aware Conversations**

Unlike web-based AI chats, Kuuzuki is **aware of your project**:

```
Project Context
├── File structure
├── Git status  
├── LSP data (types, errors)
├── Recent changes
└── Open files

        ↓ (informs)

AI Understanding
├── Better suggestions
├── Accurate file paths
├── Type-aware code
└── Project-specific advice
```

## 🔧 Development Patterns

### **1. Adding New Tools**

```typescript
// 1. Define the tool
export const myTool = {
  name: "my-tool",
  description: "Does something useful",
  parameters: {
    input: { type: "string", description: "Input parameter" }
  },
  execute: async ({ input }) => {
    // Tool implementation
    return { success: true, result: "..." }
  }
}

// 2. Register the tool
import { myTool } from "./my-tool.js"
export const tools = [
  bashTool,
  editTool,
  myTool,  // Add here
  // ...
]
```

### **2. Adding New Providers**

```typescript
// 1. Implement provider interface
class MyProvider implements Provider {
  name = "my-provider"
  models = [{ name: "my-model", id: "my-model-v1" }]
  
  async authenticate(credentials: Credentials) {
    // Auth logic
  }
  
  async *generateText(request: GenerateRequest) {
    // Generation logic
    for (const chunk of response) {
      yield { text: chunk }
    }
  }
}

// 2. Register provider
providerManager.register("my-provider", new MyProvider())
```

### **3. Adding TUI Components**

```go
// 1. Create component
type MyComponent struct {
  state MyState
}

func (c MyComponent) Update(msg tea.Msg) (MyComponent, tea.Cmd) {
  // Handle events
  return c, nil
}

func (c MyComponent) View() string {
  // Render component
  return "My component view"
}

// 2. Integrate with main model
type Model struct {
  myComponent MyComponent
  // ...
}
```

### **4. Configuration Patterns**

```typescript
// 1. Define config schema
interface MyConfig {
  enabled: boolean
  options: MyOptions
}

// 2. Add to main config
interface Config {
  myFeature: MyConfig
  // ...
}

// 3. Use in code
const config = await configLoader.load()
if (config.myFeature.enabled) {
  // Feature logic
}
```

## 🎨 UI Patterns

### **1. Responsive Layout**

```go
func (m Model) View() string {
  width, height := m.size.Width, m.size.Height
  
  if width < 80 {
    // Narrow layout
    return m.renderNarrow()
  } else {
    // Wide layout  
    return m.renderWide()
  }
}
```

### **2. Modal Dialogs**

```go
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  if m.modal.IsOpen() {
    // Modal handles all input
    m.modal, cmd = m.modal.Update(msg)
    return m, cmd
  }
  
  // Normal input handling
  return m.handleNormalInput(msg)
}
```

### **3. Keyboard Navigation**

```go
func (m Model) handleKeyPress(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
  switch msg.String() {
  case "tab":
    return m.nextComponent(), nil
  case "shift+tab":
    return m.prevComponent(), nil
  case "ctrl+c":
    return m, tea.Quit
  }
  
  // Delegate to focused component
  return m.updateFocusedComponent(msg)
}
```

## 🔍 Debugging Patterns

### **1. Structured Logging**

```typescript
import { logger } from "./util/log.js"

// Structured logging with context
logger.info("Session created", {
  sessionId: session.id,
  provider: session.config.provider,
  model: session.config.model
})

logger.error("Provider error", {
  provider: "anthropic",
  error: error.message,
  requestId: request.id
})
```

### **2. Error Boundaries**

```go
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  defer func() {
    if r := recover(); r != nil {
      log.Printf("TUI panic: %v", r)
      // Show error modal
      m.modal = errorModal(fmt.Sprintf("Error: %v", r))
    }
  }()
  
  // Normal update logic
  return m.handleUpdate(msg)
}
```

### **3. Development Helpers**

```typescript
// Debug mode helpers
if (process.env.DEBUG) {
  console.log("Request:", JSON.stringify(request, null, 2))
  console.log("Response:", JSON.stringify(response, null, 2))
}

// Performance timing
const start = performance.now()
const result = await expensiveOperation()
const duration = performance.now() - start
logger.debug("Operation completed", { duration, operation: "expensiveOperation" })
```

---

**Next:** [API Flow](API_FLOW.md) - See how requests flow through the system in detail!