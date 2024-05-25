package main

import (
	"context"

	"github.com/Hui4401/gopkg/errorx"
	"github.com/Hui4401/gopkg/logx"
)

type errorCode = int32

const (
	CodeUnknownError errorCode = -1
	CodeParamError   errorCode = 1001
)

var code2msg = map[errorCode]string{
	CodeUnknownError: "未知错误",
	CodeParamError:   "参数错误",
}

func main() {
	ctx := context.Background()

	// 初始化 logs，设置最低输出级别限制，以及同时写文件路径
	logx.Init(logx.LogModeDevelopment, "./output", true)
	defer func() {
		logx.Sync()
	}()

	// 初始化 errors，code 使用
	errorx.SetUnknownCode(CodeUnknownError)
	errorx.SetUnknownMsg(code2msg[CodeUnknownError])
	errorx.SetCode2MsgMap(code2msg)
	err := errorx.NewCodeError(CodeParamError)
	logx.InfoKvs("error example", "err", err, "code", errorx.GetErrorCode(err))

	// logs 在 context 中传递键值对到下一级方法
	ctx = logx.CtxAddKVs(ctx, "func", "main")
	exampleLogFunc(ctx)
	logx.CtxInfoKvs(ctx, "test ctx", "hello", "hi")
}

func exampleLogFunc(ctx context.Context) {
	ctx = logx.CtxAddKVs(ctx, "func", "exampleLogFunc")
	logx.CtxInfoKvs(ctx, "test ctx", "hello", "world")
}
