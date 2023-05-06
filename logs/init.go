package logs

import (
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type logMode int

const (
	LogModeDevelopment logMode = 1
	LogModeProduct     logMode = 2
)

var logger *zap.SugaredLogger

// Init logPath为空则不同时输出到文件
func Init(logMode logMode, logPath string) {
	var zapConfig zap.Config

	switch logMode {
	case LogModeDevelopment:
		zapConfig = zap.NewDevelopmentConfig()
	case LogModeProduct:
		zapConfig = zap.NewProductionConfig()
	}

	zapConfig.Encoding = "console"
	zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	zapConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")

	// 同时输出到文件
	if logPath != "" {
		if _, err := os.Stat(logPath); os.IsNotExist(err) {
			if err = os.MkdirAll(logPath, os.ModePerm); err != nil {
				log.Panicf("creat log path error: %s", err.Error())
			}
		}
		zapConfig.OutputPaths = append(zapConfig.OutputPaths, logPath+"/info.log")
		zapConfig.ErrorOutputPaths = append(zapConfig.ErrorOutputPaths, logPath+"/error.log")
	}

	_logger, err := zapConfig.Build(
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zap.PanicLevel),
	)
	if err != nil {
		log.Panicf("Init zap logger error: %s", err.Error())
	}
	logger = _logger.Sugar()
}
