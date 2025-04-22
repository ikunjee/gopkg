# asyncx

异步任务处理

## 下载

`go get github.com/ikunjee/gopkg`

## 引入
```go
import "github.com/ikunjee/gopkg/asyncx"
```

## 方法列表
| 方法名称   | 功能描述         | 详细说明               |
|---|--------------|--------------------|
| Recover | 捕获panic并打印堆栈 | [link](###Recover) |

## 详细说明

### Recover

> 用于捕获 panic 并打印错误堆栈信息
> 
> 参数 ctx 可选，最多传一个，不传则使用 context.Background()
> 
> 注意要直接跟在 defer 后面，不要嵌套多余 func，否则不生效

- 示例

```go
package main
import (
	"context"
	"fmt"
	"sync"
	"github.com/ikunjee/gopkg/asyncx"
)
func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer asyncx.Recover(ctx)
		/* 错误用法
		defer func() {
		    asyncx.Recover(ctx)
		}()
		*/
		defer wg.Done()
		panic("test panic")
	}()
	wg.Wait()
	fmt.Println("goroutine has been recovered")
}
```
