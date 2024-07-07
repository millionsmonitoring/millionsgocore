package logger

import (
	"context"
	"log/slog"
	"os"
)

func Info(ctx context.Context, msg string, args ...any) {
	slog.InfoContext(ctx, msg, args...)
}

func Error(ctx context.Context, msg string, args ...any) {
	slog.ErrorContext(ctx, msg, args...)
}

func XError(ctx context.Context, msg string, args ...any) {
	slog.ErrorContext(ctx, msg, args...)
	// add a sentry or some monitor in the application
}

func Warn(ctx context.Context, msg string, args ...any) {
	slog.WarnContext(ctx, msg, args...)
}

func Debug(ctx context.Context, msg string, args ...any) {
	slog.DebugContext(ctx, msg, args...)
}

func Fatal(ctx context.Context, msg string, args ...any) {
	slog.ErrorContext(ctx, msg, args...)
	os.Exit(1)
}
