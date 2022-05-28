package logger

import (
	"go.uber.org/zap"
)

var logger *Log

func InitLogger(config *LogConfig) {
	logger = &Log{}
	logger.init(config)
}


func Info(msg string, fields ...zap.Field) {
	logger.l.Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.l.Error(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	logger.l.Debug(msg, fields...)
}





