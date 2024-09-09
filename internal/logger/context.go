package logger

import (
	"context"
	"log/slog"
	"slices"
	"strings"
)

type ctxKey int

const logCtxKey ctxKey = 0

type ctxKV map[string]string

// LogValue возвращает данные для записи в лог
func (c ctxKV) LogValue() slog.Value {
	attrs := make([]slog.Attr, 0, len(c))

	for k, v := range c {
		attrs = append(attrs, slog.String(k, v))
	}

	slices.SortFunc(attrs, func(a, b slog.Attr) int {
		return strings.Compare(a.Key, b.Key)
	})

	return slog.GroupValue(attrs...)
}

// LogWithKvContext добавляет kv в контекст для логирования
func LogWithKvContext(ctx context.Context, key, value string) context.Context {
	c, ok := ctx.Value(logCtxKey).(ctxKV)
	if !ok {
		c = ctxKV{}
	}

	c[key] = value

	return context.WithValue(ctx, logCtxKey, c)
}

// LogWithKvMapContext добавляет map в контекст для логирования
func LogWithKvMapContext(ctx context.Context, kv map[string]string) context.Context {
	c, ok := ctx.Value(logCtxKey).(ctxKV)
	if !ok {
		c = ctxKV{}
	}

	for k, v := range kv {
		c[k] = v
	}

	return context.WithValue(ctx, logCtxKey, c)
}
