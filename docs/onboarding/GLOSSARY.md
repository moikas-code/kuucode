# Glossary

This glossary defines key terms and concepts used throughout the Kuucode project.

## 🎯 Core Terms

### **Agent**

An AI-powered assistant that can understand natural language instructions and perform coding tasks. Kuucode is an AI coding agent.

### **API**

Application Programming Interface. In Kuucode, refers to the REST API provided by the TypeScript server that clients (like the TUI) use to communicate.

### **Bubble Tea**

A Go framework for building terminal user interfaces. Used by Kuucode's TUI component for creating the interactive terminal experience.

### **Bun**

A fast JavaScript runtime and package manager. Kuucode uses Bun instead of Node.js for better performance and developer experience.

### **Client**

Any application that connects to the Kuucode server. Examples include the TUI, VS Code extension, or web clients.

### **Context**

Information about the current project, files, and environment that helps the AI provide better assistance. Includes file structure, Git status, LSP data, etc.

## 🔧 Technical Terms

### **Hono**

A lightweight web framework for TypeScript/JavaScript. Kuucode uses Hono for its HTTP server instead of Express.

### **LSP (Language Server Protocol)**

A protocol that provides language features like autocomplete, diagnostics, and hover information. Kuucode integrates with LSPs to understand code better.

### **OpenAPI Generator**

An open-source tool that generates client SDKs from OpenAPI specifications. Replaced Stainless in Kuucode for cost savings and independence.

### **Provider**

An AI service integration (like Anthropic Claude, OpenAI GPT, or local models). Providers implement a common interface for text generation.

### **SDK (Software Development Kit)**

Client libraries generated from the API specification. Kuucode generates TypeScript, Go, and Python SDKs.

### **Server-Sent Events (SSE)**

A web standard for streaming real-time data from server to client. Used for streaming AI responses in real-time.

### **Stainless**

A commercial service for generating SDKs from API specifications. Previously used by Kuucode but replaced with OpenAPI Generator.

### **Streaming**

The process of sending data in chunks as it becomes available, rather than waiting for the complete response. Enables real-time AI response display.

### **TUI (Terminal User Interface)**

A text-based user interface that runs in a terminal. Kuucode's primary interface is a TUI built with Go and Bubble Tea.

## 🎨 UI Terms

### **Component**

A reusable UI element in the TUI. Examples include chat components, file viewers, and status bars.

### **Focus**

Which UI component is currently active and receiving keyboard input.

### **Layout**

How UI components are arranged on screen. Can be responsive based on terminal size.

### **Modal**

A dialog box that appears over the main interface, requiring user interaction before continuing.

### **Theme**

A set of colors and styling rules that define the visual appearance of the TUI.

### **Viewport**

The visible area of content, especially for scrollable components like chat history or file lists.

## 🔄 Workflow Terms

### **Chat**

A conversation between the user and AI within a session. Similar to a chat conversation in messaging apps.

### **Message**

A single exchange in a chat. Can be from the user, AI assistant, or system (for tool calls).

### **Session**

A persistent conversation thread with the AI. Sessions maintain context and can be resumed across app restarts.

### **Tool Call**

When the AI decides to use a tool (like editing a file or running a command) to accomplish a task.

### **Tool**

A capability that the AI can use to interact with the system. Examples include file editing, shell commands, and code search.

## 📊 Data Terms

### **Configuration**

Settings that control how Kuucode behaves. Can be global, project-specific, or session-specific.

### **Credentials**

Authentication information like API keys for AI providers.

### **Metadata**

Additional information about messages, sessions, or other data objects.

### **Persistence**

Storing data (like sessions and messages) so they survive app restarts.

### **Schema**

A definition of the structure and validation rules for data objects.

## 🔌 Integration Terms

### **Authentication**

The process of verifying identity, typically using API keys or OAuth tokens for AI providers.

### **Git Integration**

Features that make Kuucode aware of Git repository status, changes, and history.

### **Model**

A specific AI model from a provider (like Claude 3.5 Sonnet or GPT-4).

### **OAuth**

An authentication protocol used by some AI providers for secure access.

### **Project Context**

Information about the current project including file structure, Git status, and recent changes.

## 🏗️ Architecture Terms

### **Client-Server Architecture**

A design where the TUI (client) communicates with a separate server process for business logic.

### **Event-Driven**

A programming pattern where components react to events (like keyboard input or API responses).

### **Middleware**

Code that processes HTTP requests before they reach the main handler. Used for authentication, logging, etc.

### **REST API**

A web API design pattern using HTTP methods (GET, POST, etc.) for different operations.

### **Stateless**

A design where the server doesn't maintain client state between requests. Each request contains all needed information.

## 🔧 Development Terms

### **Hot Reload**

Automatically restarting the application when code changes are detected during development.

### **Monorepo**

A repository containing multiple related packages. Kuucode uses a monorepo structure with packages for server, TUI, web, etc.

### **Package Manager**

A tool for managing dependencies. Kuucode uses Bun as its package manager.

### **Type Checking**

Verifying that TypeScript code uses types correctly. Done with `bun run typecheck`.

### **Workspace**

A collection of related packages managed together. Defined in the root package.json.

## 🚀 Deployment Terms

### **Binary**

A compiled executable file. The TUI is compiled to a binary for distribution.

### **Build**

The process of compiling and preparing code for distribution.

### **CI/CD**

Continuous Integration/Continuous Deployment. Automated testing and deployment processes.

### **Release**

A versioned distribution of the software with specific features and bug fixes.

## 📝 File Types

### **`.go`**

Go source code files. Used for the TUI component.

### **`.ts`**

TypeScript source code files. Used for the server component.

### **`.json`**

JavaScript Object Notation files. Used for configuration and data.

### **`.md`**

Markdown files. Used for documentation.

### **`.yml/.yaml`**

YAML configuration files. Used for various configuration purposes.

## 🔍 Common Abbreviations

| Abbreviation | Full Term                         | Meaning                                                                        |
| ------------ | --------------------------------- | ------------------------------------------------------------------------------ |
| **AI**       | Artificial Intelligence           | Computer systems that can perform tasks typically requiring human intelligence |
| **API**      | Application Programming Interface | Set of protocols for building software applications                            |
| **CLI**      | Command Line Interface            | Text-based interface for interacting with programs                             |
| **HTTP**     | HyperText Transfer Protocol       | Protocol for transferring data over the web                                    |
| **JSON**     | JavaScript Object Notation        | Lightweight data interchange format                                            |
| **LSP**      | Language Server Protocol          | Protocol for providing language features in editors                            |
| **REST**     | Representational State Transfer   | Architectural style for web APIs                                               |
| **SDK**      | Software Development Kit          | Collection of tools for developing applications                                |
| **SSE**      | Server-Sent Events                | Web standard for streaming data from server to client                          |
| **TUI**      | Terminal User Interface           | Text-based user interface in a terminal                                        |
| **UI**       | User Interface                    | The means by which users interact with software                                |
| **UX**       | User Experience                   | Overall experience of using a product                                          |

## 🎯 Kuucode-Specific Terms

### **Kuucode**

The name of this AI coding agent project. A fork of the original OpenCode project.

### **KuuCode (Original)**

The original project that this Kuucode was forked from, created by SST.

### **moikas-code**

The GitHub organization that owns the Kuucode fork.

### **Fork**

A copy of a repository that allows independent development. Kuucode is a fork of OpenCode.

### **Migration**

The process of moving from Stainless to OpenAPI Generator for SDK generation.

### **Provider Agnostic**

Designed to work with multiple AI providers without being tied to any specific one.

## 🔄 Process Terms

### **Onboarding**

The process of getting new developers familiar with the project. This documentation is part of onboarding.

### **Refactoring**

Restructuring existing code without changing its external behavior.

### **Debugging**

The process of finding and fixing bugs in code.

### **Profiling**

Analyzing code performance to identify bottlenecks.

### **Testing**

Verifying that code works correctly through automated tests.

---

**Don't see a term you're looking for?** Feel free to add it to this glossary or ask in the project discussions!
