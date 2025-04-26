### logs

基于 [go.uber.org/zap](https://pkg.go.dev/go.uber.org/zap) 封装的日志组件

- CtxXXX 方法：为 context 添加 kvs 信息，在打印日志时输出调用链上所有自定义信息，适合在复杂调用中追踪上下文，使用示例：

```go
package main
import (
	"context"
	"github.com/ikunjee/gopkg/logx"
)
func funcA(ctx context.Context) {
    logx.CtxInfo(ctx, "log in %s", "funcA")                    // log in funcA    {"func": "main"}
	logx.CtxInfoKvs(ctx, "log in funcA", "test", "funcA")  // log in funcA    {"func": "main", "test": "funcA"}
}
func main() {
    ctx := context.Background()
    ctx = logx.CtxAddKVs(ctx, "func", "main")
	funcA(ctx)  // ctx 传递
}
```

- LogXXX 方法：通用方法直接打印日志：

```go
package main
import (
	"context"
	"github.com/ikunjee/gopkg/logx"
)
func main() {
	logx.Info("log in %s", "main")               // log in main
	logx.InfoKvs("log in main", "test", "main")  // log in main     {"test": "main"}
}
```

- InitDefaultLogger: 配置日志级别，输出到文件，启用color等

```go
package main
import (
	"github.com/ikunjee/gopkg/logx"
)
func main() {
	// 默认 product 级别只打印 info 及以上日志，不输出到文件，不启用 color，可通过 InitDefaultLogger 自定义设置
	logx.InitDefaultLogger(logx.LogModeDevelopment, "logx/out", true)
	logx.Info("hello info")
	logx.Error("hello error")
}
```
