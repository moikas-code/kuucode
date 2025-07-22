# SDK Generation - Quick Reference

> 📖 **Full Documentation**: See [`docs/SDK_GENERATION.md`](docs/SDK_GENERATION.md) for comprehensive documentation.

This project uses [OpenAPI Generator](https://openapi-generator.tech/) to automatically generate client SDKs from the kuucode REST API.

## Quick Start

```bash
# Generate all SDKs (TypeScript, Go, Python)
./scripts/generate-sdks

# Development mode (faster)
./scripts/generate-sdks --dev

# Clean build
./scripts/generate-sdks --clean
```

## Generated SDKs

| Language   | Output Directory        | Package Name        | Usage                    |
|------------|------------------------|--------------------|--------------------------| 
| TypeScript | `./sdks/typescript/`   | `@kuucode-ai/sdk`  | Web and Node.js apps     |
| Go         | `./packages/tui/sdk/`  | `kuucode`          | TUI component (committed)|
| Python     | `./sdks/python/`       | `kuucode_ai`       | Python applications      |

## Development Workflow

1. **Make API changes** in `packages/kuucode/src/server/server.ts`
2. **Regenerate SDKs**: `./scripts/generate-sdks --dev`
3. **Test locally**: The Go TUI automatically uses the local SDK
4. **Commit changes**: Only the Go SDK is committed to the repository

## Configuration Files

- `config/typescript-config.json` - TypeScript SDK settings
- `config/go-config.json` - Go SDK settings  
- `config/python-config.json` - Python SDK settings

## Migration from Stainless

✅ **Free and open source** (was $99/month)  
✅ **Full local control** (no external API dependencies)  
✅ **Support for 50+ languages** (was limited)  
✅ **Customizable templates** (full control)

The old `./scripts/stainless` command still works (aliased to `./scripts/generate-sdks`).

---

**Need help?** Check the [full documentation](docs/SDK_GENERATION.md) or [troubleshooting guide](docs/SDK_GENERATION.md#troubleshooting).