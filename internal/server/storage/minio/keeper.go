package minio

import (
	"bytes"
	"context"
	"fmt"
	"sync"
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

// Получение множества ссылок на объекты Minio.
func (m *Minio) GetMany(ctx context.Context, bucket string, objectIDs []string) ([]string, error) {
	urlCh := make(chan string, len(objectIDs))
	errCh := make(chan model.MinioOperationError, len(objectIDs))

	var wg sync.WaitGroup
	_, cancel := context.WithCancel(ctx)
	defer cancel()

	for _, objectID := range objectIDs {
		wg.Add(1)
		go func(objectID string) {
			defer wg.Done()
			url, err := m.GetOne(ctx, bucket, objectID)
			if err != nil {
				errCh <- model.MinioOperationError{ObjectID: objectID, Error: fmt.Errorf("ошибка при получении объекта %s: %v", objectID, err)}
				cancel()
				return
			}
			urlCh <- url
		}(objectID)
	}

	go func() {
		wg.Wait()
		close(urlCh)
		close(errCh)
	}()

	var urls []string
	var errs []error
	for url := range urlCh {
		urls = append(urls, url)
	}
	for opErr := range errCh {
		errs = append(errs, opErr.Error)
	}

	if len(errs) > 0 {
		return nil, fmt.Errorf("ошибки при получении объектов: %v", errs)
	}

	return urls, nil
}

// Удаление объекта из Minio.
func (m *Minio) DeleteOne(ctx context.Context, bucket string, objectID string) error {
	err := m.RemoveObject(ctx, bucket, objectID, minio.RemoveObjectOptions{})
	if err != nil {
		return err
	}
	return nil
}
