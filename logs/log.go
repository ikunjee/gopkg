package logs

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
