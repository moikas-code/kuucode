package compat

import (
	kuucode "github.com/moikas-code/kuuzuki-sdk-go"
)

// Client type alias for compatibility
type Client = CompatClient

// Core types for compatibility
type App = kuucode.App
type Mode = kuucode.Mode
type Provider = kuucode.Provider
type Model = kuucode.Model
type Session = kuucode.Session
type SessionShare = kuucode.SessionShare
type Message = kuucode.Message
type UserMessage = kuucode.UserMessage
type AssistantMessage = kuucode.AssistantMessage
type TextPart = kuucode.TextPart
type FilePart = kuucode.FilePart
type ToolPart = kuucode.ToolPart
type StepStartPart = kuucode.StepStartPart
type StepFinishPart = kuucode.StepFinishPart
type FilePartSource = kuucode.FilePartSource
type ToolState = kuucode.ToolState
type AssistantMessageError = kuucode.AssistantMessageError

// Missing types that the UI code expects
type MessageUnion = interface{}
type PartUnion = interface{}
type FileReadParams struct {
	Path string `json:"path"`
}

// Config types
type KeybindsConfig = kuucode.KeybindsConfig

// Constants for compatibility
const (
	UserMessageRoleUser                           = "user"
	TextPartTypeText                              = "text"
	FilePartTypeFile                              = "file"
	FilePartSourceTypeFile                        = "file"
	FilePartSourceTypeSymbol                      = "symbol"
	SymbolSourceTypeSymbol                        = "symbol"
	FileSourceTypeFile                            = "file"
	TextPartInputTypeText                         = "text"
	FilePartInputTypeFile                         = "file"
	ConfigShareDisabled                           = "disabled"
	AssistantMessageErrorMessageOutputLengthError = "output_length_error"

	// Tool state status constants
	ToolPartStateStatusPending   = "pending"
	ToolPartStateStatusCompleted = "completed"
	ToolPartStateStatusError     = "error"
	ToolPartStateStatusRunning   = "running"
)

// App log level type
type AppLogParamsLevel string

// Session chat parameter types
type SessionChatParamsPartUnion interface {
	isSessionChatParamsPartUnion()
}

// Event types for compatibility - using type aliases to actual SDK types
type EventListResponseEventSessionUpdated = kuucode.EventSessionUpdated
type EventListResponseEventMessageUpdated = kuucode.EventMessageUpdated
type EventListResponseEventMessagePartUpdated = kuucode.EventMessagePartUpdated
type EventListResponseEventSessionDeleted = kuucode.EventSessionDeleted
type EventListResponseEventIdeInstalled = kuucode.EventIdeInstalled
type EventListResponseEventInstallationUpdated = kuucode.EventInstallationUpdated
type EventListResponseEventSessionError = kuucode.EventSessionError
type EventListResponseEventFileWatcherUpdated = kuucode.EventFileWatcherUpdated
type EventListResponseEventStorageWrite = kuucode.EventStorageWrite

// Helper function F for parameter building - matches kuucode.F signature
func F(value interface{}) interface{} {
	return value
}

// Helper function FString for string parameter building
func FString(value interface{}) string {
	if str, ok := value.(string); ok {
		return str
	}
	return ""
}

// AsUnion helper function for AssistantMessageError compatibility
func AsUnion(e *kuucode.AssistantMessageError) interface{} {
	return e
}

// ToolState compatibility wrapper
type ToolStateCompat struct {
	*kuucode.ToolState
}

// WrapToolState wraps a kuucode.ToolState to provide compatibility methods
func WrapToolState(ts *kuucode.ToolState) *ToolStateCompat {
	return &ToolStateCompat{ToolState: ts}
}

// GetPartId extracts the Id field from a Part union type
func GetPartId(part kuucode.Part) string {
	if part.TextPart != nil {
		return part.TextPart.Id
	}
	if part.FilePart != nil {
		return part.FilePart.Id
	}
	if part.ToolPart != nil {
		return part.ToolPart.Id
	}
	if part.StepStartPart != nil {
		return part.StepStartPart.Id
	}
	if part.StepFinishPart != nil {
		return part.StepFinishPart.Id
	}
	if part.SnapshotPart != nil {
		return part.SnapshotPart.Id
	}
	return ""
}

// GetPartMessageID extracts the MessageID field from a Part union type
func GetPartMessageID(part kuucode.Part) string {
	if part.TextPart != nil {
		return part.TextPart.MessageID
	}
	if part.FilePart != nil {
		return part.FilePart.MessageID
	}
	if part.ToolPart != nil {
		return part.ToolPart.MessageID
	}
	if part.StepStartPart != nil {
		return part.StepStartPart.MessageID
	}
	if part.StepFinishPart != nil {
		return part.StepFinishPart.MessageID
	}
	if part.SnapshotPart != nil {
		return part.SnapshotPart.MessageID
	}
	return ""
}

// GetPartSessionID extracts the SessionID field from a Part union type
func GetPartSessionID(part kuucode.Part) string {
	if part.TextPart != nil {
		return part.TextPart.SessionID
	}
	if part.FilePart != nil {
		return part.FilePart.SessionID
	}
	if part.ToolPart != nil {
		return part.ToolPart.SessionID
	}
	if part.StepStartPart != nil {
		return part.StepStartPart.SessionID
	}
	if part.StepFinishPart != nil {
		return part.StepFinishPart.SessionID
	}
	if part.SnapshotPart != nil {
		return part.SnapshotPart.SessionID
	}
	return ""
}

// PartAsUnion converts a Part union type to PartUnion interface
func PartAsUnion(part kuucode.Part) PartUnion {
	if part.TextPart != nil {
		return *part.TextPart
	}
	if part.FilePart != nil {
		return FilePart(*part.FilePart)
	}
	if part.ToolPart != nil {
		return *part.ToolPart
	}
	if part.StepStartPart != nil {
		return *part.StepStartPart
	}
	if part.StepFinishPart != nil {
		return *part.StepFinishPart
	}
	if part.SnapshotPart != nil {
		return *part.SnapshotPart
	}
	return nil
}

// GetMessageSessionID extracts the SessionID field from a Message union type
func GetMessageSessionID(msg kuucode.Message) string {
	if msg.AssistantMessage != nil {
		return msg.AssistantMessage.SessionID
	}
	if msg.UserMessage != nil {
		return msg.UserMessage.SessionID
	}
	return ""
}

// GetMessageId extracts the Id field from a Message union type
func GetMessageId(msg kuucode.Message) string {
	if msg.AssistantMessage != nil {
		return msg.AssistantMessage.Id
	}
	if msg.UserMessage != nil {
		return msg.UserMessage.Id
	}
	return ""
}

// MessageAsUnion converts a Message union type to MessageUnion interface
func MessageAsUnion(msg kuucode.Message) MessageUnion {
	if msg.AssistantMessage != nil {
		return *msg.AssistantMessage
	}
	if msg.UserMessage != nil {
		return *msg.UserMessage
	}
	return nil
}

// Status returns the status of the tool state for compatibility
func (ts *ToolStateCompat) Status() string {
	if ts.ToolState == nil {
		return ToolPartStateStatusPending
	}
	if ts.ToolState.ToolStateCompleted != nil {
		return ToolPartStateStatusCompleted
	}
	if ts.ToolState.ToolStateError != nil {
		return ToolPartStateStatusError
	}
	if ts.ToolState.ToolStatePending != nil {
		return ToolPartStateStatusPending
	}
	if ts.ToolState.ToolStateRunning != nil {
		return ToolPartStateStatusRunning
	}
	return ToolPartStateStatusPending
}

// Output returns the output for compatibility
func (ts *ToolStateCompat) Output() interface{} {
	if ts.ToolState != nil && ts.ToolState.ToolStateCompleted != nil {
		return ts.ToolState.ToolStateCompleted.Output
	}
	return nil
}

// Input returns the input for compatibility
func (ts *ToolStateCompat) Input() interface{} {
	if ts.ToolState != nil && ts.ToolState.ToolStateCompleted != nil {
		return ts.ToolState.ToolStateCompleted.Input
	}
	if ts.ToolState != nil && ts.ToolState.ToolStateRunning != nil {
		return ts.ToolState.ToolStateRunning.Input
	}
	return nil
}

// Error returns the error for compatibility
func (ts *ToolStateCompat) Error() interface{} {
	if ts.ToolState != nil && ts.ToolState.ToolStateError != nil {
		return ts.ToolState.ToolStateError.Error
	}
	return nil
}

// Metadata returns metadata for compatibility
func (ts *ToolStateCompat) Metadata() interface{} {
	if ts.ToolState != nil && ts.ToolState.ToolStateCompleted != nil {
		return ts.ToolState.ToolStateCompleted.Metadata
	}
	return nil
}

// Union types and interfaces for compatibility - simplified approach

// Log level constants
const (
	AppLogParamsLevelDebug = "debug"
	AppLogParamsLevelInfo  = "info"
	AppLogParamsLevelWarn  = "warn"
	AppLogParamsLevelError = "error"
)

// Parameter types for API calls
type TextPartInputParam struct {
	ID   string                  `json:"id"`
	Type string                  `json:"type"`
	Text string                  `json:"text"`
	Time *TextPartInputTimeParam `json:"time,omitempty"`
}

type TextPartInputTimeParam struct {
	Start *float64 `json:"start,omitempty"`
	End   *float64 `json:"end,omitempty"`
}

type FilePartInputParam struct {
	ID     string                    `json:"id"`
	Type   string                    `json:"type"`
	URL    string                    `json:"url"`
	Source *FilePartSourceUnionParam `json:"source,omitempty"`
}

// Simplified parameter types for compatibility
type FilePartSourceUnionParam interface {
	isFilePartSourceUnionParam()
}

type FileSourceParam struct {
	Type string                   `json:"type"`
	Path string                   `json:"path"`
	Text *FilePartSourceTextParam `json:"text,omitempty"`
}

func (f FileSourceParam) isFilePartSourceUnionParam() {}

type SymbolSourceParam struct {
	Type  string                   `json:"type"`
	Path  string                   `json:"path"`
	Name  string                   `json:"name"`
	Kind  string                   `json:"kind"`
	Range *SymbolSourceRangeParam  `json:"range,omitempty"`
	Text  *FilePartSourceTextParam `json:"text,omitempty"`
}

func (s SymbolSourceParam) isFilePartSourceUnionParam() {}

type FilePartSourceTextParam struct {
	Start *int32 `json:"start,omitempty"`
	End   *int32 `json:"end,omitempty"`
	Text  string `json:"text"`
}

type SymbolSourceRangeParam struct {
	Start *SymbolSourceRangeStartParam `json:"start,omitempty"`
	End   *SymbolSourceRangeEndParam   `json:"end,omitempty"`
}

type SymbolSourceRangeStartParam struct {
	Line      int32 `json:"line"`
	Character int32 `json:"character"`
}

type SymbolSourceRangeEndParam struct {
	Line      int32 `json:"line"`
	Character int32 `json:"character"`
}

// Additional missing types
type UserMessageTime struct {
	Start *float64 `json:"start,omitempty"`
	End   *float64 `json:"end,omitempty"`
}

type SymbolSourceRange struct {
	Start *SymbolSourceRangeStart `json:"start,omitempty"`
	End   *SymbolSourceRangeEnd   `json:"end,omitempty"`
}

type SymbolSourceRangeStart struct {
	Line      int32 `json:"line"`
	Character int32 `json:"character"`
}

type SymbolSourceRangeEnd struct {
	Line      int32 `json:"line"`
	Character int32 `json:"character"`
}

// File read parameters - already defined above

// Service parameter types
type AppLogParams struct {
	Level   string      `json:"level"`
	Service string      `json:"service"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type AppProvidersResponse struct {
	Providers []kuucode.Provider `json:"providers"`
	Default   map[string]string  `json:"default,omitempty"`
}

type SessionInitParams struct {
	MessageID  string `json:"messageId"`
	ProviderID string `json:"providerId"`
	ModelID    string `json:"modelId"`
}

type SessionChatParams struct {
	ProviderID string                       `json:"providerId"`
	ModelID    string                       `json:"modelId"`
	Mode       string                       `json:"mode"`
	MessageID  string                       `json:"messageId"`
	Parts      []SessionChatParamsPartUnion `json:"parts"`
}

type SessionSummarizeParams struct {
	ProviderID string `json:"providerId"`
	ModelID    string `json:"modelId"`
}

type FindFilesParams struct {
	Query string `json:"query"`
}

type FindFilesResponse struct {
	Files []kuucode.File `json:"files"`
}

type FindSymbolsParams struct {
	Query string `json:"query"`
}

type FindSymbolsResponse struct {
	Symbols []kuucode.Symbol `json:"symbols"`
}
