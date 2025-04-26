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
