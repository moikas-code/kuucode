# Development Setup

This guide will get your development environment ready for working on Kuuzuki.

## 📋 Prerequisites

### **Required Tools**

| Tool | Version | Purpose | Installation |
|------|---------|---------|--------------|
| **Bun** | Latest | Package manager & runtime | `curl -fsSL https://bun.sh/install \| bash` |
| **Go** | 1.24+ | TUI development | [golang.org/dl](https://golang.org/dl/) |
| **Git** | 2.0+ | Version control | [git-scm.com](https://git-scm.com/) |
| **curl** | Any | API testing | Usually pre-installed |

### **Optional but Recommended**

| Tool | Purpose | Installation |
|------|---------|--------------|
| **jq** | JSON processing | `brew install jq` or `apt install jq` |
| **VS Code** | Code editing | [code.visualstudio.com](https://code.visualstudio.com/) |
| **Go extension** | Go development | VS Code marketplace |

## 🚀 Quick Setup

### **1. Clone and Install**

```bash
# Clone your fork
git clone https://github.com/moikas-code/kuuzuki.git
cd kuuzuki

# Install dependencies
bun install

# Verify installation
bun run --version
go version
```

### **2. Start Development Servers**

```bash
# Terminal 1: Start the TypeScript server
bun run dev

# Terminal 2: Start the Go TUI (in another terminal)
cd packages/tui
go run .
```

### **3. Install Globally (Optional)**

```bash
# Link for global usage
cd packages/kuuzuki
bun link
bun link kuuzuki

# Test global installation
kuuzuki --help
```

### **4. Verify Everything Works**

```bash
# Test the API directly
curl http://localhost:4096/app

# Test SDK generation
./scripts/generate-sdks --dev

# Test global CLI
kuuzuki serve --port 4096

# Check TUI connection
# The TUI should connect to the server automatically
```

## 🔧 Detailed Setup

### **Environment Configuration**

Create a local config file:

```bash
# Copy example config
cp kuuzuki.json kuuzuki.local.json

# Edit for your setup
{
  "$schema": "https://kuuzuki.ai/config.json",
  "provider": {
    "anthropic": {
      "apiKey": "your-api-key-here"
    }
  }
}
```

### **Development Workflow**

```bash
# Start development (recommended workflow)
# Terminal 1: Server with hot reload
bun run dev

# Terminal 2: TUI development
cd packages/tui
go run .

# Terminal 3: Commands and testing
curl http://localhost:4096/session
./scripts/generate-sdks --dev
```

## 🛠️ Component-Specific Setup

### **TypeScript Server Development**

```bash
# Navigate to server package
cd packages/kuuzuki

# Install dependencies (if not done globally)
bun install

# Start with hot reload
bun run dev

# Type checking
bun run typecheck

# Direct execution
bun run src/index.ts --help
```

**Key Files to Know:**
- `src/index.ts` - Main entry point
- `src/server/server.ts` - API routes
- `src/provider/` - AI provider integrations
- `src/config/` - Configuration management

### **Go TUI Development**

```bash
# Navigate to TUI package
cd packages/tui

# Install Go dependencies
go mod tidy

# Run the TUI
go run .

# Build binary
go build -o kuuzuki-tui .

# Run tests
go test ./...
```

**Key Files to Know:**
- `main.go` - TUI entry point
- `internal/tui/tui.go` - Main TUI logic
- `internal/components/` - UI components
- `sdk/` - Generated Go SDK for API communication

### **Documentation Development**

```bash
# Navigate to web package
cd packages/web

# Install dependencies
bun install

# Start development server
bun run dev

# Build for production
bun run build
```

## 🔍 Verification Steps

### **1. Server Health Check**

```bash
# Start server
bun run dev

# Test endpoints
curl http://localhost:4096/app
curl http://localhost:4096/config
curl http://localhost:4096/session

# Expected: JSON responses, no errors
```

### **2. TUI Connection Test**

```bash
# With server running, start TUI
cd packages/tui
go run .

# Expected: TUI loads, connects to server
# Try typing a message to test AI integration
```

### **3. SDK Generation Test**

```bash
# Generate SDKs
./scripts/generate-sdks --dev

# Check outputs
ls -la sdks/typescript/
ls -la sdks/python/
ls -la packages/tui/sdk/

# Expected: Generated SDK files in each directory
```

## 🐛 Common Issues and Solutions

### **Issue: Bun not found**
```bash
# Solution: Install Bun
curl -fsSL https://bun.sh/install | bash
source ~/.bashrc  # or restart terminal
```

### **Issue: Go version too old**
```bash
# Check version
go version

# Solution: Update Go
# Visit https://golang.org/dl/ for latest version
```

### **Issue: Port 4096 already in use**
```bash
# Check what's using the port
lsof -i :4096

# Kill the process or use different port
kill -9 <PID>
# or
PORT=4097 bun run dev
```

### **Issue: TUI won't connect to server**
```bash
# Verify server is running
curl http://localhost:4096/app

# Check TUI logs
cd packages/tui
go run . --debug

# Common fix: Restart both server and TUI
```

### **Issue: SDK generation fails**
```bash
# Install OpenAPI Generator
npm install -g @openapitools/openapi-generator-cli

# Verify installation
openapi-generator-cli version

# Try generation again
./scripts/generate-sdks --dev
```

## 🎯 Development Environment Tips

### **Recommended VS Code Extensions**

```json
{
  "recommendations": [
    "golang.go",
    "ms-vscode.vscode-typescript-next",
    "bradlc.vscode-tailwindcss",
    "ms-vscode.vscode-json"
  ]
}
```

### **Useful Aliases**

Add to your `~/.bashrc` or `~/.zshrc`:

```bash
# Kuuzuki development aliases
alias kdev="cd ~/path/to/kuuzuki && bun run dev"
alias ktui="cd ~/path/to/kuuzuki/packages/tui && go run ."
alias ksdk="cd ~/path/to/kuuzuki && ./scripts/generate-sdks --dev"
alias ktest="cd ~/path/to/kuuzuki && bun run typecheck && cd packages/tui && go test ./..."
```

### **Git Configuration**

```bash
# Set up Git hooks (already done by postinstall)
./scripts/hooks

# Recommended Git config
git config --local core.autocrlf false
git config --local pull.rebase true
```

## 📊 Development Workflow

### **Daily Development**

```bash
# 1. Start your day
git pull origin main
bun install  # if package.json changed

# 2. Start development servers
bun run dev          # Terminal 1
cd packages/tui && go run .  # Terminal 2

# 3. Make changes
# Edit files in your preferred editor

# 4. Test changes
# TUI automatically reflects server changes
# Restart TUI if you change Go code

# 5. Generate SDKs (if API changed)
./scripts/generate-sdks --dev

# 6. Commit changes
git add .
git commit -m "Your descriptive commit message"
```

### **Feature Development**

```bash
# 1. Create feature branch
git checkout -b feature/your-feature-name

# 2. Develop and test
# ... make your changes ...

# 3. Ensure everything works
bun run typecheck
cd packages/tui && go test ./...
./scripts/generate-sdks --dev

# 4. Commit and push
git add .
git commit -m "Add your feature description"
git push -u origin feature/your-feature-name
```

## 🎉 You're Ready!

If you've completed this setup successfully, you should have:

- ✅ **Server running** on http://localhost:4096
- ✅ **TUI connecting** to the server
- ✅ **SDK generation** working
- ✅ **Development tools** configured

**Next Steps:**
- **[Architecture](ARCHITECTURE.md)** - Understand how everything fits together
- **[Codebase Tour](CODEBASE_TOUR.md)** - Explore the code structure
- **[Common Tasks](COMMON_TASKS.md)** - Learn frequent development patterns

---

**Having issues?** Check the [Debugging Guide](DEBUGGING.md) or create an issue on GitHub!