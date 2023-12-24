package errors

import (
	"errors"
)

type errorWithCode struct {
	error
	code int32
}

func WithCode(err error, code int32) error {
	return errorWithCode{
		error: err,
		code:  code,
	}
}

func NewCodeError(code int32) error {
	return errorWithCode{
		code: code,
	}
}

func NewCodeErrorWithMsg(code int32, msg string) error {
	return errorWithCode{
		error: errors.New(msg),
		code:  code,
	}
}

func (err errorWithCode) Error() string {
	if err.error != nil {
		return err.error.Error()
	}

	return GetErrorMsgByCode(err.code)
}

func GetErrorCode(err error) int32 {
	if codeError, ok := err.(errorWithCode); ok {
		return codeError.code
	}

	return codeUnknown
}

func GetErrorMsgByCode(code int32) string {
	if code2msg == nil {
		return msgUnknown
	}
	msg, ok := code2msg[code]
	if !ok {
		return code2msg[codeUnknown]
	}

	return msg
}
