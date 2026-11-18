package logger

import (
	"context"
	"log/slog"
	"os"
)

var Log *slog.Logger

const CorrelationIDKey = "correlation_id"

func InitLogger() {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	Log = slog.New(handler)

}

func getLogArgs(ctx context.Context, args []any) []any {
	if ctx != nil {
		if cid, ok := ctx.Value(CorrelationIDKey).(string); ok {
			return append(args, slog.String(CorrelationIDKey, cid))
		}
	}
	return args
}


func Info(ctx context.Context, msg string, args ...any) {
	Log.InfoContext(ctx, msg, getLogArgs(ctx, args)...)
}

func Error(ctx context.Context, msg string, args ...any) {
	Log.ErrorContext(ctx, msg, getLogArgs(ctx, args)...)
}


func Warn(ctx context.Context, msg string, args ...any) {
	Log.WarnContext(ctx, msg, getLogArgs(ctx, args)...)
}
