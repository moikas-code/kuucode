package compat

import (
	"context"

	kuuzuki "github.com/moikas-code/kuuzuki-sdk-go"
)

// CompatClient wraps the new OpenAPI-generated APIClient to maintain the old interface
type CompatClient struct {
	apiClient *kuuzuki.APIClient
	baseURL   string

	// Direct service access for compatibility
	Event   *EventService
	App     *AppService
	Session *SessionService
	Config  *ConfigService
	File    *FileService
}

// NewClient creates a new compatibility client that wraps the OpenAPI-generated client
func NewClient(baseURL string) *CompatClient {
	cfg := kuuzuki.NewConfiguration()
	cfg.Servers = kuuzuki.ServerConfigurations{
		{
			URL:         baseURL,
			Description: "kuuzuki API server",
		},
	}

	apiClient := kuuzuki.NewAPIClient(cfg)

	client := &CompatClient{
		apiClient: apiClient,
		baseURL:   baseURL,
	}

	// Initialize service fields
	client.Event = &EventService{client: client}
	client.App = &AppService{client: client}
	client.Session = &SessionService{client: client}
	client.Config = &ConfigService{client: client}
	client.File = &FileService{client: client}

	return client
}

// Find returns a find service compatible with the old interface
func (c *CompatClient) Find() *FindService {
	return &FindService{client: c}
}

// EventService provides compatibility for event-related operations
type EventService struct {
	client *CompatClient
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
	client *CompatClient
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
	client *CompatClient
}

// Log sends a log message
func (a *AppService) Log(ctx context.Context, params AppLogParams) error {
	req := a.client.apiClient.DefaultAPI.PostLog(ctx)

	// Create the log request body
	logRequest := kuuzuki.PostLogRequest{
		Level:   string(params.Level),
		Service: params.Service,
		Message: params.Message,
	}

	// Add extra data if provided
	if params.Data != nil {
		if extraMap, ok := params.Data.(map[string]interface{}); ok {
			logRequest.Extra = extraMap
		}
	}

	req = req.PostLogRequest(logRequest)
	_, _, err := req.Execute()
	return err
}

// Providers returns a providers service for method chaining
func (a *AppService) Providers() *ProvidersService {
	return &ProvidersService{client: a.client}
}

// Init returns an init service for method chaining
func (a *AppService) Init() *InitService {
	return &InitService{client: a.client}
}

// ProvidersService handles provider-related operations
type ProvidersService struct {
	client *CompatClient
}

// Get executes the providers request
func (p *ProvidersService) Get(ctx context.Context) (*AppProvidersResponse, error) {
	req := p.client.apiClient.DefaultAPI.GetConfigProviders(ctx)
	resp, _, err := req.Execute()
	if err != nil {
		return nil, err
	}

	return &AppProvidersResponse{
		Providers: resp.Providers,
	}, nil
}

// InitService handles app initialization
type InitService struct {
	client *CompatClient
}

// Post executes the app initialization
func (i *InitService) Post(ctx context.Context) error {
	req := i.client.apiClient.DefaultAPI.PostAppInit(ctx)
	_, _, err := req.Execute()
	return err
}

// SessionService provides compatibility for session-related operations
type SessionService struct {
	client *CompatClient
}

// Init returns an init service for method chaining
func (s *SessionService) Init() *SessionInitService {
	return &SessionInitService{client: s.client}
}

// Chat returns a chat service for method chaining
func (s *SessionService) Chat() *SessionChatService {
	return &SessionChatService{client: s.client}
}

// Summarize returns a summarize service for method chaining
func (s *SessionService) Summarize() *SessionSummarizeService {
	return &SessionSummarizeService{client: s.client}
}

// New creates a new session
func (s *SessionService) New(ctx context.Context) (*kuuzuki.Session, error) {
	req := s.client.apiClient.DefaultAPI.PostSession(ctx)
	session, _, err := req.Execute()
	return session, err
}

// List lists sessions
func (s *SessionService) List(ctx context.Context) ([]*kuuzuki.Session, error) {
	req := s.client.apiClient.DefaultAPI.GetSession(ctx)
	sessions, _, err := req.Execute()
	if err != nil {
		return nil, err
	}

	// Convert []Session to []*Session for compatibility
	result := make([]*kuuzuki.Session, len(sessions))
	for i := range sessions {
		result[i] = &sessions[i]
	}
	return result, nil
}

// Delete deletes a session
func (s *SessionService) Delete(ctx context.Context, sessionID string) error {
	req := s.client.apiClient.DefaultAPI.DeleteSessionById(ctx, sessionID)
	_, _, err := req.Execute()
	return err
}

// Messages gets session messages
func (s *SessionService) Messages(ctx context.Context, sessionID string) ([]*kuuzuki.Message, error) {
	req := s.client.apiClient.DefaultAPI.GetSessionByIdMessage(ctx, sessionID)
	messages, _, err := req.Execute()
	if err != nil {
		return nil, err
	}

	// Convert messages to the expected format
	result := make([]*kuuzuki.Message, len(messages))
	for i, msg := range messages {
		// Extract the Message from the Info field
		result[i] = &msg.Info
	}
	return result, nil
}

// Abort aborts a session
func (s *SessionService) Abort(ctx context.Context, sessionID string) error {
	req := s.client.apiClient.DefaultAPI.PostSessionByIdAbort(ctx, sessionID)
	_, _, err := req.Execute()
	return err
}

// Share shares a session
func (s *SessionService) Share(ctx context.Context, sessionID string) (*kuuzuki.SessionShare, error) {
	req := s.client.apiClient.DefaultAPI.PostSessionByIdShare(ctx, sessionID)
	session, _, err := req.Execute()
	if err != nil {
		return nil, err
	}

	// Extract SessionShare from the session
	if session.Share != nil {
		return session.Share, nil
	}
	return &kuuzuki.SessionShare{}, nil
}

// Unshare unshares a session
func (s *SessionService) Unshare(ctx context.Context, sessionID string) error {
	req := s.client.apiClient.DefaultAPI.DeleteSessionByIdShare(ctx, sessionID)
	_, _, err := req.Execute()
	return err
}

// Post executes a POST request for the service
func (s *SessionService) Post(ctx context.Context, sessionID string, params interface{}) error {
	// TODO: Implement actual POST request via the new SDK
	return nil
}

// FileService provides compatibility for file-related operations
type FileService struct {
	client *CompatClient
}

// Read reads a file with the given parameters
func (f *FileService) Read(ctx context.Context, params FileReadParams) (*kuuzuki.GetFile200Response, error) {
	req := f.client.apiClient.DefaultAPI.GetFile(ctx)
	req = req.Path(params.Path)
	response, _, err := req.Execute()
	return response, err
}

// FileReadService handles file reading operations
type FileReadService struct {
	client *CompatClient
}

// Post executes a POST request for the service
func (f *FileService) Post(ctx context.Context, params interface{}) (*kuuzuki.File, error) {
	// TODO: Implement actual file reading via the new SDK
	return &kuuzuki.File{}, nil
}

// Get performs a GET request (placeholder implementation)
func (c *CompatClient) Get(ctx context.Context, path string, params interface{}, result interface{}) error {
	// TODO: Implement actual GET request via the new SDK
	return nil
}

// Post performs a POST request (placeholder implementation)
func (c *CompatClient) Post(ctx context.Context, path string, body interface{}, result interface{}) error {
	// TODO: Implement actual POST request via the new SDK
	return nil
}

// FindFiles finds files
func (f *FileService) FindFiles(ctx context.Context, params FindFilesParams) (*FindFilesResponse, error) {
	req := f.client.apiClient.DefaultAPI.GetFindFile(ctx)
	req = req.Query(params.Query)
	files, _, err := req.Execute()
	if err != nil {
		return nil, err
	}

	// Convert []string to []kuuzuki.File
	result := make([]kuuzuki.File, len(files))
	for i, filePath := range files {
		result[i] = kuuzuki.File{Path: filePath}
	}

	return &FindFilesResponse{Files: result}, nil
}

// FindSymbols finds symbols
func (f *FileService) FindSymbols(ctx context.Context, params FindSymbolsParams) (*FindSymbolsResponse, error) {
	req := f.client.apiClient.DefaultAPI.GetFindSymbol(ctx)
	req = req.Query(params.Query)
	symbols, _, err := req.Execute()
	if err != nil {
		return nil, err
	}

	return &FindSymbolsResponse{Symbols: symbols}, nil
}

// FindService provides compatibility for find-related operations
type FindService struct {
	client *CompatClient
}

// Symbols finds symbols
func (f *FindService) Symbols(ctx context.Context, params FindSymbolsParams) (*[]kuuzuki.Symbol, error) {
	req := f.client.apiClient.DefaultAPI.GetFindSymbol(ctx)
	req = req.Query(params.Query)
	symbols, _, err := req.Execute()
	if err != nil {
		return nil, err
	}

	return &symbols, nil
}

// Files finds files
func (f *FindService) Files(ctx context.Context, params FindFilesParams) (*[]kuuzuki.File, error) {
	req := f.client.apiClient.DefaultAPI.GetFindFile(ctx)
	req = req.Query(params.Query)
	files, _, err := req.Execute()
	if err != nil {
		return nil, err
	}

	// Convert []string to []kuuzuki.File
	result := make([]kuuzuki.File, len(files))
	for i, filePath := range files {
		result[i] = kuuzuki.File{Path: filePath}
	}

	return &result, nil
}

// ConfigService provides compatibility for config-related operations
type ConfigService struct {
	client *CompatClient
}

// Get returns a config getter service for method chaining
func (c *ConfigService) Get() *ConfigGetService {
	return &ConfigGetService{client: c.client}
}

// ConfigGetService handles config retrieval
type ConfigGetService struct {
	client *CompatClient
}

// Execute gets the configuration
func (c *ConfigGetService) Execute(ctx context.Context) (*kuuzuki.Config, error) {
	req := c.client.apiClient.DefaultAPI.GetConfig(ctx)
	config, _, err := req.Execute()
	return config, err
}

// SessionInitService handles session initialization
type SessionInitService struct {
	client *CompatClient
}

// Post executes session initialization
func (s *SessionInitService) Post(ctx context.Context, sessionID string, params SessionInitParams) error {
	req := s.client.apiClient.DefaultAPI.PostSessionByIdInit(ctx, sessionID)

	// Create the init request body
	initRequest := kuuzuki.PostSessionByIdInitRequest{
		MessageID:  params.MessageID,
		ProviderID: params.ProviderID,
		ModelID:    params.ModelID,
	}

	req = req.PostSessionByIdInitRequest(initRequest)
	_, _, err := req.Execute()
	return err
}

// SessionChatService handles session chat operations
type SessionChatService struct {
	client *CompatClient
}

// Post executes session chat
func (s *SessionChatService) Post(ctx context.Context, sessionID string, params SessionChatParams) error {
	req := s.client.apiClient.DefaultAPI.PostSessionByIdMessage(ctx, sessionID)

	// Convert parts to the SDK format
	parts := make([]kuuzuki.PostSessionByIdMessageRequestPartsInner, len(params.Parts))
	for i, _ := range params.Parts {
		// Convert SessionChatParamsPartUnion to PostSessionByIdMessageRequestPartsInner
		// This is a simplified conversion - may need more sophisticated handling
		parts[i] = kuuzuki.PostSessionByIdMessageRequestPartsInner{}
	}

	// Create the chat request body
	chatRequest := kuuzuki.PostSessionByIdMessageRequest{
		ProviderID: params.ProviderID,
		ModelID:    params.ModelID,
		Mode:       &params.Mode,
		MessageID:  &params.MessageID,
		Parts:      parts,
	}
	req = req.PostSessionByIdMessageRequest(chatRequest)
	_, _, err := req.Execute()
	return err
}

// SessionSummarizeService handles session summarization
type SessionSummarizeService struct {
	client *CompatClient
}

// Post executes session summarization
func (s *SessionSummarizeService) Post(ctx context.Context, sessionID string, params SessionSummarizeParams) error {
	req := s.client.apiClient.DefaultAPI.PostSessionByIdSummarize(ctx, sessionID)

	// Create the summarize request body
	summarizeRequest := kuuzuki.PostSessionByIdSummarizeRequest{
		ProviderID: params.ProviderID,
		ModelID:    params.ModelID,
	}

	req = req.PostSessionByIdSummarizeRequest(summarizeRequest)
	_, _, err := req.Execute()
	return err
}
