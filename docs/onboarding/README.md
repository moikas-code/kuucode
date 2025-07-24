# Welcome to Kuuzuki! 🚀

This onboarding documentation will help you understand the Kuuzuki project architecture, codebase, and development workflow.

## 📚 Documentation Overview

This folder contains everything you need to get up to speed with Kuuzuki development:

### 🎯 Start Here (Recommended Order)

1. **[Project Overview](PROJECT_OVERVIEW.md)** - What Kuuzuki is and does
2. **[Development Setup](DEVELOPMENT_SETUP.md)** - Get your environment ready
3. **[Architecture](ARCHITECTURE.md)** - Understand the system design
4. **[Codebase Tour](CODEBASE_TOUR.md)** - Guided walkthrough of the code

### 🔍 Deep Dive Topics

5. **[Key Concepts](KEY_CONCEPTS.md)** - Important patterns and ideas
6. **[API Flow](API_FLOW.md)** - How requests flow through the system
7. **[TUI Guide](TUI_GUIDE.md)** - Understanding the terminal interface
8. **[Provider System](PROVIDER_SYSTEM.md)** - How AI providers work

### 🛠️ Development Resources

9. **[Session Management](SESSION_MANAGEMENT.md)** - Data handling and persistence
10. **[Configuration](CONFIGURATION.md)** - Config system deep dive
11. **[Debugging](DEBUGGING.md)** - Troubleshooting and development tools
12. **[Common Tasks](COMMON_TASKS.md)** - Frequent development patterns

### 📖 Reference

13. **[Glossary](GLOSSARY.md)** - Terms and definitions

## 🎓 Learning Paths

### **I'm New to the Project**
→ Start with Project Overview → Development Setup → Architecture

### **I Want to Add Features**
→ Codebase Tour → Key Concepts → API Flow → Common Tasks

### **I'm Debugging Issues**
→ Debugging → Architecture → Glossary

### **I Want to Understand the TUI**
→ TUI Guide → Architecture → Key Concepts

### **I'm Working on AI Integration**
→ Provider System → API Flow → Session Management

## 🚀 Quick Start Checklist

Before diving into the docs, make sure you have:

- [ ] **Cloned the repository**: `git clone https://github.com/moikas-code/kuuzuki.git`
- [ ] **Installed dependencies**: `bun install`
- [ ] **Started the server**: `bun run dev`
- [ ] **Tested the TUI**: `cd packages/tui && go run .`

## 💡 How to Use This Documentation

### **Visual Learner?**
- Look for 📊 diagrams in Architecture and API Flow
- Check out the code examples in Codebase Tour
- Follow the visual guides in TUI Guide

### **Hands-On Learner?**
- Start with Development Setup
- Try the examples in Common Tasks
- Experiment with the code as you read

### **Conceptual Learner?**
- Begin with Project Overview and Key Concepts
- Read Architecture for the big picture
- Use Glossary for terminology

## 🆘 Getting Help

If you get stuck:

1. **Check the Debugging guide** for common issues
2. **Search the Glossary** for unfamiliar terms
3. **Look at Common Tasks** for development patterns
4. **Create an issue** at [moikas-code/kuuzuki/issues](https://github.com/moikas-code/kuuzuki/issues)

## 📝 Contributing to Documentation

Found something unclear or missing? Please:

1. **Open an issue** describing what's confusing
2. **Submit a PR** with improvements
3. **Ask questions** - they help improve the docs for everyone

## 🎯 Your Journey

```
Week 1: Foundation
├── Project Overview (understand what you're building)
├── Development Setup (get environment working)
└── Architecture (see the big picture)

Week 2: Implementation
├── Codebase Tour (learn the code structure)
├── Key Concepts (understand core patterns)
└── API Flow (trace request handling)

Week 3: Specialization
├── Choose your focus area (TUI, Providers, etc.)
├── Deep dive into relevant guides
└── Start contributing!
```

## 🌟 What Makes Kuuzuki Special

As you learn about Kuuzuki, you'll discover:

- **Clean Architecture**: Separation between TUI, server, and providers
- **Modern Stack**: TypeScript server, Go TUI, OpenAPI Generator
- **Extensible Design**: Easy to add new providers and features
- **Developer-Friendly**: Great tooling and development experience
- **Independent**: No external service dependencies

---

**Ready to dive in?** Start with [Project Overview](PROJECT_OVERVIEW.md) to understand what you're building!

*Happy coding! 🎉*