package logger

import "go.uber.org/zap"

var logger *zap.Logger

func GetLogger() *zap.SugaredLogger {
	if logger == nil {
		logger, _ = zap.NewDevelopment()
		defer logger.Sync()
	}
	sugar := logger.Sugar()
	return sugar
}
