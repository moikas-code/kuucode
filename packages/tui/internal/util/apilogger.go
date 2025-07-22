package util

import (
	"context"
	"log/slog"
	"sync"

	kuucode "github.com/moikas-code/kuucode-sdk-go"
)

type APILogHandler struct {
	client  *kuucode.Client
	service string
	level   slog.Level
	attrs   []slog.Attr
	groups  []string
	mu      sync.Mutex
	queue   chan kuucode.AppLogParams
}

func NewAPILogHandler(ctx context.Context, client *kuucode.Client, service string, level slog.Level) *APILogHandler {
	result := &APILogHandler{
		client:  client,
		service: service,
		level:   level,
		attrs:   make([]slog.Attr, 0),
		groups:  make([]string, 0),
		queue:   make(chan kuucode.AppLogParams, 100_000),
	}
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case params := <-result.queue:
				_, err := client.App.Log(context.Background(), params)
				if err != nil {
					slog.Error("Failed to log to API", "error", err)
				}
			}
		}
	}()
	return result
}

func (h *APILogHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.level
}

func (h *APILogHandler) Handle(ctx context.Context, r slog.Record) error {
	var apiLevel kuucode.AppLogParamsLevel
	switch r.Level {
	case slog.LevelDebug:
		apiLevel = kuucode.AppLogParamsLevelDebug
	case slog.LevelInfo:
		apiLevel = kuucode.AppLogParamsLevelInfo
	case slog.LevelWarn:
		apiLevel = kuucode.AppLogParamsLevelWarn
	case slog.LevelError:
		apiLevel = kuucode.AppLogParamsLevelError
	default:
		apiLevel = kuucode.AppLogParamsLevelInfo
	}

	extra := make(map[string]any)

	h.mu.Lock()
	for _, attr := range h.attrs {
		extra[attr.Key] = attr.Value.Any()
	}
	h.mu.Unlock()

	r.Attrs(func(attr slog.Attr) bool {
		extra[attr.Key] = attr.Value.Any()
		return true
	})

	params := kuucode.AppLogParams{
		Service: kuucode.F(h.service),
		Level:   kuucode.F(apiLevel),
		Message: kuucode.F(r.Message),
	}

	if len(extra) > 0 {
		params.Extra = kuucode.F(extra)
	}

	h.queue <- params

	return nil
}

// WithAttrs returns a new Handler whose attributes consist of
// both the receiver's attributes and the arguments.
func (h *APILogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	h.mu.Lock()
	defer h.mu.Unlock()

	newHandler := &APILogHandler{
		client:  h.client,
		service: h.service,
		level:   h.level,
		attrs:   make([]slog.Attr, len(h.attrs)+len(attrs)),
		groups:  make([]string, len(h.groups)),
	}

	copy(newHandler.attrs, h.attrs)
	copy(newHandler.attrs[len(h.attrs):], attrs)
	copy(newHandler.groups, h.groups)

	return newHandler
}

// WithGroup returns a new Handler with the given group appended to
// the receiver's existing groups.
func (h *APILogHandler) WithGroup(name string) slog.Handler {
	h.mu.Lock()
	defer h.mu.Unlock()

	newHandler := &APILogHandler{
		client:  h.client,
		service: h.service,
		level:   h.level,
		attrs:   make([]slog.Attr, len(h.attrs)),
		groups:  make([]string, len(h.groups)+1),
	}

	copy(newHandler.attrs, h.attrs)
	copy(newHandler.groups, h.groups)
	newHandler.groups[len(h.groups)] = name

	return newHandler
}
