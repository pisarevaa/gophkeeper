package user

import (
	"github.com/pisarevaa/gophkeeper/internal/server/config"
	"github.com/pisarevaa/gophkeeper/internal/server/storage"
)

type UserService struct { //nolint:revive // it's ok
	Config  config.Config
	Storage storage.Storage
}

type Option func(*UserService)

func WithConfig(config config.Config) Option {
	return func(s *UserService) {
		s.Config = config
	}
}

func WithStorage(storage storage.Storage) Option {
	return func(s *UserService) {
		s.Storage = storage
	}
}

// Создание сервиса.
func NewService(options ...Option) *UserService {
	h := &UserService{}
	for _, option := range options {
		option(h)
	}
	return h
}
