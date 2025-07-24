# Development Guide

This guide covers the essential information for developing Kuuzuki locally.

## Table of Contents

- [Setup](#setup)
- [Architecture](#architecture)
- [Development Workflow](#development-workflow)
- [SDK Generation](#sdk-generation)
- [Testing](#testing)
- [Contributing](#contributing)

## Setup

### Prerequisites

- **Bun** (recommended) or Node.js 18+
- **Go** 1.24+ (for TUI component)
- **Git**

### Installation

```bash
# Clone the repository
git clone https://github.com/moikas-code/kuuzuki.git
cd kuuzuki

# Install dependencies
bun install

# Start development server
bun run dev
```

## Architecture

Kuuzuki uses a client/server architecture:

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   TUI (Go)      │───▶│  Server (TS)    │───▶│   Providers     │
│                 │    │                 │    │   (Anthropic,   │
│ Terminal UI     │    │ API & Business  │    │    OpenAI, etc) │
│ User Interface  │    │ Logic           │    │                 │
└─────────────────┘    └─────────────────┘    └─────────────────┘
```

### Key Components

- **`packages/kuuzuki/`**: TypeScript server (API, business logic)
- **`packages/tui/`**: Go terminal user interface
- **`packages/web/`**: Documentation website
- **`sdks/`**: Generated client SDKs
- **`scripts/`**: Build and development scripts

## Development Workflow

### 1. Running Components

**Main Server** (TypeScript):
```bash
# Development mode with hot reload
bun run dev

# Or directly
bun run packages/kuuzuki/src/index.ts
```

**TUI** (Go):
```bash
cd packages/tui
go run .
```

**Documentation Site**:
```bash
cd packages/web
bun run dev
```

### 2. Making Changes

#### Server API Changes

1. Edit files in `packages/kuuzuki/src/server/`
2. The dev server will automatically restart
3. Regenerate SDKs if you changed the API:
   ```bash
   ./scripts/generate-sdks --dev
   ```

#### TUI Changes

1. Edit Go files in `packages/tui/`
2. Restart the TUI:
   ```bash
   cd packages/tui
   go run .
   ```

#### Documentation Changes

1. Edit files in `packages/web/src/content/`
2. The dev server will hot reload

### 3. Testing Changes

```bash
# Type check all packages
bun run typecheck

# Test the CLI directly
bun run packages/kuuzuki/src/index.ts --help

# Test TUI with local server
cd packages/tui && go run .
```

## SDK Generation

Kuuzuki automatically generates client SDKs from the OpenAPI specification.

### Quick Commands

```bash
# Generate all SDKs
./scripts/generate-sdks

# Development mode (faster)
./scripts/generate-sdks --dev

# Clean previous builds
./scripts/generate-sdks --clean
```

### When to Regenerate

- After changing API endpoints in `packages/kuuzuki/src/server/`
- After modifying request/response types
- Before committing API changes

### Generated SDKs

| Language   | Location              | Usage                |
|------------|-----------------------|----------------------|
| TypeScript | `./sdks/typescript/`  | Web/Node.js clients  |
| Go         | `./packages/tui/sdk/` | TUI component        |
| Python     | `./sdks/python/`      | Python clients       |

> 📖 **Full SDK Documentation**: [`SDK_GENERATION.md`](../SDK_GENERATION.md)

## Testing

### Manual Testing

```bash
# Start the server
bun run dev

# In another terminal, test commands
bun run packages/kuuzuki/src/index.ts auth login
bun run packages/kuuzuki/src/index.ts --help
```

### TUI Testing

```bash
# Start server first
bun run dev

# In another terminal, run TUI
cd packages/tui
go run .
```

### API Testing

```bash
# Start server
bun run dev

# Test API endpoints
curl http://localhost:4096/app
curl http://localhost:4096/session
```

## Contributing

### Code Style

- **TypeScript**: Follow existing patterns, use Prettier
- **Go**: Follow Go conventions, use `gofmt`
- **Documentation**: Use clear, concise language

### Commit Guidelines

```bash
# Good commit messages
git commit -m "Add new session management endpoint"
git commit -m "Fix TUI rendering issue with long messages"
git commit -m "Update TypeScript SDK configuration"

# Include SDK updates when changing APIs
git add packages/tui/sdk/
git commit -m "Add user preferences API

- Add GET/POST /user/preferences endpoints
- Update Go SDK with new types
- Add TUI integration for preferences"
```

### Pull Request Process

1. **Fork** the repository
2. **Create** a feature branch
3. **Make** your changes
4. **Regenerate** SDKs if needed: `./scripts/generate-sdks`
5. **Test** your changes locally
6. **Commit** with clear messages
7. **Push** and create a pull request

### Development Tips

1. **Use `--dev` flag** for faster SDK generation during development
2. **Keep the server running** in one terminal while developing
3. **Check logs** in `~/.local/share/kuuzuki/log/` for debugging
4. **Use the config file** `kuuzuki.json` for local settings
5. **Test with different providers** to ensure compatibility

## Common Tasks

### Adding a New API Endpoint

1. **Add endpoint** in `packages/kuuzuki/src/server/server.ts`:
   ```typescript
   app.get('/api/v1/new-endpoint', async (c) => {
     return c.json({ message: 'Hello World' })
   })
   ```

2. **Regenerate SDKs**:
   ```bash
   ./scripts/generate-sdks --dev
   ```

3. **Test the endpoint**:
   ```bash
   curl http://localhost:4096/api/v1/new-endpoint
   ```

4. **Use in TUI** (Go SDK auto-updated):
   ```go
   response, err := client.NewEndpoint.Get(ctx)
   ```

### Adding a New Configuration Option

1. **Add to config schema** in `packages/kuuzuki/src/config/`
2. **Update validation** and types
3. **Document** in `packages/web/src/content/docs/config.mdx`
4. **Test** with example config

### Debugging Issues

1. **Check server logs**:
   ```bash
   tail -f ~/.local/share/kuuzuki/log/*.log
   ```

2. **Enable debug mode**:
   ```bash
   DEBUG=1 bun run dev
   ```

3. **Validate OpenAPI spec**:
   ```bash
   jq empty openapi.json
   ```

4. **Check TUI logs**:
   ```bash
   cd packages/tui
   go run . --debug
   ```

## Resources

- **OpenAPI Generator**: [openapi-generator.tech](https://openapi-generator.tech/)
- **Bun Documentation**: [bun.sh/docs](https://bun.sh/docs)
- **Go Documentation**: [golang.org/doc](https://golang.org/doc/)
- **Project Issues**: [github.com/moikas-code/kuuzuki/issues](https://github.com/moikas-code/kuuzuki/issues)

---

*Happy coding! 🚀*