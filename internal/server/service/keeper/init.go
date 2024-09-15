package keeper

import (
	"github.com/pisarevaa/gophkeeper/internal/server/config"
	"github.com/pisarevaa/gophkeeper/internal/server/storage/db"
	"github.com/pisarevaa/gophkeeper/internal/server/storage/minio"
)

type KeeperService struct { //nolint:revive // it's okey
	Config  config.Config
	Storage db.KeeperStorage
	Minio   minio.MinioStorage
}

type Option func(*KeeperService)

func WithConfig(config config.Config) Option {
	return func(s *KeeperService) {
		s.Config = config
	}
}

func WithStorage(storage db.KeeperStorage) Option {
	return func(s *KeeperService) {
		s.Storage = storage
	}
}

func WithMinio(minio minio.MinioStorage) Option {
	return func(s *KeeperService) {
		s.Minio = minio
	}
}

// Создание сервиса.
func NewService(options ...Option) *KeeperService {
	h := &KeeperService{}
	for _, option := range options {
		option(h)
	}
	return h
}
