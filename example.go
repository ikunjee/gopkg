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

	logs.Init(logs.LogModeDevelopment, "./output")
	defer func() {
		logs.Sync()
	}()

	errors.SetUnknownCode(CodeUnknownError)
	errors.SetUnknownMsg(code2msg[CodeUnknownError])
	errors.SetCode2MsgMap(code2msg)
	err := errors.NewCodeError(CodeParamError)
	logs.InfoKvs("error", err.Error(), "error code", errors.GetErrorCode(err))

	ctx = logs.CtxAddKVs(ctx, "func", "main")
	exampleLogFunc(ctx)
	logs.CtxInfoKvs(ctx, "hello", "hi")
}

func exampleLogFunc(ctx context.Context) {
	ctx = logs.CtxAddKVs(ctx, "func", "exampleLogFunc")
	logs.CtxInfoKvs(ctx, "hello", "world")
}
