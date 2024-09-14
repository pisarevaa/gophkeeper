package auth

import (
	"github.com/pisarevaa/gophkeeper/internal/server/config"
	"github.com/pisarevaa/gophkeeper/internal/server/storage"
)

type AuthService struct { //nolint:revive // it's okey
	Config  config.Config
	Storage storage.AuthStorage
}

type Option func(*AuthService)

func WithConfig(config config.Config) Option {
	return func(s *AuthService) {
		s.Config = config
	}
}

func WithStorage(storage storage.AuthStorage) Option {
	return func(s *AuthService) {
		s.Storage = storage
	}
}

// Создание сервиса.
func NewService(options ...Option) *AuthService {
	h := &AuthService{}
	for _, option := range options {
		option(h)
	}
	return h
}
