# 🔄 Kuucode Parity Tracking - Quick Start

> **TL;DR**: Run `bun scripts/parity-daily` for daily checks, `bun scripts/parity dashboard` for status

## What This Is

A comprehensive system to track and integrate valuable improvements from the upstream `sst/opencode` repository into your `kuuzuki` fork while maintaining your unique features.

## Current Status

✅ **System Active**: 50 upstream commits discovered and triaged  
🔍 **Pending Review**: 14 high-priority items ready for integration  
📈 **Integration Rate**: 4% (will improve as you review items)  
🤖 **Automation**: Weekly checks + GitHub Actions enabled

## Daily Workflow (2 minutes)

```bash
# Check what needs attention
bun scripts/parity-daily
```

This shows:

- New upstream commits
- High-priority items to review
- Current integration status

## Weekly Review (30-60 minutes)

```bash
# 1. See what needs review
bun scripts/parity list pending high

# 2. Review a specific commit
bun scripts/parity review abc123
# (Edit the generated template in kb/reviews/abc123-review.md)

# 3. Make decision
bun scripts/parity update abc123 integrated "Easy bug fix, no conflicts"
# OR
bun scripts/parity update abc123 skipped "Conflicts with kuuzuki features"

# 4. Check progress
bun scripts/parity dashboard
```

## Key Commands

| Command                                     | Purpose                      |
| ------------------------------------------- | ---------------------------- |
| `bun scripts/parity dashboard`              | Show current status overview |
| `bun scripts/parity list pending high`      | See high-priority items      |
| `bun scripts/parity review <hash>`          | Create review template       |
| `bun scripts/parity update <hash> <status>` | Update decision              |
| `bun scripts/parity-daily`                  | Quick daily check            |

## Integration Decision Guide

### ✅ Usually Integrate

- Security fixes
- Bug fixes with no conflicts
- Performance improvements
- Simple feature enhancements

### 🔍 Review Carefully

- New features (may enhance kuuzuki)
- API changes (check compatibility)
- UI/UX improvements (align with kuuzuki vision)

### ❌ Usually Skip

- Branding changes (opencode → kuuzuki handled separately)
- Features conflicting with kuuzuki enhancements
- Experimental/WIP code

## Files You Should Know

```
.parity-status.json           # Tracks all decisions
scripts/parity               # Main management tool
scripts/parity-daily         # Quick daily check
docs/PARITY_TRACKING.md      # Complete documentation
kb/active/parity-*.md        # Knowledge base docs
```

## Getting Started

1. **Check current status**: `bun scripts/parity dashboard`
2. **See high-priority items**: `bun scripts/parity list pending high`
3. **Pick an easy one**: Look for bug fixes or small improvements
4. **Review it**: `bun scripts/parity review <hash>`
5. **Make decision**: `bun scripts/parity update <hash> integrated "reason"`

## Automation

- **GitHub Actions**: Runs weekly Monday 9 AM UTC
- **Auto-triage**: Categorizes new commits by priority
- **Issue Creation**: Creates GitHub issues for high-priority items
- **Status Updates**: Keeps `.parity-status.json` current

## Need More Details?

📖 **Full Documentation**: `docs/PARITY_TRACKING.md`  
📚 **Knowledge Base**: `kb/active/parity-*.md`  
🛠️ **Command Help**: `bun scripts/parity` (no args)

## Current High-Priority Items

Run `bun scripts/parity list pending high` to see the 14 high-priority upstream improvements waiting for your review, including:

- SDK generation improvements
- TUI bug fixes and enhancements
- Mouse interaction fixes
- API improvements

**Start with simple bug fixes for quick wins!**
