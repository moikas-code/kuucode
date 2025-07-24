# Kuuzuki Parity Tracking System

## Overview
Comprehensive system for tracking and managing parity between kuuzuki (our fork) and opencode (upstream).

## System Components

### 1. Enhanced Parity Status Schema
- **File**: `.parity-status.json`
- **Version**: 2.0
- **Features**: Rich metadata, statistics, categorization

### 2. Management Tools
- **Main Tool**: `scripts/parity` - Full-featured parity management
- **Daily Check**: `scripts/parity-daily` - Quick daily status check
- **Commands Available**:
  - `check` - Find new upstream commits
  - `triage` - Auto-categorize commits
  - `dashboard` - Status overview
  - `update` - Change commit status
  - `review` - Create review templates
  - `list` - Filter and view commits

### 3. Review Templates
- **Detailed**: `kb/templates/parity-review-template.md`
- **Quick**: `kb/templates/parity-decision-quick.md`
- **Storage**: `kb/reviews/` for completed reviews

### 4. Automation
- **GitHub Workflow**: `.github/workflows/parity-check.yml`
- **Schedule**: Weekly Monday 9 AM UTC
- **Features**: Auto-triage, issue creation, status updates

## Workflow

### Daily Development
```bash
# Quick daily check
bun scripts/parity-daily

# Check specific timeframe
bun scripts/parity check "3 days ago" --auto-triage

# View dashboard
bun scripts/parity dashboard
```

### Weekly Review Process
1. **Discovery**: Automated weekly check finds new commits
2. **Triage**: Auto-categorization by priority/type/impact
3. **Review**: Manual review of high-priority items
4. **Decision**: Integrate/Skip/Defer with reasoning
5. **Implementation**: Adapt and integrate selected changes
6. **Documentation**: Update KB with decisions and notes

### Review Process
```bash
# Create review template
bun scripts/parity review abc123

# Edit the generated template in kb/reviews/abc123-review.md

# Update status after decision
bun scripts/parity update abc123 integrated "Easy adaptation, good feature"
```

## Status Categories

### Commit Status
- `pending` - Awaiting review
- `reviewed` - Reviewed but not decided
- `integrated` - Successfully integrated
- `skipped` - Decided not to integrate
- `conflict` - Integration conflicts exist

### Priority Levels
- `critical` - Security fixes, urgent bugs
- `high` - Important features, core fixes
- `medium` - Nice-to-have features
- `low` - Documentation, minor changes

### Categories
- `feature` - New functionality
- `bugfix` - Bug fixes
- `security` - Security-related changes
- `refactor` - Code improvements
- `docs` - Documentation changes

### Impact Areas
- `tui` - Terminal UI components
- `sdk` - SDK and API changes
- `github` - GitHub integration
- `core` - Core functionality
- `docs` - Documentation

## Integration Guidelines

### Auto-Integration Candidates
- Security fixes (after review)
- Simple bug fixes with no conflicts
- Documentation improvements

### Manual Review Required
- New features
- Breaking changes
- Core architecture changes
- Changes affecting kuuzuki-specific features

### Skip Candidates
- Branding-specific changes (opencode → kuuzuki handled separately)
- Features that conflict with kuuzuki enhancements
- Experimental/WIP features

## Metrics & Reporting

### Key Metrics
- Integration rate (target: >80% for valuable changes)
- Review time (target: <1 week for high priority)
- Conflict rate (target: <10%)
- Time behind upstream (target: <2 weeks)

### Reports
- Weekly automated reports via GitHub issues
- Dashboard always available via `bun scripts/parity dashboard`
- Historical tracking in `.parity-status.json`

## Best Practices

### Decision Making
1. **Value Assessment**: Does this benefit kuuzuki users?
2. **Effort Analysis**: Cost vs. benefit of integration
3. **Conflict Check**: Does it interfere with kuuzuki features?
4. **Maintenance Impact**: Long-term maintenance burden

### Documentation
- Always document reasoning for skip decisions
- Link to KB entries for complex integrations
- Track adaptation patterns for future reference
- Maintain historical context

### Automation Balance
- Automate discovery and triage
- Keep human judgment for decisions
- Automate routine status updates
- Manual review for complex changes

## Future Enhancements

### Planned Improvements
- Automated conflict detection
- Integration testing automation
- Slack/Discord notifications
- Web dashboard interface
- ML-based priority prediction

### Integration Opportunities
- CI/CD pipeline integration
- IDE plugin for parity management
- Automated adaptation suggestions
- Community contribution tracking