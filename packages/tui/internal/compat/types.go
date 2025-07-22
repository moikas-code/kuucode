package compat

import (
	kuucode "github.com/moikas-code/kuucode-sdk-go"
)

// Client type alias for compatibility
type Client = Client

// Parameter types for compatibility
type AppLogParams struct {
	Level   AppLogParamsLevel `json:"level"`
	Service string            `json:"service"`
	Message string            `json:"message"`
	Data    interface{}       `json:"data,omitempty"`
}

type AppLogParamsLevel string

const (
	AppLogParamsLevelDebug AppLogParamsLevel = "DEBUG"
	AppLogParamsLevelInfo  AppLogParamsLevel = "INFO"
	AppLogParamsLevelWarn  AppLogParamsLevel = "WARN"
	AppLogParamsLevelError AppLogParamsLevel = "ERROR"
)

type SessionInitParams struct {
	Provider string `json:"provider"`
	Model    string `json:"model"`
	Mode     string `json:"mode,omitempty"`
}

type SessionChatParams struct {
	Provider string                       `json:"provider"`
	Model    string                       `json:"model"`
	Parts    []SessionChatParamsPartUnion `json:"parts"`
}

type SessionChatParamsPartUnion interface {
	isSessionChatParamsPartUnion()
}

type SessionSummarizeParams struct {
	Provider string `json:"provider"`
	Model    string `json:"model"`
}

type FileReadParams struct {
	Path string `json:"path"`
}

type FindFilesParams struct {
	Query string `json:"query"`
}

type FindSymbolsParams struct {
	Query string `json:"query"`
}

// Response types for compatibility
type AppProvidersResponse struct {
	Providers []kuucode.Provider `json:"providers"`
}

type FileReadResponse struct {
	Content string `json:"content"`
}

type FindFilesResponse struct {
	Files []kuucode.File `json:"files"`
}

type FindSymbolsResponse struct {
	Symbols []kuucode.Symbol `json:"symbols"`
}

// Union types for compatibility
type MessageUnion interface {
	isMessageUnion()
}

type PartUnion interface {
	isPartUnion()
}

// Event types for compatibility
type EventListResponseEventSessionUpdated struct{}
type EventListResponseEventMessageUpdated struct{}
type EventListResponseEventMessagePartUpdated struct{}
type EventListResponseEventInstallationUpdated struct{}
type EventListResponseEventIdeInstalled struct{}
type EventListResponseEventSessionDeleted struct{}
type EventListResponseEventSessionError struct{}
type EventListResponseEventFileWatcherUpdated struct{}

// Tool state constants for compatibility
const (
	ToolPartStateStatusError     = "error"
	ToolPartStateStatusPending   = "pending"
	ToolPartStateStatusCompleted = "completed"
)

// Config constants for compatibility
const (
	ConfigShareDisabled = "disabled"
)

// Message type constants for compatibility
const (
	UserMessageRoleUser = "user"
)

// Part type constants for compatibility
const (
	TextPartTypeText         = "text"
	FilePartTypeFile         = "file"
	TextPartInputTypeText    = "text"
	FilePartInputTypeFile    = "file"
	FilePartSourceTypeFile   = "file"
	FilePartSourceTypeSymbol = "symbol"
	SymbolSourceTypeSymbol   = "symbol"
	FileSourceTypeFile       = "file"
)

// Error types for compatibility
type AssistantMessageErrorMessageOutputLengthError struct{}

// Helper function F for compatibility
func F(value interface{}) interface{} {
	return value
}

// Time types for compatibility
type UserMessageTime struct {
	Start *float32 `json:"start,omitempty"`
	End   *float32 `json:"end,omitempty"`
}

type TextPartInputTimeParam struct {
	Start *float32 `json:"start,omitempty"`
	End   *float32 `json:"end,omitempty"`
}

// Input parameter types for compatibility
type TextPartInputParam struct {
	ID   string                  `json:"id"`
	Type string                  `json:"type"`
	Text string                  `json:"text"`
	Time *TextPartInputTimeParam `json:"time,omitempty"`
}

type FilePartInputParam struct {
	ID       string                    `json:"id"`
	Type     string                    `json:"type"`
	Filename string                    `json:"filename"`
	URL      string                    `json:"url"`
	Source   *FilePartSourceUnionParam `json:"source,omitempty"`
}

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
	Start   string `json:"start"`
	End     string `json:"end"`
	Content string `json:"content"`
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

// Implement union interfaces for existing types
func (kuucode.UserMessage) isMessageUnion()      {}
func (kuucode.AssistantMessage) isMessageUnion() {}

func (kuucode.TextPart) isPartUnion()       {}
func (kuucode.FilePart) isPartUnion()       {}
func (kuucode.ToolPart) isPartUnion()       {}
func (kuucode.StepStartPart) isPartUnion()  {}
func (kuucode.StepFinishPart) isPartUnion() {}

func (TextPartInputParam) isSessionChatParamsPartUnion() {}
func (FilePartInputParam) isSessionChatParamsPartUnion() {}

// Compatibility wrappers for existing types to add missing fields/methods
type SessionWrapper struct {
	*kuucode.Session
}

func (s *SessionWrapper) ID() string {
	if s.Session != nil && s.Session.Id != nil {
		return *s.Session.Id
	}
	return ""
}

type ProviderWrapper struct {
	*kuucode.Provider
}

func (p *ProviderWrapper) ID() string {
	if p.Provider != nil && p.Provider.Id != nil {
		return *p.Provider.Id
	}
	return ""
}

type ModelWrapper struct {
	*kuucode.Model
}

func (m *ModelWrapper) ID() string {
	if m.Model != nil && m.Model.Id != nil {
		return *m.Model.Id
	}
	return ""
}

// Extended types with compatibility methods
type ToolStateExtended struct {
	kuucode.ToolState
}

func (t *ToolStateExtended) Status() string {
	// TODO: Map from new SDK tool state to old status format
	return "pending"
}

func (t *ToolStateExtended) Output() interface{} {
	// TODO: Extract output from new SDK tool state
	return nil
}

func (t *ToolStateExtended) Input() interface{} {
	// TODO: Extract input from new SDK tool state
	return nil
}

func (t *ToolStateExtended) Error() interface{} {
	// TODO: Extract error from new SDK tool state
	return nil
}

func (t *ToolStateExtended) Metadata() interface{} {
	// TODO: Extract metadata from new SDK tool state
	return nil
}
