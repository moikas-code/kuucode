# Kuucode API Documentation

This document describes the Kuucode REST API that powers the client SDKs and TUI.

## Base URL

- **Development**: `http://localhost:4096`
- **Production**: `https://api.kuucode.ai`

## Authentication

The API uses bearer token authentication:

```bash
# Include in headers
Authorization: Bearer <your-token>
```

## OpenAPI Specification

The complete API specification is available at:

- **JSON**: `GET /doc` - OpenAPI 3.0 specification
- **Interactive**: `GET /docs` - Swagger UI (if available)

## Core Endpoints

### App Management

```http
GET /app
```
Get application information and status.

```http
POST /app/init
```
Initialize the application.

### Session Management

```http
GET /session
```
List all sessions.

```http
POST /session
```
Create a new session.

```http
DELETE /session/{id}
```
Delete a session.

```http
GET /session/{id}/message
```
Get messages for a session.

```http
POST /session/{id}/message
```
Send a message to a session.

### File Operations

```http
GET /file?path={path}
```
Read a file.

```http
GET /file/status?path={path}
```
Get file status information.

### Search and Find

```http
GET /find?query={query}
```
Search for text in files.

```http
GET /find/file?pattern={pattern}
```
Find files by pattern.

```http
GET /find/symbol?query={query}
```
Find symbols in code.

### Configuration

```http
GET /config
```
Get current configuration.

```http
GET /config/providers
```
Get available providers.

### Events (Server-Sent Events)

```http
GET /event
```
Stream real-time events.

## SDK Usage Examples

### TypeScript

```typescript
import { Configuration, DefaultApi } from '@kuucode-ai/sdk';

const config = new Configuration({
  basePath: 'http://localhost:4096',
  accessToken: 'your-token'
});

const api = new DefaultApi(config);

// List sessions
const sessions = await api.sessionList();

// Create a session
const newSession = await api.sessionCreate({
  name: 'My Session'
});

// Send a message
const response = await api.sessionChat({
  id: newSession.id,
  message: 'Hello, Kuucode!'
});
```

### Go

```go
package main

import (
    "context"
    "log"
    
    "github.com/moikas-code/kuucode-sdk-go"
)

func main() {
    client := kuucode.NewClient()
    
    // List sessions
    sessions, err := client.Session.List(context.TODO())
    if err != nil {
        log.Fatal(err)
    }
    
    // Create a session
    session, err := client.Session.Create(context.TODO(), kuucode.SessionCreateParams{
        Name: kuucode.F("My Session"),
    })
    if err != nil {
        log.Fatal(err)
    }
    
    // Send a message
    response, err := client.Session.Chat(context.TODO(), session.ID, kuucode.SessionChatParams{
        Message: kuucode.F("Hello, Kuucode!"),
    })
    if err != nil {
        log.Fatal(err)
    }
}
```

### Python

```python
import kuucode_ai
from kuucode_ai.rest import ApiException

configuration = kuucode_ai.Configuration(
    host="http://localhost:4096",
    access_token="your-token"
)

with kuucode_ai.ApiClient(configuration) as api_client:
    api_instance = kuucode_ai.DefaultApi(api_client)
    
    try:
        # List sessions
        sessions = api_instance.session_list()
        
        # Create a session
        new_session = api_instance.session_create(
            session_create_request=kuucode_ai.SessionCreateRequest(
                name="My Session"
            )
        )
        
        # Send a message
        response = api_instance.session_chat(
            id=new_session.id,
            session_chat_request=kuucode_ai.SessionChatRequest(
                message="Hello, Kuucode!"
            )
        )
        
    except ApiException as e:
        print(f"Exception: {e}")
```

## Error Handling

The API returns standard HTTP status codes:

- **200**: Success
- **400**: Bad Request
- **401**: Unauthorized
- **404**: Not Found
- **500**: Internal Server Error

Error responses include a JSON body:

```json
{
  "error": "Error message",
  "code": "ERROR_CODE",
  "details": {}
}
```

## Rate Limiting

The API may implement rate limiting. Check response headers:

- `X-RateLimit-Limit`: Requests per time window
- `X-RateLimit-Remaining`: Remaining requests
- `X-RateLimit-Reset`: Time when limit resets

## Streaming Events

The `/event` endpoint provides server-sent events:

```javascript
const eventSource = new EventSource('http://localhost:4096/event');

eventSource.onmessage = function(event) {
  const data = JSON.parse(event.data);
  console.log('Event:', data);
};
```

## Development

### Generating SDKs

After modifying the API, regenerate client SDKs:

```bash
./scripts/generate-sdks --dev
```

### Testing the API

```bash
# Start the server
bun run dev

# Test endpoints
curl http://localhost:4096/app
curl -H "Authorization: Bearer token" http://localhost:4096/session
```

### API Documentation

The OpenAPI specification is automatically generated from the server code. To update:

1. Modify endpoints in `packages/kuucode/src/server/server.ts`
2. Restart the server: `bun run dev`
3. Fetch updated spec: `curl http://localhost:4096/doc > openapi.json`
4. Regenerate SDKs: `./scripts/generate-sdks`

---

For more details, see the [full OpenAPI specification](http://localhost:4096/doc) or the [SDK generation documentation](SDK_GENERATION.md).