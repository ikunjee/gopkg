package logs

import (
	"context"
)

func Sync() {
	_ = logger.Sync()
}

func DebugKvs(kvs ...interface{}) {
	logger.Debugw("", kvs...)
}

func InfoKvs(kvs ...interface{}) {
	logger.Infow("", kvs...)
}

func WarnKvs(kvs ...interface{}) {
	logger.Warnw("", kvs...)
}

func ErrorKvs(kvs ...interface{}) {
	logger.Errorw("", kvs...)
}

func PanicKvs(kvs ...interface{}) {
	logger.Panicw("", kvs...)
}

func CtxDebugKvs(ctx context.Context, kvs ...interface{}) {
	logger.Debugw("", kvs...)
}

func CtxInfoKvs(ctx context.Context, kvs ...interface{}) {
	logger.Infow("", kvs...)
}

func CtxWarnKvs(ctx context.Context, kvs ...interface{}) {
	logger.Warnw("", kvs...)
}

func CtxErrorKvs(ctx context.Context, kvs ...interface{}) {
	logger.Errorw("", kvs...)
}

func CtxPanicKvs(ctx context.Context, kvs ...interface{}) {
	logger.Panicw("", kvs...)
}
