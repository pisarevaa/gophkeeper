package minio

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/pisarevaa/gophkeeper/internal/server/config"
)

type Minio struct {
	*minio.Client
}

// Создание подключения к Minio.
func NewMinio(cfg config.Minio) (*Minio, error) {
	client, err := minio.New(cfg.Host, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.User, cfg.Password, ""),
		Secure: cfg.UseSSL,
	})
	if err != nil {
		return nil, err
	}
	slog.Info("Connected to Minio")
	coon := &Minio{client}
	ctx := context.Background()
	// Проверка наличия бакета и его создание, если не существует
	exists, err := coon.BucketExists(ctx, cfg.Bucket)
	if err != nil {
		return nil, err
	}
	if !exists {
		err := coon.MakeBucket(ctx, cfg.Bucket, minio.MakeBucketOptions{})
		if err != nil {
			return nil, err
		}
		slog.Info(fmt.Sprint("Created bucket", cfg.Bucket))
	}
	return coon, nil
}
