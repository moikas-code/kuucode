<p align="center">
  <a href="https://kuucode.ai">
    <picture>
      <source srcset="packages/web/src/assets/logo-ornate-dark.svg" media="(prefers-color-scheme: dark)">
      <source srcset="packages/web/src/assets/logo-ornate-light.svg" media="(prefers-color-scheme: light)">
      <img src="packages/web/src/assets/logo-ornate-light.svg" alt="kuucode logo">
    </picture>
  </a>
</p>
<p align="center">AI coding agent, built for the terminal.</p>
<p align="center">
  <a href="https://kuucode.ai/discord"><img alt="Discord" src="https://img.shields.io/discord/1391832426048651334?style=flat-square&label=discord" /></a>
  <a href="https://www.npmjs.com/package/kuucode-ai"><img alt="npm" src="https://img.shields.io/npm/v/kuucode-ai?style=flat-square" /></a>
  <a href="https://github.com/moikas-code/kuucode/actions/workflows/publish.yml"><img alt="Build status" src="https://img.shields.io/github/actions/workflow/status/moikas-code/kuucode/publish.yml?style=flat-square&branch=dev" /></a>
</p>

[![kuucode Terminal UI](packages/web/src/assets/lander/screenshot.png)](https://kuucode.ai)

---

### Installation

```bash
# YOLO
curl -fsSL https://kuucode.ai/install | bash

# Package managers
npm i -g kuucode-ai@latest        # or bun/pnpm/yarn
brew install moikas-code/tap/kuucode      # macOS
paru -S kuucode-bin               # Arch Linux
```

> [!TIP]
> Remove versions older than 0.1.x before installing.

#### Installation Directory

The install script respects the following priority order for the installation path:

1. `$KUUCODE_INSTALL_DIR` - Custom installation directory
2. `$XDG_BIN_DIR` - XDG Base Directory Specification compliant path
3. `$HOME/bin` - Standard user binary directory (if exists or can be created)
4. `$HOME/.kuucode/bin` - Default fallback

```bash
# Examples
KUUCODE_INSTALL_DIR=/usr/local/bin curl -fsSL https://kuucode.ai/install | bash
XDG_BIN_DIR=$HOME/.local/bin curl -fsSL https://kuucode.ai/install | bash
```

### Documentation

For more info on how to configure kuucode [**head over to our docs**](https://kuucode.ai/docs).

### Contributing

kuucode is an opinionated tool so any fundamental feature needs to go through a
design process with the core team.

> [!IMPORTANT]
> We do not accept PRs for core features.

However we still merge a ton of PRs - you can contribute:

- Bug fixes
- Improvements to LLM performance
- Support for new providers
- Fixes for env specific quirks
- Missing standard behavior
- Documentation

Take a look at the git history to see what kind of PRs we end up merging.

> [!NOTE]
> If you do not follow the above guidelines we might close your PR.

To run kuucode locally you need.

- Bun
- Golang 1.24.x

And run.

```bash
$ bun install
$ bun run packages/kuucode/src/index.ts
```

#### Development Notes

**API Client**: After making changes to the TypeScript API endpoints in `packages/kuucode/src/server/server.ts`, you can regenerate the SDKs locally using `./scripts/generate-sdks` or `bun run generate-sdks`.

**SDK Generation**: This project uses OpenAPI Generator to create client SDKs for TypeScript, Go, and Python. See [`SDK_GENERATION.md`](SDK_GENERATION.md) for details.

### FAQ

#### How is this different than Claude Code?

It's very similar to Claude Code in terms of capability. Here are the key differences:

- 100% open source
- Not coupled to any provider. Although Anthropic is recommended, kuucode can be used with OpenAI, Google or even local models. As models evolve the gaps between them will close and pricing will drop so being provider agnostic is important.
- A focus on TUI. kuucode is built by neovim users and the creators of [terminal.shop](https://terminal.shop); we are going to push the limits of what's possible in the terminal.
- A client/server architecture. This for example can allow kuucode to run on your computer, while you can drive it remotely from a mobile app. Meaning that the TUI frontend is just one of the possible clients.

#### What's the other repo?

The other confusingly named repo has no relation to this one. You can [read the story behind it here](https://x.com/thdxr/status/1933561254481666466).

---

**Join our community** [Discord](https://discord.gg/kuucode) | [YouTube](https://www.youtube.com/c/sst-dev) | [X.com](https://x.com/SST_dev)
