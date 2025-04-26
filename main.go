package main

import (
	"github.com/ikunjee/gopkg/logx"
)

func main() {
	// 默认 product 级别只打印 info 及以上日志，不输出到文件，不启用 color，可通过 InitDefaultLogger 自定义设置
	logx.InitDefaultLogger(logx.LogModeDevelopment, "logx/output", true)
	logx.Info("hello info")
	logx.Error("hello error")
}
