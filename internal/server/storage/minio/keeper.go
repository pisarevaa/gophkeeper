package minio

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/pisarevaa/gophkeeper/internal/server/model"
)

// Загрузка объекта Minio.
func (m *Minio) CreateOne(ctx context.Context, bucket string, file model.UploadedFile) (string, error) {
	objectID := uuid.New().String()
	reader := bytes.NewReader(file.Data)
	_, err := m.PutObject(
		context.Background(),
		bucket,
		objectID,
		reader,
		int64(len(file.Data)),
		minio.PutObjectOptions{},
	)
	if err != nil {
		return "", fmt.Errorf("ошибка при создании объекта %s: %v", file.FileName, err)
	}
	return objectID, nil
}

// Получение ссылки на объект Minio.
func (m *Minio) GetOne(ctx context.Context, bucket string, objectID string) (string, error) {
	linkExpiedAt := time.Second * 24 * 60 * 60
	url, err := m.PresignedGetObject(ctx, bucket, objectID, linkExpiedAt, nil)
	if err != nil {
		return "", fmt.Errorf("ошибка при получении URL для объекта %s: %v", objectID, err)
	}
	return url.String(), nil
}

// Удаление объекта из Minio.
func (m *Minio) DeleteOne(ctx context.Context, bucket string, objectID string) error {
	err := m.RemoveObject(ctx, bucket, objectID, minio.RemoveObjectOptions{})
	if err != nil {
		return err
	}
	return nil
}
