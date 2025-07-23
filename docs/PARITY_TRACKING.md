# Kuucode Parity Tracking System

## Quick Reference

### Daily Commands

```bash
# Quick daily check (2-3 minutes)
bun scripts/parity-daily

# Check dashboard anytime
bun scripts/parity dashboard

# List high priority items
bun scripts/parity list pending high
```

### Weekly Review Process

```bash
# 1. Check for new upstream commits
bun scripts/parity check "1 week ago" --auto-triage

# 2. Review high priority items
bun scripts/parity list pending high

# 3. Create review for specific commit
bun scripts/parity review abc123

# 4. Update status after decision
bun scripts/parity update abc123 integrated "Easy adaptation, good feature"
bun scripts/parity update abc123 skipped "Conflicts with kuucode features"

# 5. Check progress
bun scripts/parity dashboard
```

## System Overview

### Files & Structure

```
.parity-status.json           # Main status tracking file
scripts/parity               # Main management tool
scripts/parity-daily         # Quick daily check
kb/templates/                # Review templates
kb/reviews/                  # Completed reviews
kb/active/parity-*.md        # Documentation
.github/workflows/parity-check.yml  # Weekly automation
```

### Status Categories

- **pending** - Awaiting review
- **integrated** - Successfully integrated into kuucode
- **skipped** - Decided not to integrate
- **conflict** - Has integration conflicts
- **reviewed** - Reviewed but decision pending

### Priority Levels

- **critical** - Security fixes, urgent bugs
- **high** - Important features, core fixes
- **medium** - Nice-to-have features
- **low** - Documentation, minor changes

## Complete Command Reference

### Discovery & Triage

```bash
# Check for new commits since timeframe
bun scripts/parity check "1 week ago"
bun scripts/parity check "3 days ago" --auto-triage

# Auto-triage existing commits
bun scripts/parity triage
```

### Status & Reporting

```bash
# Show dashboard
bun scripts/parity dashboard
bun scripts/parity status

# List commits by filter
bun scripts/parity list                    # All commits
bun scripts/parity list pending           # Pending only
bun scripts/parity list pending high      # High priority pending
bun scripts/parity list integrated        # Successfully integrated
```

### Review & Decision Making

```bash
# Create review template for commit
bun scripts/parity review abc123

# Update commit status
bun scripts/parity update abc123 integrated "Reason for integration"
bun scripts/parity update abc123 skipped "Reason for skipping"
bun scripts/parity update abc123 conflict "Files that conflict"
```

## Integration Workflow

### 1. Discovery (Automated)

- Weekly GitHub Action runs Monday 9 AM UTC
- Finds new upstream commits
- Auto-triages by priority/category/impact
- Creates GitHub issues for high-priority items

### 2. Review Process

```bash
# Start with high priority items
bun scripts/parity list pending high

# Create detailed review
bun scripts/parity review [commit-hash]
# Edit the generated template in kb/reviews/[hash]-review.md

# Make decision
bun scripts/parity update [commit-hash] [status] "reasoning"
```

### 3. Integration Steps

For commits marked as "integrated":

1. **Analyze the upstream commit**

   ```bash
   git show [commit-hash]
   git show --stat [commit-hash]
   ```

2. **Adapt to kuucode**

   - Change opencode → kuucode branding
   - Resolve any conflicts with kuucode features
   - Test functionality

3. **Document the integration**
   - Update KB with adaptation notes
   - Link to review documentation

### 4. Testing Integration

```bash
# Build and test after integration
cd packages/tui && go build
cd sdks/github && bun install && bun run tsc --noEmit
```

## Decision Guidelines

### Auto-Integrate Candidates

- ✅ Security fixes (after review)
- ✅ Simple bug fixes with no conflicts
- ✅ Documentation improvements
- ✅ Performance optimizations

### Manual Review Required

- 🔍 New features
- 🔍 Breaking changes
- 🔍 Core architecture changes
- 🔍 Changes affecting kuucode-specific features

### Skip Candidates

- ❌ Branding-specific changes (opencode-specific)
- ❌ Features that conflict with kuucode enhancements
- ❌ Experimental/WIP features
- ❌ Changes that reduce kuucode's competitive advantage

## Automation

### GitHub Actions Workflow

- **Trigger**: Weekly Monday 9 AM UTC + manual dispatch
- **Actions**:
  - Fetch upstream changes
  - Auto-triage new commits
  - Generate parity report
  - Create GitHub issues for high-priority items
  - Commit updated status

### Local Automation

```bash
# Add to your daily routine
echo 'bun scripts/parity-daily' >> ~/.bashrc

# Or create a cron job for weekly checks
# 0 9 * * 1 cd /path/to/kuucode && bun scripts/parity check "1 week ago" --auto-triage
```

## Troubleshooting

### Common Issues

**"Commit not found"**

```bash
# Fetch latest upstream
git fetch upstream
bun scripts/parity check "1 month ago" --auto-triage
```

**"No upstream remote"**

```bash
git remote add upstream https://github.com/sst/opencode.git
git fetch upstream
```

**"Permission denied on scripts"**

```bash
chmod +x scripts/parity
chmod +x scripts/parity-daily
```

### Reset Parity Status

```bash
# Backup current status
cp .parity-status.json .parity-status.json.backup

# Reset and re-triage
rm .parity-status.json
bun scripts/parity check "1 month ago" --auto-triage
```

## Success Metrics

### Target Goals

- **Integration Rate**: >80% for valuable changes
- **Review Time**: <1 week for high priority items
- **Conflict Rate**: <10% of integrations
- **Time Behind Upstream**: <2 weeks

### Monitoring

```bash
# Check current metrics
bun scripts/parity dashboard

# Weekly review questions:
# 1. Are we reviewing high-priority items quickly?
# 2. Is our integration rate reasonable?
# 3. Are we staying current with valuable changes?
# 4. Do we need to adjust triage rules?
```

## Examples

### Example: Integrating a Bug Fix

```bash
# 1. Review the commit
bun scripts/parity review e4f754ee

# 2. Analyze the change
git show e4f754ee

# 3. Test integration
# (apply changes manually, test functionality)

# 4. Update status
bun scripts/parity update e4f754ee integrated "Mouse selection fix, no conflicts"
```

### Example: Skipping a Feature

```bash
# 1. Review the commit
bun scripts/parity review xyz789

# 2. Determine it conflicts with kuucode features

# 3. Update status
bun scripts/parity update xyz789 skipped "Conflicts with kuucode's enhanced TUI features"
```

## Quick Start Checklist

- [ ] Run `bun scripts/parity-daily` to see current status
- [ ] Review high-priority items: `bun scripts/parity list pending high`
- [ ] Pick 1-2 simple bug fixes to integrate first
- [ ] Set up weekly review schedule (30-60 minutes)
- [ ] Enable GitHub Actions workflow for automation
- [ ] Document your first integration experience

## Need Help?

- Check `kb/active/parity-*.md` for detailed documentation
- Review templates in `kb/templates/`
- Look at completed reviews in `kb/reviews/`
- Run `bun scripts/parity` for command help
