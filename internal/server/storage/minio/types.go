package minio

import (
	"context"

	"github.com/pisarevaa/gophkeeper/internal/shared/model"
)

type MinioStorage interface { //nolint:revive // it's okey
	CreateOne(ctx context.Context, bucket string, file model.UploadedFile) (string, error)
	GetOne(ctx context.Context, bucket string, objectID string) (string, error)
	GetMany(ctx context.Context, bucket string, objectIDs []string) ([]string, error)
	DeleteOne(ctx context.Context, bucket string, objectID string) error
}
