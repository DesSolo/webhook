package logger

import (
	"context"
	"log/slog"

	"github.com/go-chi/chi/v5/middleware"
)

// LogContextHandler wraps slog.Handler with custom context
type LogContextHandler struct {
	slog.Handler
}

// NewLogContextHandler creates new LogContextHandler
func NewLogContextHandler(h slog.Handler) *LogContextHandler {
	return &LogContextHandler{
		Handler: h,
	}
}

// Handle handles slog record and adds reqID and kv
func (h LogContextHandler) Handle(ctx context.Context, r slog.Record) error {
	if reqID := middleware.GetReqID(ctx); reqID != "" {
		r.Add("reqID", reqID)
	}

	if kv, ok := ctx.Value(logCtxKey).(ctxKV); ok {
		r.Add("kv", kv)
	}

	return h.Handler.Handle(ctx, r)
}

// WithAttrs returns a Handler that wraps h in a group with the given attributes.
func (h LogContextHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return LogContextHandler{
		Handler: h.Handler.WithAttrs(attrs),
	}
}

// WithGroup returns a Handler that wraps h in a group with the given name.
func (h LogContextHandler) WithGroup(name string) slog.Handler {
	return LogContextHandler{
		Handler: h.Handler.WithGroup(name),
	}
}
