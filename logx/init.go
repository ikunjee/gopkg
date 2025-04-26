package logx

import (
	"log"
	"log/slog"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type logMode int

const (
	LogModeDevelopment logMode = 1
	LogModeProduct     logMode = 2
)

var defaultLogger *zap.SugaredLogger

func init() {
	InitDefaultLogger(LogModeProduct, "", true)
}

// InitDefaultLogger 初始化默认logger， logPath 指定同时输出到文件，为空则不输出
func InitDefaultLogger(logMode logMode, logPath string, enableColor bool) {
	var zapConfig zap.Config
	switch logMode {
	case LogModeDevelopment:
		zapConfig = zap.NewDevelopmentConfig()
	case LogModeProduct:
		zapConfig = zap.NewProductionConfig()
	}

	zapConfig.Encoding = "console"
	zapConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
	zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	if enableColor {
		zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	if logPath != "" {
		if _, err := os.Stat(logPath); os.IsNotExist(err) {
			if err = os.MkdirAll(logPath, os.ModePerm); err != nil {
				slog.Error("creat log path err: %v", err)
				panic(err)
			}
		}
		zapConfig.OutputPaths = append(zapConfig.OutputPaths, logPath+"/info.log")
		zapConfig.ErrorOutputPaths = append(zapConfig.ErrorOutputPaths, logPath+"/error.log")
	}

	_logger, err := zapConfig.Build(zap.AddCallerSkip(1), zap.AddStacktrace(zap.PanicLevel))
	if err != nil {
		log.Panicf("InitDefaultLogger zap defaultLogger error: %s", err.Error())
	}

	defaultLogger = _logger.Sugar()
}

func Sync() {
	_ = defaultLogger.Sync()
}
