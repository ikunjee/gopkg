package logx

func Debug(template string, args ...any) {
	defaultLogger.Debugf(template, args...)
}

func Info(template string, args ...any) {
	defaultLogger.Infof(template, args...)
}

func Warn(template string, args ...any) {
	defaultLogger.Warnf(template, args...)
}

func Error(template string, args ...any) {
	defaultLogger.Errorf(template, args...)
}

func Panic(template string, args ...any) {
	defaultLogger.Panicf(template, args...)
}

func DebugKvs(msg string, kvs ...any) {
	defaultLogger.Debugw(msg, kvs...)
}

func InfoKvs(msg string, kvs ...any) {
	defaultLogger.Infow(msg, kvs...)
}

func WarnKvs(msg string, kvs ...any) {
	defaultLogger.Warnw(msg, kvs...)
}

func ErrorKvs(msg string, kvs ...any) {
	defaultLogger.Errorw(msg, kvs...)
}

func PanicKvs(msg string, kvs ...any) {
	defaultLogger.Panicw(msg, kvs...)
}
