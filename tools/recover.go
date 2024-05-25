package tools

import (
	"context"
	"log/slog"
	"runtime/debug"
)

func Recover(ctx context.Context) {
	if e := recover(); e != nil {
		slog.ErrorContext(ctx, "panic recovered", "err", e, "stack", string(debug.Stack()))
	}
}
