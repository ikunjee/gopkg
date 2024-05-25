package logx

func DebugKvs(msg string, kvs ...any) {
	logger.Debugw(msg, kvs...)
}

func InfoKvs(msg string, kvs ...any) {
	logger.Infow(msg, kvs...)
}

func WarnKvs(msg string, kvs ...any) {
	logger.Warnw(msg, kvs...)
}

func ErrorKvs(msg string, kvs ...any) {
	logger.Errorw(msg, kvs...)
}

func PanicKvs(msg string, kvs ...any) {
	logger.Panicw(msg, kvs...)
}
