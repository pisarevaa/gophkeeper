package handler

import (
	"go.uber.org/zap"

	"github.com/go-playground/validator/v10"
	"github.com/pisarevaa/gophkeeper/internal/server/config"
	"github.com/pisarevaa/gophkeeper/internal/server/storage"
)

type Handler struct {
	Config    config.Config
	Logger    *zap.SugaredLogger
	Storage   storage.Storage
	Validator *validator.Validate
}

type Option func(*Handler)

func WithConfig(config config.Config) Option {
	return func(s *Handler) {
		s.Config = config
	}
}

func WithLogger(logger *zap.SugaredLogger) Option {
	return func(s *Handler) {
		s.Logger = logger
	}
}

func WithStorage(storage storage.Storage) Option {
	return func(s *Handler) {
		s.Storage = storage
	}
}

func WithValidator(validator *validator.Validate) Option {
	return func(s *Handler) {
		s.Validator = validator
	}
}

// Создание хедлера роутера.
func NewHandler(options ...Option) *Handler {
	h := &Handler{}
	for _, option := range options {
		option(h)
	}
	return h
}
