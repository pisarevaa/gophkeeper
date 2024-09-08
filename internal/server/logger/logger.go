package logger

import (
	"go.uber.org/zap"
)

// Инициализация логера.
func NewLogger() *zap.SugaredLogger {
	logger := zap.NewExample().Sugar()
	defer logger.Sync() //nolint:errcheck // ignore check
	return logger
}
