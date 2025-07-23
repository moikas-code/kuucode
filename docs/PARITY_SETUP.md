# Parity Notification Setup Guide

## 🔒 Secure Webhook Storage

### Option 1: Local Environment Variables (Recommended)

1. **Create a `.env` file** (not committed to git):

```bash
# In your kuucode project root
echo "DISCORD_WEBHOOK_URL=https://discordapp.com/api/webhooks/1397447455800561767/" > .env
```

2. **Load environment variables** before running scripts:

```bash
# Option A: Use dotenv (install if needed)
bun add -d dotenv
export $(cat .env | xargs) && npm run parity-daily

# Option B: Export directly
export DISCORD_WEBHOOK_URL="https://discordapp.com/api/webhooks/1397447455800561767/"
npm run parity-daily
```

### Option 2: Shell Profile (Persistent)

Add to your `~/.bashrc`, `~/.zshrc`, or `~/.profile`:

```bash
export DISCORD_WEBHOOK_URL="https://discordapp.com/api/webhooks/1397447455800561767/"
```

Then reload: `source ~/.bashrc` (or restart terminal)

### Option 3: GitHub Secrets (For Actions)

1. **Go to your GitHub repository**
2. **Settings** → **Secrets and variables** → **Actions**
3. **New repository secret**:
   - Name: `DISCORD_WEBHOOK_URL`
   - Value: `https://discordapp.com/api/webhooks/1397447455800561767/`

## 🧪 Testing Setup

```bash
# Test with environment variable
export DISCORD_WEBHOOK_URL="your_webhook_url"
npm run parity-notify daily "Testing secure webhook setup!"

# Verify it works
npm run parity-daily
```

## 📁 File Security

### ✅ Safe (these files are git-ignored):

- `.env` - Local environment variables
- `.env.local` - Local overrides

### ⚠️ Be Careful:

- `.env.example` - Template file (no real secrets)
- `scripts/parity-notify` - Now reads from environment

### ❌ Never Commit:

- Webhook URLs in source code
- `.env` files with real secrets
- Hardcoded tokens or keys

## 🔄 Migration from Hardcoded

The webhook has been removed from the source code and now requires:

```bash
export DISCORD_WEBHOOK_URL="your_webhook_here"
```

This ensures:

- ✅ Secrets not in git history
- ✅ Different webhooks per environment
- ✅ Easy rotation without code changes
- ✅ Secure CI/CD integration

## 🚀 Quick Setup

```bash
# 1. Set your webhook
export DISCORD_WEBHOOK_URL="https://discordapp.com/api/webhooks/1397447455800561767/"

# 2. Test it works
npm run parity-notify daily "Setup complete!"

# 3. Run daily check
npm run parity-daily
```
