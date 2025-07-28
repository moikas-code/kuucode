# kuuzuki GitHub Action

A GitHub Action that integrates [kuuzuki](https://kuuzuki.ai) directly into your GitHub workflow.

Mention `/kuuzuki` in your comment, and kuuzuki will execute tasks within your GitHub Actions runner.

## Features

#### Triage and explain issues

```bash
/kuuzuki explain this issue
```

#### Fix or implement issues - opencode will create a PR with the changes.

```bash
/kuuzuki fix this
```

#### Review PRs and make changes

```bash
Delete the attachment from S3 when the note is removed /oc
```

## Installation

Run the following command in the terminal from your GitHub repo:

```bash
opencode github install
```

This will walk you through installing the GitHub app, creating the workflow, and setting up secrets.

### Manual Setup

1. Install the GitHub app https://github.com/apps/opencode-agent. Make sure it is installed on the target repository.
2. Add the following workflow file to `.github/workflows/opencode.yml` in your repo. Set the appropriate `model` and required API keys in `env`.

   ```yml
   name: opencode

   on:
     issue_comment:
       types: [created]

   jobs:
     opencode:
       if: |
         contains(github.event.comment.body, '/oc') ||
         contains(github.event.comment.body, '/opencode')
       runs-on: ubuntu-latest
       permissions:
         id-token: write
       steps:
         - name: Checkout repository
           uses: actions/checkout@v4
           with:
             fetch-depth: 1

         - name: Run opencode
           uses: sst/opencode/github@latest
           env:
             ANTHROPIC_API_KEY: ${{ secrets.ANTHROPIC_API_KEY }}
           with:
             model: anthropic/claude-sonnet-4-20250514
   ```

3. Store the API keys in secrets. In your organization or project **settings**, expand **Secrets and variables** on the left and select **Actions**. Add the required API keys.

## Support

This is an early release. If you encounter issues or have feedback, please create an issue at https://github.com/sst/opencode/issues.

## Development

To test locally:

1. Navigate to a test repo (e.g. `hello-world`):

   ```bash
   cd hello-world
   ```

2. Run:

   ```bash
   MODEL=anthropic/claude-sonnet-4-20250514 \
     ANTHROPIC_API_KEY=sk-ant-api03-1234567890 \
     GITHUB_RUN_ID=dummy \
     bun /path/to/opencode/packages/kuuzuki/src/index.ts github run \
     --token 'github_pat_1234567890' \
     --event '{"eventName":"issue_comment",...}'
   ```

   - `MODEL`: The model used by opencode. Same as the `MODEL` defined in the GitHub workflow.
   - `ANTHROPIC_API_KEY`: Your model provider API key. Same as the keys defined in the GitHub workflow.
   - `GITHUB_RUN_ID`: Dummy value to emulate GitHub action environment.
   - `/path/to/opencode`: Path to your cloned opencode repo. `bun /path/to/opencode/packages/kuuzuki/src/index.ts` runs your local version of `opencode`.
   - `--token`: A GitHub persontal access token. This token is used to verify you have `admin` or `write` access to the test repo. Generate a token [here](https://github.com/settings/personal-access-tokens).
   - `--event`: Mock GitHub event payload (see templates below).

### Issue comment event

```
--event '{"eventName":"issue_comment","repo":{"owner":"sst","repo":"hello-world"},"actor":"fwang","payload":{"issue":{"number":4},"comment":{"id":1,"body":"hey opencode, summarize thread"}}}'
```

Replace:

- `"owner":"sst"` with repo owner
- `"repo":"hello-world"` with repo name
- `"actor":"fwang"` with the GitHub username of commentor
- `"number":4` with the GitHub issue id
- `"body":"hey opencode, summarize thread"` with comment body

### Issue comment with image attachment.

```
--event '{"eventName":"issue_comment","repo":{"owner":"sst","repo":"hello-world"},"actor":"fwang","payload":{"issue":{"number":4},"comment":{"id":1,"body":"hey opencode, what is in my image ![Image](https://github.com/user-attachments/assets/xxxxxxxx)"}}}'
```

Replace the image URL `https://github.com/user-attachments/assets/xxxxxxxx` with a valid GitHub attachment (you can generate one by commenting with an image in any issue).

### PR comment event

```
--event '{"eventName":"issue_comment","repo":{"owner":"sst","repo":"hello-world"},"actor":"fwang","payload":{"issue":{"number":4,"pull_request":{}},"comment":{"id":1,"body":"hey opencode, summarize thread"}}}'
```
