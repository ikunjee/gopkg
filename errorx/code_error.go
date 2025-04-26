package errorx

import (
	"errors"
)

type errorWithCode struct {
	error error
	code  int64
}

func (err errorWithCode) Error() string {
	if err.error != nil {
		return err.error.Error()
	}

	return GetErrorMsgByCode(err.code)
}

func NewCodeWrapError(err error, code int64) error {
	return errorWithCode{
		error: err,
		code:  code,
	}
}

func NewCodeError(code int64) error {
	return errorWithCode{
		code: code,
	}
}

func NewCodeMsgError(code int64, msg string) error {
	return errorWithCode{
		error: errors.New(msg),
		code:  code,
	}
}

func GetErrorCode(err error) int64 {
	if codeError, ok := err.(errorWithCode); ok {
		return codeError.code
	}

	return codeUnknown
}

func GetErrorMsgByCode(code int64) string {
	msg, ok := codeMsgMap[code]
	if !ok {
		return codeMsgMap[codeUnknown]
	}

	return msg
}
