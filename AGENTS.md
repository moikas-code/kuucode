# TUI Agent Guidelines

## Style

- prefer single word variable/function names
- avoid try catch where possible - prefer to let exceptions bubble up
- avoid else statements where possible
- do not make useless helper functions - inline functionality unless the
  function is reusable or composable
- prefer Bun apis

## Workflow

- you can regenerate the SDKs (Go, TypeScript, Python) by calling ./scripts/generate-sdks
- we use bun and the moidvk mcp for everything
