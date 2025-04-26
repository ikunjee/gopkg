package errorx

import (
	"errors"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

var (
	originalErr = errors.New("original error")
)

func TestNewError(t *testing.T) {
	initCodeMsg()

	convey.Convey("NewCodeWrapError", t, func() {
		convey.Convey("nil origin error", func() {
			err := NewCodeWrapError(nil, 100)
			codeErr, ok := err.(errorWithCode)
			convey.So(ok, convey.ShouldBeTrue)
			convey.So(codeErr.code, convey.ShouldEqual, 100)
			convey.So(codeErr.error, convey.ShouldEqual, nil)
		})
		convey.Convey("origin error", func() {
			err := NewCodeWrapError(originalErr, 100)
			codeErr, ok := err.(errorWithCode)
			convey.So(ok, convey.ShouldBeTrue)
			convey.So(codeErr.code, convey.ShouldEqual, 100)
			convey.So(codeErr.error, convey.ShouldEqual, originalErr)
		})
	})

	convey.Convey("NewCodeError", t, func() {
		err := NewCodeError(100)
		codeErr, ok := err.(errorWithCode)
		convey.So(ok, convey.ShouldBeTrue)
		convey.So(codeErr.code, convey.ShouldEqual, 100)
		convey.So(codeErr.error, convey.ShouldEqual, nil)
	})

	convey.Convey("NewCodeMsgError", t, func() {
		err := NewCodeMsgError(100, "test message")
		codeErr, ok := err.(errorWithCode)
		convey.So(ok, convey.ShouldBeTrue)
		convey.So(codeErr.code, convey.ShouldEqual, 100)
		convey.So(codeErr.error.Error(), convey.ShouldEqual, "test message")
	})
}

func TestGetError(t *testing.T) {
	initCodeMsg()

	convey.Convey("default config", t, func() {
		convey.Convey("wrap error", func() {
			err := NewCodeWrapError(originalErr, 100)
			convey.So(GetErrorCode(err), convey.ShouldEqual, 100)
			convey.So(err.Error(), convey.ShouldEqual, originalErr.Error())
		})
		convey.Convey("code error", func() {
			err := NewCodeError(100)
			convey.So(GetErrorCode(err), convey.ShouldEqual, 100)
			convey.So(err.Error(), convey.ShouldEqual, msgUnknown)
		})
		convey.Convey("code msg error", func() {
			err := NewCodeMsgError(100, "test message")
			convey.So(GetErrorCode(err), convey.ShouldEqual, 100)
			convey.So(err.Error(), convey.ShouldEqual, "test message")
		})
		convey.Convey("get msg by code", func() {
			convey.So(GetErrorMsgByCode(100), convey.ShouldEqual, msgUnknown)
			convey.So(GetErrorMsgByCode(-1), convey.ShouldEqual, msgUnknown)
		})
	})

	convey.Convey("config unknown", t, func() {
		initCodeMsg()
		SetUnknownCode(-100)
		SetUnknownMsg("unknown!")
		convey.Convey("code error", func() {
			err := NewCodeError(100)
			convey.So(GetErrorCode(err), convey.ShouldEqual, 100)
			convey.So(err.Error(), convey.ShouldEqual, "unknown!")
		})
		convey.Convey("get msg by code", func() {
			convey.So(GetErrorMsgByCode(100), convey.ShouldEqual, "unknown!")
			convey.So(GetErrorMsgByCode(-100), convey.ShouldEqual, "unknown!")
		})
	})

	convey.Convey("config code msg map", t, func() {
		initCodeMsg()
		SetCodeMsgMap(map[int64]string{
			-1:  "unknown!",
			100: "test message",
		})
		convey.Convey("code error", func() {
			err := NewCodeError(100)
			convey.So(GetErrorCode(err), convey.ShouldEqual, 100)
			convey.So(err.Error(), convey.ShouldEqual, "test message")
		})
		convey.Convey("get msg by code", func() {
			convey.So(GetErrorMsgByCode(100), convey.ShouldEqual, "test message")
			convey.So(GetErrorMsgByCode(-100), convey.ShouldEqual, "unknown error")
			convey.So(GetErrorMsgByCode(-1), convey.ShouldEqual, "unknown!")
		})
	})
}

func initCodeMsg() {
	SetUnknownCode(-1)
	SetUnknownMsg("unknown error")
}
