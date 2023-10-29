package main

import (
	"context"

	"github.com/Hui4401/gopkg/errors"
	"github.com/Hui4401/gopkg/logs"
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
	logs.Init(logs.LogModeDevelopment, "./output")
	defer func() {
		logs.Sync()
	}()

	// 初始化 errors，code 使用
	errors.SetUnknownCode(CodeUnknownError)
	errors.SetUnknownMsg(code2msg[CodeUnknownError])
	errors.SetCode2MsgMap(code2msg)
	err := errors.NewCodeError(CodeParamError)
	logs.InfoKvs("error example", "err", err, "code", errors.GetErrorCode(err))

	// logs 在 context 中传递键值对到下一级方法
	ctx = logs.CtxAddKVs(ctx, "func", "main")
	exampleLogFunc(ctx)
	logs.CtxInfoKvs(ctx, "test ctx", "hello", "hi")
}

func exampleLogFunc(ctx context.Context) {
	ctx = logs.CtxAddKVs(ctx, "func", "exampleLogFunc")
	logs.CtxInfoKvs(ctx, "test ctx", "hello", "world")
}
