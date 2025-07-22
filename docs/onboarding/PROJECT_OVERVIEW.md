# Project Overview

## What is Kuucode?

Kuucode is an **AI coding agent built for the terminal**. It's a fork of the original OpenCode project, redesigned to be independent, cost-effective, and fully customizable.

## 🎯 Core Purpose

Kuucode helps developers by:
- **Writing and editing code** based on natural language instructions
- **Understanding codebases** through intelligent analysis
- **Debugging issues** with AI-powered insights
- **Explaining code** in plain English
- **Refactoring** and improving existing code

## 🏗️ High-Level Architecture

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Terminal UI   │───▶│  Kuucode Server │───▶│  AI Providers   │
│     (Go TUI)    │    │   (TypeScript)  │    │ (Claude, GPT,   │
│                 │    │                 │    │  Local models)  │
│ • Chat interface│    │ • API endpoints │    │                 │
│ • File browser  │    │ • Business logic│    │ • Text generation│
│ • Code display  │    │ • Provider mgmt │    │ • Code analysis │
└─────────────────┘    └─────────────────┘    └─────────────────┘
```

## 🚀 Key Features

### **Intelligent Code Assistance**
- Natural language to code conversion
- Code explanation and documentation
- Bug detection and fixing suggestions
- Refactoring recommendations

### **Multi-Provider Support**
- **Anthropic Claude** (recommended)
- **OpenAI GPT models**
- **Local models** via Ollama
- **75+ providers** through Models.dev

### **Terminal-First Experience**
- **Native TUI** built with Go and Bubble Tea
- **Keyboard shortcuts** for power users
- **Themeable interface** with multiple color schemes
- **Fast and responsive** - no web browser needed

### **Project-Aware**
- **Automatic LSP integration** for better code understanding
- **Git repository awareness**
- **File watching** and change detection
- **Context-aware suggestions**

### **Session Management**
- **Persistent conversations** across restarts
- **Shareable sessions** via web links
- **Message history** and search
- **Multiple concurrent sessions**

## 🎨 What Makes Kuucode Different

### **vs. GitHub Copilot**
- **Terminal-based** instead of editor-integrated
- **Conversational** rather than just autocomplete
- **Multi-provider** instead of OpenAI-only
- **Open source** and self-hostable

### **vs. ChatGPT/Claude Web**
- **Code-specialized** with LSP integration
- **Project context** awareness
- **Terminal workflow** integration
- **Offline capable** with local models

### **vs. Original OpenCode**
- **Independent** - no external service dependencies
- **Cost-effective** - free SDK generation (was $99/month)
- **Customizable** - full control over features and providers
- **Community-driven** - open to contributions

## 🛠️ Technology Stack

### **Frontend (TUI)**
- **Language**: Go 1.24+
- **Framework**: Bubble Tea (terminal UI)
- **Features**: Real-time updates, keyboard navigation, theming

### **Backend (Server)**
- **Language**: TypeScript
- **Runtime**: Bun (fast JavaScript runtime)
- **Framework**: Hono (lightweight web framework)
- **Features**: REST API, streaming responses, provider management

### **AI Integration**
- **SDK**: AI SDK (universal provider interface)
- **Providers**: 75+ supported through Models.dev
- **Streaming**: Real-time response streaming
- **Context**: Intelligent context management

### **Development Tools**
- **SDK Generation**: OpenAPI Generator (free alternative to Stainless)
- **Package Manager**: Bun for speed and simplicity
- **Documentation**: Astro-based website
- **Testing**: Bun test runner

## 📊 Project Structure

```
kuucode/
├── packages/
│   ├── kuucode/          # TypeScript server
│   │   ├── src/
│   │   │   ├── server/   # API endpoints
│   │   │   ├── provider/ # AI provider integrations
│   │   │   ├── auth/     # Authentication
│   │   │   └── config/   # Configuration management
│   │   └── bin/          # CLI binaries
│   ├── tui/              # Go terminal interface
│   │   ├── internal/     # TUI components
│   │   ├── sdk/          # Generated Go SDK
│   │   └── main.go       # TUI entry point
│   └── web/              # Documentation website
├── sdks/                 # Generated client SDKs
│   ├── typescript/       # TypeScript SDK
│   └── python/           # Python SDK
├── scripts/              # Build and development scripts
└── docs/                 # Documentation
```

## 🎯 Target Users

### **Primary Users**
- **Terminal-loving developers** who prefer CLI tools
- **Privacy-conscious users** who want local/self-hosted options
- **Multi-language developers** who work across different ecosystems
- **Power users** who want customizable AI coding assistance

### **Use Cases**
- **Daily coding assistance** - writing, debugging, explaining code
- **Learning new languages/frameworks** - AI-guided exploration
- **Code review and refactoring** - improving existing codebases
- **Documentation generation** - explaining complex code
- **Rapid prototyping** - quickly building proof-of-concepts

## 🌟 Project Vision

### **Short Term (3-6 months)**
- **Stable core features** - reliable coding assistance
- **Provider ecosystem** - support for major AI providers
- **Performance optimization** - fast, responsive experience
- **Documentation** - comprehensive guides and examples

### **Medium Term (6-12 months)**
- **Plugin system** - extensible architecture
- **Advanced features** - code generation templates, project scaffolding
- **Team features** - shared sessions, collaboration tools
- **Integration ecosystem** - VS Code extension, GitHub Actions

### **Long Term (1+ years)**
- **AI model training** - custom models for specific domains
- **Enterprise features** - SSO, audit logs, compliance
- **Mobile companion** - remote access to coding sessions
- **Community marketplace** - shared prompts, templates, plugins

## 🔄 Project History

### **Fork Context**
Kuucode is a fork of OpenCode (by SST), created to:
- **Remove external dependencies** (Stainless SDK generation)
- **Reduce costs** (eliminate $99/month subscription)
- **Enable customization** (full control over features)
- **Foster independence** (no reliance on external services)

### **Key Changes Made**
- **Rebranded** from OpenCode to Kuucode
- **Migrated** from Stainless to OpenAPI Generator
- **Updated** repository references to moikas-code organization
- **Improved** documentation and onboarding
- **Enhanced** development workflow

### **Why Fork in Beta?**
- **Clean codebase** - less technical debt
- **Architectural flexibility** - easier to modify core patterns
- **Learning opportunity** - understand AI agent development
- **Innovation space** - try different approaches

## 🎮 Getting Started

### **Quick Demo**
1. **Start the server**: `bun run dev`
2. **Open the TUI**: `cd packages/tui && go run .`
3. **Ask a question**: "Explain how authentication works in this codebase"
4. **Watch the magic**: AI analyzes your code and provides insights

### **Next Steps**
- **[Development Setup](DEVELOPMENT_SETUP.md)** - Get your environment ready
- **[Architecture](ARCHITECTURE.md)** - Understand the system design
- **[Codebase Tour](CODEBASE_TOUR.md)** - Explore the code structure

## 🤝 Community

### **Contributing**
- **Bug reports** - Help improve stability
- **Feature requests** - Share your ideas
- **Code contributions** - Add new capabilities
- **Documentation** - Help others learn

### **Communication**
- **GitHub Issues** - Bug reports and feature requests
- **Discussions** - General questions and ideas
- **Pull Requests** - Code contributions

---

**Ready to dive deeper?** Continue with [Development Setup](DEVELOPMENT_SETUP.md) to get your environment ready for development!