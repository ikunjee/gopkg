package asyncx

import (
	"context"
	"log/slog"
	"runtime/debug"
)

// Recover 捕获 panic 并打印错误堆栈信息
// 参数 ctx 可选，不传则使用 context.Background()
// 注意要直接跟在 defer 后面，不要嵌套多余 func，否则不生效
func Recover(ctx ...context.Context) {
	if err := recover(); err != nil {
		var c context.Context
		if len(ctx) > 0 {
			c = ctx[0]
		} else {
			c = context.Background()
		}
		slog.ErrorContext(c, "panic recovered", "err", err)
		debug.PrintStack()
	}
}
