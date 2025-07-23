# GitHub Actions Parity Setup

## ✅ Configuration Complete

You've successfully configured Discord notifications for GitHub Actions! Here's what's set up:

### 🔐 GitHub Secrets

- **Secret Name**: `DISCORD_WEBHOOK_URL`
- **Secret Value**: Your Discord webhook URL
- **Location**: Repository Settings → Secrets and variables → Actions

### 🤖 Automated Workflows

#### 1. Daily Parity Check (`.github/workflows/parity-daily.yml`)

- **Schedule**: Daily at 9 AM UTC
- **Triggers**: Scheduled + Manual
- **Actions**:
  - Fetches upstream changes
  - Runs auto-triage
  - Sends Discord notifications
  - Uploads parity status artifacts

#### 2. Weekly Parity Review (`.github/workflows/parity-check.yml`)

- **Schedule**: Weekly on Mondays at 9 AM UTC
- **Triggers**: Scheduled + Manual with parameters
- **Actions**:
  - Comprehensive parity analysis
  - Creates GitHub issues for high-priority items
  - Sends Discord notifications
  - Commits updated parity status

### 📬 Notification Types

| Trigger                 | Notification   | When                           |
| ----------------------- | -------------- | ------------------------------ |
| **Daily Check**         | Daily status   | Every day at 9 AM UTC          |
| **High Priority Found** | Critical alert | When urgent commits detected   |
| **Weekly Review**       | Summary report | Mondays with detailed analysis |
| **Manual Trigger**      | Custom message | When workflows run manually    |

### 🎛️ Manual Triggers

```bash
# Trigger daily check manually
gh workflow run "Daily Parity Check"

# Trigger weekly review with custom timeframe
gh workflow run "Parity Check" -f since="2 weeks ago"
```

### 📊 Monitoring

**GitHub Actions Tab**: View workflow runs and logs
**Discord Channel**: Receive real-time notifications
**Artifacts**: Download parity status files for analysis
**Issues**: Auto-created for high-priority items

### 🔧 Customization

**Change notification schedule**:

```yaml
# In .github/workflows/parity-daily.yml
schedule:
  - cron: "0 9 * * *" # Daily at 9 AM UTC
  - cron: "0 21 * * *" # Also at 9 PM UTC
```

**Modify notification messages**:

```bash
# In the workflow files, update the parity-notify calls
bun scripts/parity-notify daily "Custom message here"
```

### 🚨 Troubleshooting

**No notifications received?**

1. Check GitHub Actions logs for errors
2. Verify `DISCORD_WEBHOOK_URL` secret is set correctly
3. Test webhook manually: `npm run parity-notify daily "test"`

**Workflow not running?**

1. Ensure workflows are enabled in repository settings
2. Check if repository has recent activity (GitHub may pause workflows)
3. Manually trigger to test: `gh workflow run "Daily Parity Check"`

### 🎯 Next Steps

1. **Monitor first automated run** (next 9 AM UTC)
2. **Adjust schedule** if needed for your timezone
3. **Customize notification messages** for your team
4. **Add more notification channels** (Slack, email, etc.)

Your parity system is now fully automated! 🎉
