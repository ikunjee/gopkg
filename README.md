# gopkg

一些 Golang 开发时常用组件的封装

## Installation

`go get github.com/Hui4401/gopkg`

## 组件

使用示例见 [example.go](./example.go)

### errors

在原始 error 的基础上添加了 code 属性，便于服务开发时的错误码传递

### logs

基于 go.uber.org/zap 封装的日志组件

- 方便的为 context 添加 log kv 信息，在打印日志时输出调用链上所有自定义添加的日志
