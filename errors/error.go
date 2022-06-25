package errors

type errorWithCode struct {
    error
    code int32
}

func (err errorWithCode) Error() string {
    return err.error.Error()
}

func NewCodeError(code int32) error {
    return errorWithCode{
        code: code,
    }
}

func GetErrorCode(err error) int32 {
    if codeError, ok := err.(errorWithCode); ok {
        return codeError.code
    }
    return unknownCode
}

func GetErrorMsg(err error) string {
    if codeError, ok := err.(errorWithCode); ok {
        return getMsg(codeError.code)
    }
    return unknownMsg
}

func getMsg(errCode int32) string {
    msg, ok := msgMap[errCode]
    if !ok {
        return unknownMsg
    }
    return msg
}
