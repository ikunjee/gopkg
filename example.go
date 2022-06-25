package main

import (
    "github.com/Hui4401/gopkg/errors"
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
    errors.SetUnknownCode(CodeUnknownError)
    errors.SetCode2Msg(code2msg)
    err := errors.NewCodeError(CodeParamError)
    println(errors.GetErrorCode(err))
    println(errors.GetErrorMsg(err))
}
