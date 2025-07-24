# TUI SDK Migration Task

## Overview

Migrate the TUI codebase from the old custom SDK interface to the new OpenAPI-generated SDK structure.

## Background

- Successfully generated new OpenAPI SDK for Go, TypeScript, and Python
- TUI code was written for old custom SDK interface
- New SDK has different structure: field names (ID→Id, URL→Url), client interface, type system
- Current compiled binary works but cannot recompile due to SDK incompatibility

## Scope

Approximately 15-20 files need updates:

- `cmd/kuuzuki/main.go` - Client initialization
- `internal/api/api.go` - API service layer
- `internal/app/app.go` - Core app logic
- `internal/util/apilogger.go` - Logging integration
- `internal/components/chat/*.go` - Chat components
- `internal/components/dialog/*.go` - Dialog components
- `internal/completions/*.go` - Completion handlers
- `internal/commands/command.go` - Command handling
- `internal/tui/tui.go` - Main TUI logic

## Migration Strategy Options

### Option 1: Full Migration (Recommended)

- Update all TUI code to use new SDK directly
- Benefits: Clean, maintainable, uses standard OpenAPI patterns
- Effort: 3-5 days
- Risk: Medium (systematic changes)

### Option 2: Compatibility Layer

- Create wrapper that adapts new SDK to old interface
- Benefits: Minimal TUI code changes
- Effort: 2-3 days
- Risk: High (complex type mapping, maintenance burden)

### Option 3: Custom SDK Generation

- Generate SDK with templates matching old interface
- Benefits: No TUI code changes needed
- Effort: 1-2 days
- Risk: Low (maintains existing patterns)

## Key Changes Needed

### Client Interface

```go
// Old
httpClient := kuuzuki.NewClient(option.WithBaseURL(url))

// New
cfg := kuuzuki.NewConfiguration()
cfg.Servers = kuuzuki.ServerConfigurations{{URL: url}}
httpClient := kuuzuki.NewAPIClient(cfg)
```

### Field Names

- All `ID` fields → `Id`
- All `URL` fields → `Url`
- Pointer handling for optional fields

### Type System

- `kuuzuki.MessageUnion` → New union handling
- `kuuzuki.PartUnion` → New part types
- Event types restructured
- Parameter types changed

### API Methods

- `httpClient.Event.ListStreaming()` → New event API
- Service method signatures changed
- Response structures different

## Implementation Plan

### Phase 1: Core Infrastructure (Day 1)

1. Update main.go client initialization
2. Update API service layer
3. Update logging integration
4. Test basic compilation

### Phase 2: Data Types (Day 2)

1. Update all ID/URL field references
2. Fix pointer/optional field handling
3. Update struct literals
4. Test data flow

### Phase 3: API Integration (Day 3)

1. Update event handling
2. Update API method calls
3. Update parameter construction
4. Test API communication

### Phase 4: UI Components (Day 4)

1. Update chat components
2. Update dialog components
3. Update completion handlers
4. Test UI functionality

### Phase 5: Testing & Polish (Day 5)

1. End-to-end testing
2. Error handling verification
3. Performance validation
4. Documentation updates

## Success Criteria

- [ ] TUI compiles without errors
- [ ] All existing functionality works
- [ ] No regression in performance
- [ ] Authentication flow works
- [ ] Chat interface functional
- [ ] File operations work
- [ ] Event streaming operational

## Current Status

- ✅ New SDK generated successfully
- ✅ Existing compiled binary works (interim solution)
- ❌ Cannot recompile due to SDK incompatibility
- ⏳ Migration task defined and ready to implement

## Priority

**Medium** - Current binary works for immediate needs, but migration needed for:

- Future TUI development
- Bug fixes requiring recompilation
- New feature additions
- Maintenance and updates

## Dependencies

- OpenAPI-generated SDK (completed)
- Go development environment
- Test authentication setup

## Estimated Effort

**3-5 days** for full migration approach (recommended)

## Notes

- Existing binary can be used while migration is in progress
- No impact on npm publishing or user functionality
- Migration only affects development/compilation workflow

## Implementation Checklist

### Pre-Implementation

- [ ] Set up development environment
- [ ] Create feature branch for migration
- [ ] Backup current working binary
- [ ] Document current API usage patterns

### Phase 1: Core Infrastructure

- [ ] Update `cmd/kuuzuki/main.go` client initialization
- [ ] Update `internal/api/api.go` service layer
- [ ] Update `internal/util/apilogger.go` logging integration
- [ ] Test basic compilation

### Phase 2: Data Types

- [ ] Update all `ID` → `Id` field references
- [ ] Update all `URL` → `Url` field references
- [ ] Fix pointer/optional field handling
- [ ] Update struct literals
- [ ] Test data flow

### Phase 3: API Integration

- [ ] Update event handling system
- [ ] Update API method calls
- [ ] Update parameter construction
- [ ] Test API communication

### Phase 4: UI Components

- [ ] Update chat components
- [ ] Update dialog components
- [ ] Update completion handlers
- [ ] Test UI functionality

### Phase 5: Testing & Polish

- [ ] End-to-end testing
- [ ] Error handling verification
- [ ] Performance validation
- [ ] Documentation updates
- [ ] Create new binary
- [ ] Verify all functionality works

## Risk Mitigation

- Keep existing binary as fallback
- Implement in feature branch
- Test each phase thoroughly before proceeding
- Document any breaking changes discovered
- Have rollback plan ready
