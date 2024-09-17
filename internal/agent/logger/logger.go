package logger

import (
	"log/slog"
	"os"
)

// Инициализация логера.
func NewLogger() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
}
