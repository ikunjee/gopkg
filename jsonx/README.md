# jsonx

常用 json 方法封装，设计要点：
- 提供更多简洁实用的方法，减少日常冗余代码，如 ToString 忽略错误，泛型方法 UnmarshalStringWithType 无需提前定义结构等
- 底层封装标准库或流行第三方库，对外提供统一方法可自由选择，目前已支持：
  - （默认）Bytedance 开源的 [sonic](https://github.com/bytedance/sonic)
  - 标准库 encoding/json

- 选项模式（Option）提供灵活的配置方式，方法级配置 > 全局配置 > 默认配置，可根据需求灵活选择，[配置方法](#配置方法)



## 下载

`go get github.com/ikunjee/gopkg`

## 引入
```go
import "github.com/ikunjee/gopkg/jsonx"
```

## Json方法
| 方法名称   | 功能描述                           | 详细说明               |
|---|--------------------------------|--------------------|
| ToString | 序列化为 string，忽略错误               | [link](###ToString) |
| ToByte | 序列化为 bytes，忽略错误                | [link](###ToByte) |
| MarshalString | 常规 MarshalString               | [link](###MarshalString) |
| Marshal | 常规 Marshal                     | [link](###Marshal) |
| UnmarshalStringWithType | 泛型 UnmarshalString 方法，无需提前定义结构 | [link](###UnmarshalStringWithType) |
| UnmarshalWithType | 泛型 Unmarshal 方法，无需提前定义结构       | [link](###UnmarshalWithType) |
| UnmarshalString | 常规 UnmarshalString             | [link](###UnmarshalString) |
| Unmarshal | 常规 Unmarshal                   | [link](###Unmarshal) |
| ConfigDefault | 全局配置               | [link](###ConfigDefault) |

## 详细说明

### ToString

> 日常绝大多数场景下，json marshal 并不会出错，直接 ToString 即可，除非序列化一个不能序列化的对象如 channel

- 示例

```go
package main
import (
	"fmt"
	"github.com/ikunjee/gopkg/jsonx"
)
func main() {
    fmt.Println(jsonx.ToString(map[string]string{"a": "123"}))
}
```

### ToByte

> 日常绝大多数场景下，json marshal 并不会出错，直接 ToByte 即可，除非序列化一个不能序列化的对象如 channel

- 示例

```go
package main
import (
	"fmt"
	"github.com/ikunjee/gopkg/jsonx"
)
func main() {
    fmt.Println(jsonx.ToByte(map[string]string{"a": "123"}))
}
```

### MarshalString

```go
package main
import (
	"fmt"
	"github.com/ikunjee/gopkg/jsonx"
)
func main() {
	result, err := jsonx.MarshalString(123)
}
```

### Marshal

```go
package main
import (
	"fmt"
	"github.com/ikunjee/gopkg/jsonx"
)
func main() {
	result, err := jsonx.Marshal([]byte("123"))
}
```

### UnmarshalStringWithType

> 泛型反序列化方法，将字符串反序列化为指定类型，无需提前定义结构，传入空字符串时返回零值而不是报错。

- 示例

```go
package main
import (
	"fmt"
	"github.com/ikunjee/gopkg/jsonx"
)
func main() {
	jsonStr := `{"key": "value"}`
	result, err := jsonx.UnmarshalStringWithType[map[string]any](jsonStr)
}
```

### UnmarshalWithType

> 泛型反序列化方法，将字节切片反序列化为指定类型，无需提前定义结构，传入空字节切片时返回零值而不是报错。

- 示例

```go
package main
import (
	"fmt"
	"github.com/ikunjee/gopkg/jsonx"
)
func main() {
	jsonBytes := []byte(`{"key": "value"}`)
	result, err := jsonx.UnmarshalWithType[map[string]any](jsonBytes)
}
```

### UnmarshalString

```go
package main
import (
	"fmt"
	"github.com/ikunjee/gopkg/jsonx"
)
func main() {
	jsonStr := `{"key": "value"}`
	var anyMap map[string]any
	err := jsonx.UnmarshalString(jsonStr, &anyMap)
}
```

### Unmarshal

```go
package main
import (
	"fmt"
	"github.com/ikunjee/gopkg/jsonx"
)
func main() {
	jsonBytes := []byte(`{"key": "value"}`)
	var anyMap map[string]any
	err := jsonx.Unmarshal(jsonBytes, &anyMap)
}
```

## 配置方法

> 使用选项模式（option）实现用不同的配置控制 json 行为， 优先级：临时配置 > 全局配置 > 默认配置

### 可选options
- **func WithJsonLibType(jsonLibType)**：切换底层 json 库，默认值：sonic
- **WithEscapeHTML(bool)**：marshal 时是否转义 HTML 字符，默认 false
- **WithUseNumber(bool)**：unmarshal 数字时是否转成 number 类型，默认值：false

### 临时配置

> json方法调用时可传入option，仅对当前方法生效，方法级配置优先级高于全局配置

- 示例

```go
package main
import (
	"fmt"
	"github.com/ikunjee/gopkg/jsonx"
)
func main() {
    fmt.Println(jsonx.ToString("</>", jsonx.WithEscapeHTML(true)))
}
```

### 全局配置

> 使用 ConfigDefault 方法设置全局 json 行为

- 示例

```go
package main
import (
	"fmt"
	"github.com/ikunjee/gopkg/jsonx"
)
func main() {
    jsonx.ConfigDefault(WithEscapeHTML(true), WithUseNumber(true))  // html字符转义、数字使用 number 类型，全局生效
}
```

### 默认配置

若未设置过全局配置，且方法调用时也未传入 option，则使用默认配置：

- jsonLibType：sonic
- escapeHTML：false
- useNumber：false