## Style

- prefer single word variable/function names
- avoid try catch where possible - prefer to let exceptions bubble up
- avoid else statements where possible
- do not make useless helper functions - inline functionality unless the
  function is reusable or composable
- prefer Bun apis

## Workflow

- you can regenerate the SDKs (Go, TypeScript, Python) by calling ./scripts/generate-sdks
- we use bun and the moidvk mcp for everything
  CRITICAL MCP Tool Usage Rules:

✅ Correct Tool Names:

• KB-MCP: kb-mcp_kb_read, kb-mcp_kb_update, kb-
mcp_kb_search
• Moidvk: moidvk_check_code_practices, moidvk_format_code,
moidvk_scan_security_vulnerabilities

❌ Never Use:

• Mcp**kb**kb_update (wrong casing/format)
• kb_update (missing namespace)
• mcp**moidvk**format_code (double underscores)

📋 Before Using Any MCP Tool:

1. Verify exact tool name spelling (case-sensitive)
2. Check required parameters - don't guess
3. Use namespace prefix: kb-mcp* or moidvk*
4. Test simple operations first

🔧 Tool Name Format:

• Knowledge Base: kb-mcp*[action]
• Development Tools: moidvk*[function]
