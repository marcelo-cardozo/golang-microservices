package log

import (
	zapT "github.com/marcelo-cardozo/golang-microservices/src/api/log/zap"
	zap "go.uber.org/zap"
)

func Info(msg string, fields ...zap.Field) {
	zapT.Info(msg, fields...)
}