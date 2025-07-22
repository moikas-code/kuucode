# Common Development Tasks

This guide covers the most frequent development tasks you'll encounter when working on Kuucode.

## 🚀 Getting Started Tasks

### **Start Development Environment**

```bash
# Option 1: Using global kuucode command
kuucode serve

# Option 2: Using development script
cd kuucode
bun run dev

# Terminal 2: Start the TUI  
cd packages/tui
go run .

# Terminal 3: Development commands
# Use this for testing, SDK generation, etc.
```

### **Test Your Changes**

```bash
# Type check TypeScript
bun run typecheck

# Test Go code
cd packages/tui && go test ./...

# Test API endpoints
curl http://localhost:4096/app
curl http://localhost:4096/session

# Generate and test SDKs
./scripts/generate-sdks --dev
```

## 🔧 API Development

### **Add a New API Endpoint**

1. **Define the endpoint** in `packages/kuucode/src/server/server.ts`:

```typescript
// Add to the Hono app
app.get('/my-endpoint', async (c) => {
  const result = await myBusinessLogic()
  return c.json(result)
})

app.post('/my-endpoint', async (c) => {
  const body = await c.req.json()
  const result = await processData(body)
  return c.json(result)
})
```

2. **Add types** if needed in `packages/kuucode/src/types.ts`:

```typescript
export interface MyEndpointRequest {
  data: string
  options?: MyOptions
}

export interface MyEndpointResponse {
  success: boolean
  result: string
}
```

3. **Regenerate SDKs**:

```bash
./scripts/generate-sdks --dev
```

4. **Test the endpoint**:

```bash
# GET request
curl http://localhost:4096/my-endpoint

# POST request
curl -X POST http://localhost:4096/my-endpoint \
  -H "Content-Type: application/json" \
  -d '{"data": "test"}'
```

### **Add Request Validation**

```typescript
import { z } from 'zod'

const MyRequestSchema = z.object({
  data: z.string(),
  options: z.object({
    flag: z.boolean().optional()
  }).optional()
})

app.post('/my-endpoint', async (c) => {
  const body = await c.req.json()
  
  // Validate request
  const parsed = MyRequestSchema.parse(body)
  
  const result = await processData(parsed)
  return c.json(result)
})
```

### **Add Streaming Endpoint**

```typescript
import { streamText } from 'ai'

app.post('/stream-endpoint', async (c) => {
  const { prompt } = await c.req.json()
  
  return streamText({
    model: provider.model,
    prompt: prompt,
    onChunk: (chunk) => {
      // Stream to client via SSE
      eventStream.send(sessionId, {
        type: 'chunk',
        content: chunk.text
      })
    }
  })
})
```

## 🛠️ Tool Development

### **Create a New Tool**

1. **Create tool file** `packages/kuucode/src/tool/my-tool.ts`:

```typescript
export const myTool = {
  name: "my-tool",
  description: "Does something useful with files",
  parameters: {
    filePath: {
      type: "string",
      description: "Path to the file to process"
    },
    options: {
      type: "object",
      properties: {
        recursive: { type: "boolean" },
        pattern: { type: "string" }
      }
    }
  },
  execute: async ({ filePath, options = {} }) => {
    try {
      // Tool implementation
      const result = await processFile(filePath, options)
      
      return {
        success: true,
        result: result,
        message: `Processed ${filePath} successfully`
      }
    } catch (error) {
      return {
        success: false,
        error: error.message
      }
    }
  }
}
```

2. **Register the tool** in `packages/kuucode/src/tool/index.ts`:

```typescript
import { myTool } from "./my-tool.js"

export const tools = [
  bashTool,
  editTool,
  readTool,
  myTool,  // Add your tool here
  // ... other tools
]
```

3. **Test the tool**:

```bash
# Regenerate SDKs to include new tool
./scripts/generate-sdks --dev

# Test via TUI or API
# The AI can now use your tool automatically
```

### **Tool Best Practices**

```typescript
export const goodTool = {
  name: "good-tool",
  description: "Clear, specific description of what this tool does",
  parameters: {
    // Required parameters first
    requiredParam: {
      type: "string",
      description: "Clear description of what this parameter does"
    },
    // Optional parameters with defaults
    optionalParam: {
      type: "boolean",
      description: "Optional parameter with clear purpose",
      default: false
    }
  },
  execute: async (params) => {
    // Validate inputs
    if (!params.requiredParam) {
      return { success: false, error: "requiredParam is required" }
    }
    
    try {
      // Do the work
      const result = await doWork(params)
      
      // Return structured result
      return {
        success: true,
        result: result,
        message: "Operation completed successfully"
      }
    } catch (error) {
      // Handle errors gracefully
      return {
        success: false,
        error: error.message,
        details: error.stack
      }
    }
  }
}
```

## 🎨 TUI Development

### **Add a New TUI Component**

1. **Create component** `packages/tui/internal/components/mycomponent/mycomponent.go`:

```go
package mycomponent

import (
    "github.com/charmbracelet/bubbles/v2/key"
    tea "github.com/charmbracelet/bubbletea/v2"
    "github.com/charmbracelet/lipgloss/v2"
)

type Model struct {
    state  MyState
    width  int
    height int
}

type MyState struct {
    data   []string
    cursor int
}

func New() Model {
    return Model{
        state: MyState{
            data: []string{"Item 1", "Item 2", "Item 3"},
        },
    }
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "up", "k":
            if m.state.cursor > 0 {
                m.state.cursor--
            }
        case "down", "j":
            if m.state.cursor < len(m.state.data)-1 {
                m.state.cursor++
            }
        case "enter":
            // Handle selection
            return m, m.handleSelection()
        }
    case tea.WindowSizeMsg:
        m.width = msg.Width
        m.height = msg.Height
    }
    
    return m, nil
}

func (m Model) View() string {
    var items []string
    
    for i, item := range m.state.data {
        cursor := " "
        if i == m.state.cursor {
            cursor = ">"
        }
        
        style := lipgloss.NewStyle()
        if i == m.state.cursor {
            style = style.Foreground(lipgloss.Color("205"))
        }
        
        items = append(items, fmt.Sprintf("%s %s", cursor, style.Render(item)))
    }
    
    return lipgloss.JoinVertical(lipgloss.Left, items...)
}

func (m Model) handleSelection() tea.Cmd {
    selected := m.state.data[m.state.cursor]
    // Do something with selection
    return nil
}
```

2. **Integrate with main TUI** in `packages/tui/internal/tui/tui.go`:

```go
import "github.com/moikas-code/kuucode/internal/components/mycomponent"

type Model struct {
    // ... existing components
    myComponent mycomponent.Model
    
    // ... rest of model
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmds []tea.Cmd
    
    // Update your component
    m.myComponent, cmd = m.myComponent.Update(msg)
    cmds = append(cmds, cmd)
    
    // ... handle other components
    
    return m, tea.Batch(cmds...)
}

func (m Model) View() string {
    // Include your component in the layout
    return lipgloss.JoinHorizontal(
        lipgloss.Top,
        m.chat.View(),
        m.myComponent.View(),
    )
}
```

### **Handle Keyboard Shortcuts**

```go
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+n":
            return m, m.createNewSession()
        case "ctrl+o":
            return m, m.openFile()
        case "ctrl+s":
            return m, m.saveSession()
        case "tab":
            return m.nextComponent(), nil
        case "shift+tab":
            return m.prevComponent(), nil
        }
    }
    
    return m, nil
}
```

### **Add Modal Dialogs**

```go
// In your component
func (m Model) showConfirmDialog(message string) tea.Cmd {
    return func() tea.Msg {
        return ShowModalMsg{
            Type: "confirm",
            Title: "Confirm Action",
            Message: message,
            Callback: func(confirmed bool) tea.Cmd {
                if confirmed {
                    return m.performAction()
                }
                return nil
            },
        }
    }
}
```

## 🔌 Provider Development

### **Add a New AI Provider**

1. **Create provider** `packages/kuucode/src/provider/my-provider.ts`:

```typescript
import { Provider, GenerateRequest, Model } from './types.js'

export class MyProvider implements Provider {
  name = 'my-provider'
  models: Model[] = [
    { id: 'my-model-v1', name: 'My Model v1' },
    { id: 'my-model-v2', name: 'My Model v2' }
  ]
  
  private apiKey?: string
  
  async authenticate(credentials: { apiKey: string }) {
    this.apiKey = credentials.apiKey
    
    // Test authentication
    const response = await fetch('https://api.myprovider.com/auth/test', {
      headers: { 'Authorization': `Bearer ${this.apiKey}` }
    })
    
    if (!response.ok) {
      throw new Error('Authentication failed')
    }
  }
  
  async *generateText(request: GenerateRequest) {
    const response = await fetch('https://api.myprovider.com/generate', {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${this.apiKey}`,
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        model: request.model,
        messages: request.messages,
        stream: true
      })
    })
    
    const reader = response.body?.getReader()
    if (!reader) throw new Error('No response body')
    
    while (true) {
      const { done, value } = await reader.read()
      if (done) break
      
      const chunk = new TextDecoder().decode(value)
      const lines = chunk.split('\n').filter(line => line.trim())
      
      for (const line of lines) {
        if (line.startsWith('data: ')) {
          const data = JSON.parse(line.slice(6))
          if (data.text) {
            yield { text: data.text }
          }
        }
      }
    }
  }
}
```

2. **Register provider** in `packages/kuucode/src/provider/manager.ts`:

```typescript
import { MyProvider } from './my-provider.js'

export class ProviderManager {
  private providers = new Map<string, Provider>()
  
  constructor() {
    // Register built-in providers
    this.register('anthropic', new AnthropicProvider())
    this.register('openai', new OpenAIProvider())
    this.register('my-provider', new MyProvider())  // Add here
  }
}
```

3. **Add configuration schema**:

```typescript
// In config schema
interface ProviderConfig {
  anthropic?: AnthropicConfig
  openai?: OpenAIConfig
  'my-provider'?: MyProviderConfig  // Add config
}

interface MyProviderConfig {
  apiKey: string
  model?: string
  baseUrl?: string
}
```

## 📊 Configuration Tasks

### **Add New Configuration Option**

1. **Update config types** in `packages/kuucode/src/config/types.ts`:

```typescript
export interface Config {
  // ... existing config
  myFeature: MyFeatureConfig
}

export interface MyFeatureConfig {
  enabled: boolean
  options: {
    setting1: string
    setting2: number
  }
}
```

2. **Add default values** in `packages/kuucode/src/config/defaults.ts`:

```typescript
export const defaultConfig: Config = {
  // ... existing defaults
  myFeature: {
    enabled: false,
    options: {
      setting1: 'default-value',
      setting2: 42
    }
  }
}
```

3. **Add validation** in `packages/kuucode/src/config/schema.ts`:

```typescript
import { z } from 'zod'

const MyFeatureConfigSchema = z.object({
  enabled: z.boolean(),
  options: z.object({
    setting1: z.string(),
    setting2: z.number().min(0)
  })
})

export const ConfigSchema = z.object({
  // ... existing schema
  myFeature: MyFeatureConfigSchema
})
```

4. **Use in code**:

```typescript
const config = await configLoader.load()

if (config.myFeature.enabled) {
  const result = await processWithOptions(config.myFeature.options)
}
```

## 🧪 Testing Tasks

### **Test API Endpoints**

```bash
# Start server
bun run dev

# Test endpoints with curl
curl -X GET http://localhost:4096/app
curl -X POST http://localhost:4096/session \
  -H "Content-Type: application/json" \
  -d '{"name": "Test Session"}'

# Test with authentication
curl -X GET http://localhost:4096/session \
  -H "Authorization: Bearer your-token"
```

### **Test TUI Components**

```go
// Create test file: packages/tui/internal/components/mycomponent/mycomponent_test.go
package mycomponent

import (
    "testing"
    tea "github.com/charmbracelet/bubbletea/v2"
)

func TestMyComponent(t *testing.T) {
    model := New()
    
    // Test initial state
    if model.state.cursor != 0 {
        t.Errorf("Expected cursor to be 0, got %d", model.state.cursor)
    }
    
    // Test key handling
    model, _ = model.Update(tea.KeyMsg{Type: tea.KeyDown})
    if model.state.cursor != 1 {
        t.Errorf("Expected cursor to be 1 after down key, got %d", model.state.cursor)
    }
}
```

### **Test Tools**

```typescript
// Test your tool directly
import { myTool } from '../src/tool/my-tool.js'

const result = await myTool.execute({
  filePath: 'test-file.txt',
  options: { recursive: true }
})

console.log('Tool result:', result)
```

## 🚀 Build and Release Tasks

### **Generate SDKs**

```bash
# Development (faster, local only)
./scripts/generate-sdks --dev

# Production (full generation)
./scripts/generate-sdks

# Clean build (remove previous)
./scripts/generate-sdks --clean

# Check generated files
ls -la sdks/typescript/
ls -la sdks/python/
ls -la packages/tui/sdk/
```

### **Build TUI Binary**

```bash
cd packages/tui

# Build for current platform
go build -o kuucode-tui .

# Build for multiple platforms
GOOS=linux GOARCH=amd64 go build -o kuucode-tui-linux .
GOOS=darwin GOARCH=amd64 go build -o kuucode-tui-darwin .
GOOS=windows GOARCH=amd64 go build -o kuucode-tui-windows.exe .
```

### **Type Check Everything**

```bash
# Check TypeScript
bun run typecheck

# Check Go
cd packages/tui && go vet ./...

# Check formatting
cd packages/tui && go fmt ./...
```

## 🐛 Debugging Tasks

### **Debug Server Issues**

```bash
# Start with debug logging
DEBUG=1 bun run dev

# Check specific endpoints
curl -v http://localhost:4096/app

# Monitor logs
tail -f ~/.local/share/kuucode/log/*.log
```

### **Debug TUI Issues**

```bash
# Run with debug output
cd packages/tui
go run . --debug

# Check for panics
go run . 2>&1 | grep -i panic

# Use Go debugger
dlv debug .
```

### **Debug SDK Generation**

```bash
# Validate OpenAPI spec
jq empty openapi.json

# Test OpenAPI Generator
openapi-generator-cli validate -i openapi.json

# Generate with verbose output
DEBUG=1 ./scripts/generate-sdks --dev
```

---

**Need help with a specific task?** Check the other onboarding docs or create an issue on GitHub!