package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/pisarevaa/gophkeeper/internal/server/config"
	"github.com/pisarevaa/gophkeeper/internal/server/service/auth"
	"github.com/pisarevaa/gophkeeper/internal/server/service/keeper"
)

type Handler struct {
	Config        config.Config
	Validator     *validator.Validate
	AuthService   auth.Service
	KeeperService keeper.Service
}

type Option func(*Handler)

func WithConfig(config config.Config) Option {
	return func(s *Handler) {
		s.Config = config
	}
}

func WithValidator(validator *validator.Validate) Option {
	return func(s *Handler) {
		s.Validator = validator
	}
}

func WithAuthService(authService auth.Service) Option {
	return func(s *Handler) {
		s.AuthService = authService
	}
}

func WithKeeperService(keeperService keeper.Service) Option {
	return func(s *Handler) {
		s.KeeperService = keeperService
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
