package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/pisarevaa/gophkeeper/internal/agent/config"
	"github.com/pisarevaa/gophkeeper/internal/agent/request"
)

type Service struct {
	Client    request.Requester
	Validator *validator.Validate
	Config    *config.Config
}

type Option func(*Service)

func WithClient(client request.Requester) Option {
	return func(s *Service) {
		s.Client = client
	}
}

func WithValidator(validator *validator.Validate) Option {
	return func(s *Service) {
		s.Validator = validator
	}
}

func WithConfig(config *config.Config) Option {
	return func(s *Service) {
		s.Config = config
	}
}

// Создание сервиса.
func NewService(options ...Option) *Service {
	h := &Service{}
	for _, option := range options {
		option(h)
	}
	return h
}
