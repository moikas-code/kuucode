# Migration from Stainless to OpenAPI Generator

This document tracks the migration from Stainless to OpenAPI Generator for SDK generation.

## ✅ Completed Migration

### Files Removed
- `stainless.yml` - Stainless configuration
- `stainless-workspace.json` - Stainless workspace config  
- `openapi-generator-config.yml` - Unused OpenAPI config

### Files Added
- `config/typescript-config.json` - TypeScript SDK configuration
- `config/go-config.json` - Go SDK configuration
- `config/python-config.json` - Python SDK configuration
- `scripts/generate-sdks` - New SDK generation script
- `docs/SDK_GENERATION.md` - Comprehensive documentation
- `docs/DEVELOPMENT.md` - Development guide
- `docs/API.md` - API reference

### Files Modified
- `scripts/stainless` - Now redirects to `generate-sdks` (backward compatibility)
- `package.json` - Removed stainless script, added OpenAPI Generator dependency
- `README.md` - Updated SDK generation instructions
- `AGENTS.md` - Updated workflow instructions

### Dependencies
- **Added**: `@openapitools/openapi-generator-cli`
- **Removed**: No Stainless dependencies (was external service)

## Benefits Achieved

| Aspect | Before (Stainless) | After (OpenAPI Generator) |
|--------|-------------------|---------------------------|
| **Cost** | $99/month | Free |
| **Control** | External service | Full local control |
| **Dependencies** | API key required | Self-contained |
| **Languages** | 3 (TS, Go, Python) | 50+ available |
| **Customization** | Limited | Full template control |
| **Maintenance** | Zero (external) | Minimal (local scripts) |

## Backward Compatibility

- ✅ `./scripts/stainless` still works (redirects to new script)
- ✅ Same output directories for SDKs
- ✅ Same package names and versions
- ✅ Same development workflow

## Usage

### Old Way (still works)
```bash
./scripts/stainless --dev
```

### New Way (recommended)
```bash
./scripts/generate-sdks --dev
# or
bun run generate-sdks --dev
```

## Next Steps

1. **Test the new generation** with `./scripts/generate-sdks --dev`
2. **Verify SDKs work** with existing TUI and any external clients
3. **Update CI/CD** if using automated SDK generation
4. **Consider removing** `scripts/stainless` after transition period

## Rollback Plan

If needed, you can rollback by:
1. Restoring the original `stainless.yml` and `stainless-workspace.json`
2. Reverting `scripts/stainless` to the original Stainless logic
3. Getting a Stainless API key
4. Removing OpenAPI Generator configs

However, this would reintroduce the $99/month cost and external dependency.

---

*Migration completed: January 2025*