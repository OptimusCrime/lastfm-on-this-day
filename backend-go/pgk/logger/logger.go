package logger

import (
	"context"
	"log/slog"
)

func FromContext(ctx context.Context) *slog.Logger {
	if l, ok := ctx.Value("logger").(*slog.Logger); ok {
		return l
	}
	return slog.Default()
}
