# Enhanced Error Handling for kuucode

## Problems Addressed

### 1. Token Limit Errors
- `AI_APICallError: input length and max_tokens exceed context limit: 173941 + 32000 > 200000`
- Various token limit error formats not being caught

### 2. API Overload Errors  
- `{"type":"overloaded_error","message":"Overloaded"}`
- Anthropic rate limiting without proper retry logic

### 3. Empty Messages After Summarization
- `AI_APICallError: messages: at least one message is required`
- Summarization process removing all messages

## Enhanced Solution

### 1. Comprehensive Token Limit Detection
```typescript
const isTokenLimitError =
  errorMessage.includes("prompt is too long") ||
  errorMessage.includes("tokens >") ||
  errorMessage.includes("maximum") ||
  errorMessage.includes("context limit") ||
  errorMessage.includes("exceed") ||
  error?.name === "AI_APICallError" ||
  // Regex patterns for various formats
  /\d+\s*tokens?\s*>\s*\d+/i.test(errorMessage) ||           // "208405 tokens > 200000"
  /\d+\s*\+\s*\d+\s*>\s*\d+/i.test(errorMessage) ||         // "173941 + 32000 > 200000"
  /input\s+length.*exceed/i.test(errorMessage) ||            // "input length and max_tokens exceed"
  /context\s+limit/i.test(errorMessage)                      // "context limit"
```

### 2. API Overload Handling with Exponential Backoff
```typescript
const isOverloadedError = 
  errorType === "overloaded_error" ||
  errorMessage.includes("Overloaded") ||
  errorMessage.includes("rate limit") ||
  errorMessage.includes("too many requests")

// Exponential backoff: 2s, 4s, 8s
const backoffMs = Math.pow(2, retryCount + 1) * 1000
await new Promise(resolve => setTimeout(resolve, backoffMs))
```

### 3. Empty Messages Protection
```typescript
const isEmptyMessagesError =
  errorMessage.includes("at least one message is required") ||
  errorMessage.includes("messages: at least one")
```

## Error Handling Flow

1. **Token Limit Errors**: Auto-summarize and retry (max 2 attempts)
2. **Overload Errors**: Exponential backoff retry (max 3 attempts)  
3. **Empty Messages**: Clear error message, suggest new conversation
4. **Other Errors**: Pass through unchanged

## Retry Limits
- **Token Limit**: 2 retries with summarization
- **Overload**: 3 retries with exponential backoff (2s, 4s, 8s)
- **Empty Messages**: No retry, immediate clear error

## Benefits
✅ **Robust Pattern Matching**: Catches all known token limit error formats
✅ **Rate Limit Handling**: Proper backoff for API overload situations  
✅ **Graceful Degradation**: Clear error messages when recovery isn't possible
✅ **Comprehensive Logging**: Detailed error tracking for debugging

## Status
✅ **Implementation Complete** - Enhanced error handling deployed
⏳ **Ready for Testing** - Should handle all reported error scenarios

## Files Modified
- `packages/kuucode/src/session/index.ts` - Enhanced error handling logic