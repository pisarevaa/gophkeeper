package handler

import (
	"go.uber.org/zap"

	"github.com/go-playground/validator/v10"
	"github.com/pisarevaa/fastlog/internal/config"
	"github.com/pisarevaa/fastlog/internal/producer"
	"github.com/pisarevaa/fastlog/internal/storage"
)

type Handler struct {
	Config    config.Config
	Logger    *zap.SugaredLogger
	Storage   storage.Storage
	Validator *validator.Validate
	Producer  producer.Producer
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

func WithProducer(producer producer.Producer) Option {
	return func(s *Handler) {
		s.Producer = producer
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
