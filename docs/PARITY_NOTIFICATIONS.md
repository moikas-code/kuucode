# Kuucode Parity Notification System

## Overview

Automated Discord notifications for upstream parity tracking, keeping you informed of important changes and required actions.

## Features

### 🔔 Notification Types

1. **Daily Status** (Blue) - Regular parity check results
2. **Critical Alerts** (Red) - High priority commits requiring immediate attention
3. **Integration Updates** (Green) - Successful upstream integrations
4. **Security Alerts** (Orange) - Security-related upstream changes

### 📊 Rich Context

Each notification includes:

- Current integration rate and statistics
- List of pending high-priority commits
- Direct links to review actions
- Timestamp and system attribution

## Usage

### Manual Notifications

```bash
# Send different types of notifications
npm run parity-notify daily "Custom daily message"
npm run parity-notify critical "High priority security fix available"
npm run parity-notify integration "Successfully integrated feature X"
npm run parity-notify security "CVE fix available upstream"
```

### Automated Notifications

- **Daily checks**: Automatically triggered by `npm run parity-daily`
- **GitHub Actions**:
  - Daily at 9 AM UTC (`.github/workflows/parity-daily.yml`)
  - Weekly reviews on Mondays (`.github/workflows/parity-check.yml`)
- **Manual triggers**: Run `npm run parity-daily` anytime or trigger workflows via GitHub

## Configuration

### Discord Webhook

Currently configured for: `#kuucode-parity` channel

- **Local development**: Uses `DISCORD_WEBHOOK_URL` environment variable
- **GitHub Actions**: Uses `DISCORD_WEBHOOK_URL` repository secret
- **Security**: Webhook URL stored securely, not in source code

### Notification Logic

- **High priority items found** → Critical notification with commit details
- **No high priority items** → Daily status notification
- **Manual integrations** → Integration success notifications

## Automation Setup

### GitHub Actions (Recommended)

```yaml
# .github/workflows/parity-daily.yml
- Runs daily at 9 AM UTC
- Fetches upstream changes
- Sends Discord notifications
- Uploads parity status artifacts
```

### Local Cron (Alternative)

```bash
# Add to crontab for daily 9 AM checks
0 9 * * * cd /path/to/kuucode && npm run parity-daily
```

## Notification Examples

### Daily Status (No Issues)

```
📅 Daily Parity Check
Daily parity check complete - no high priority items pending

📊 Integration Rate: 39.3%
✅ Integrated: 24
🔍 Pending: 0
```

### Critical Alert (High Priority Items)

```
🚨 Critical Parity Alert
Found 2 high priority commits requiring review

📊 Integration Rate: 39.3%
✅ Integrated: 24
🔍 Pending: 2

🎯 2 Commits:
• abc123 [high] security: Fix authentication bypass
• def456 [high] bugfix: Resolve memory leak in TUI
```

### Integration Success

```
✅ Integration Complete
Successfully integrated upstream commit b7b0cdbd - session sorting improvement

📊 Integration Rate: 39.5%
✅ Integrated: 25
🔍 Pending: 0
```

## Benefits

1. **Proactive Monitoring** - Never miss important upstream changes
2. **Prioritized Alerts** - Focus on high-impact items first
3. **Team Coordination** - Shared visibility into parity status
4. **Historical Tracking** - Maintain context of integration decisions
5. **Automated Workflow** - Reduce manual monitoring overhead

## Next Steps

1. **Security Enhancement**: Move webhook URL to environment variables
2. **Multi-Channel Support**: Add different channels for different alert types
3. **Integration Automation**: Auto-create PRs for simple upstream changes
4. **Metrics Dashboard**: Web interface for parity trends and analytics
