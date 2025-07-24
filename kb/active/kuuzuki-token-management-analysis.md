# Kuuzuki Token Management Analysis

## Key Files and Functions for Token Limits and API Calls

### 1. Core Token Management - `packages/kuuzuki/src/session/index.ts`

**Key Constants:**
- `OUTPUT_TOKEN_MAX = 32_000` (line 47) - Maximum output tokens allowed

**Token Limit Checking Logic:**
- **Function**: `chat()` around lines 561-572
- **Logic**: Checks if total tokens exceed 90% of context limit
- **Formula**: `tokens > Math.max((model.info.limit.context - outputLimit) * 0.9, 0)`
- **Action**: Triggers automatic summarization when limit approached

**Token Calculation:**
```typescript
const tokens = previous.tokens.input + previous.tokens.cache.read + 
               previous.tokens.cache.write + previous.tokens.output
```

**Auto-Summarization Trigger:**
- Line 565: Calls `summarize()` when token limit reached
- Line 570: Recursively calls `chat(input)` after summarization

### 2. Usage Tracking - `getUsage()` function (lines 1193-1215)

**Token Structure:**
```typescript
const tokens = {
  input: usage.inputTokens ?? 0,
  output: usage.outputTokens ?? 0,
  reasoning: 0,
  cache: {
    write: metadata?.["anthropic"]?.["cacheCreationInputTokens"] ?? 0,
    read: usage.cachedInputTokens ?? 0,
  },
}
```

**Cost Calculation:**
- Input tokens: `tokens.input * model.cost.input / 1_000_000`
- Output tokens: `tokens.output * model.cost.output / 1_000_000`
- Cache read/write: Separate pricing for cached tokens

### 3. Provider Integration - `packages/kuuzuki/src/provider/provider.ts`

**API Client Setup:**
- **Anthropic**: Lines 41-69, custom fetch with OAuth tokens
- **GitHub Copilot**: Lines 70-130, token refresh mechanism
- **OpenAI**: Lines 132+, standard API key authentication

**Model Limits Structure:**
```typescript
limit: {
  context: number,  // Context window size
  output: number,   // Max output tokens
}
```

### 4. Model Definitions - `packages/kuuzuki/src/provider/models.ts`

**Model Schema (lines 26-29):**
```typescript
limit: z.object({
  context: z.number(),  // Context window limit
  output: z.number(),   // Output token limit
}),
```

**Data Source**: Models fetched from `https://models.dev/api.json`

### 5. Error Handling - `packages/kuuzuki/src/session/index.ts`

**API Key Errors:**
- Line 1022: `LoadAPIKeyError.isInstance(e)` - Handles authentication failures
- Creates `MessageV2.AuthError` for provider authentication issues

**Error Types:**
- `LoadAPIKeyError`: API authentication failures
- `NamedError.Unknown`: Generic error handling
- `BusyError`: Session concurrency control

### 6. Context Trimming Logic

**Message Filtering:**
- Line 576-577: Filters messages after last summary
- Line 531-551: Message trimming in preparation phase
- **Summarization**: `summarize()` function (lines 1109+) creates condensed conversation history

**Provider-Specific Transforms:**
- `packages/kuuzuki/src/provider/transform.ts`
- Anthropic-specific caching controls for context optimization

### 7. Session Management

**Context Window Management:**
- Tracks cumulative token usage across conversation
- Automatically triggers summarization at 90% of context limit
- Maintains conversation continuity through summarization

**Token Tracking:**
- Real-time token counting for input/output/cache
- Cost calculation per provider pricing
- Usage metadata preservation

## Key Findings

1. **No hardcoded 200,000 limit found** - Uses dynamic model limits from models.dev
2. **32,000 output token maximum** hardcoded as safety limit
3. **90% threshold** for context window management
4. **Automatic summarization** prevents context overflow
5. **Provider-agnostic** token management with provider-specific optimizations
6. **Real-time cost tracking** with decimal precision
7. **Comprehensive error handling** for API failures and limits

## Critical Functions

- `Session.chat()` - Main token limit checking and management
- `getUsage()` - Token counting and cost calculation  
- `Session.summarize()` - Context trimming through summarization
- `Provider.getModel()` - Model limit retrieval
- `ProviderTransform.message()` - Provider-specific optimizations