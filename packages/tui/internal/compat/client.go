package compat

import (
	"context"

	kuucode "github.com/moikas-code/kuucode-sdk-go"
)

// CompatClient wraps the new OpenAPI-generated APIClient to maintain the old interface
type CompatClient struct {
	apiClient *kuucode.APIClient
	baseURL   string
}

// NewClient creates a new compatibility client that wraps the OpenAPI-generated client
func NewClient(baseURL string) *Client {
	cfg := kuucode.NewConfiguration()
	cfg.Servers = kuucode.ServerConfigurations{
		{
			URL:         baseURL,
			Description: "kuucode server",
		},
	}

	return &Client{
		apiClient: kuucode.NewAPIClient(cfg),
		baseURL:   baseURL,
	}
}

// Event returns an event service compatible with the old interface
func (c *Client) Event() *EventService {
	return &EventService{client: c}
}

// App returns an app service compatible with the old interface
func (c *Client) App() *AppService {
	return &AppService{client: c}
}

// Session returns a session service compatible with the old interface
func (c *Client) Session() *SessionService {
	return &SessionService{client: c}
}

// File returns a file service compatible with the old interface
func (c *Client) File() *FileService {
	return &FileService{client: c}
}

// EventService provides compatibility for event-related operations
type EventService struct {
	client *Client
}

// ListStreaming provides a mock streaming interface for events
// TODO: Implement proper streaming when the new SDK supports it
func (e *EventService) ListStreaming(ctx context.Context) *EventStream {
	return &EventStream{
		ctx:    ctx,
		client: e.client,
		events: make(chan interface{}, 100),
	}
}

// EventStream provides a mock event stream
type EventStream struct {
	ctx    context.Context
	client *Client
	events chan interface{}
	err    error
}

func (s *EventStream) Next() bool {
	// TODO: Implement actual event streaming
	// For now, return false to prevent infinite loops
	return false
}

func (s *EventStream) Current() EventWrapper {
	return EventWrapper{}
}

func (s *EventStream) Err() error {
	return s.err
}

// EventWrapper wraps events to provide the old AsUnion interface
type EventWrapper struct {
	event interface{}
}

func (w EventWrapper) AsUnion() interface{} {
	return w.event
}

// AppService provides compatibility for app-related operations
type AppService struct {
	client *Client
}

// Log sends a log message (placeholder implementation)
func (a *AppService) Log(ctx context.Context, params AppLogParams) error {
	// TODO: Implement actual logging via the new SDK
	return nil
}

// Providers gets the list of providers (placeholder implementation)
func (a *AppService) Providers(ctx context.Context) (*AppProvidersResponse, error) {
	// TODO: Implement actual provider fetching via the new SDK
	return &AppProvidersResponse{}, nil
}

// SessionService provides compatibility for session-related operations
type SessionService struct {
	client *Client
}

// Init initializes a session (placeholder implementation)
func (s *SessionService) Init(ctx context.Context, sessionID string, params SessionInitParams) error {
	// TODO: Implement actual session initialization via the new SDK
	return nil
}

// Chat sends a chat message (placeholder implementation)
func (s *SessionService) Chat(ctx context.Context, sessionID string, params SessionChatParams) error {
	// TODO: Implement actual chat via the new SDK
	return nil
}

// Summarize summarizes a session (placeholder implementation)
func (s *SessionService) Summarize(ctx context.Context, sessionID string, params SessionSummarizeParams) error {
	// TODO: Implement actual summarization via the new SDK
	return nil
}

// FileService provides compatibility for file-related operations
type FileService struct {
	client *Client
}

// Read reads a file (placeholder implementation)
func (f *FileService) Read(ctx context.Context, params FileReadParams) (*FileReadResponse, error) {
	// TODO: Implement actual file reading via the new SDK
	return &FileReadResponse{}, nil
}

// FindFiles finds files (placeholder implementation)
func (f *FileService) FindFiles(ctx context.Context, params FindFilesParams) (*FindFilesResponse, error) {
	// TODO: Implement actual file finding via the new SDK
	return &FindFilesResponse{}, nil
}

// FindSymbols finds symbols (placeholder implementation)
func (f *FileService) FindSymbols(ctx context.Context, params FindSymbolsParams) (*FindSymbolsResponse, error) {
	// TODO: Implement actual symbol finding via the new SDK
	return &FindSymbolsResponse{}, nil
}
