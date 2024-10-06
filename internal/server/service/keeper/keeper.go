package keeper

import (
	"context"
	"net/http"

	"github.com/pisarevaa/gophkeeper/internal/shared/model"
)

// Получение всех данных пользователя по ID.
func (s *KeeperService) GetData(ctx context.Context, userID int64) ([]model.Keeper, int, error) {
	data, err := s.Storage.GetDataByUserID(ctx, userID)
	if err != nil {
		return data, http.StatusInternalServerError, err
	}
	return data, 0, nil
}

// Получение данных по ID.
func (s *KeeperService) GetDataByID(ctx context.Context, userID int64, dataID int64) (model.Keeper, int, error) {
	data, err := s.Storage.GetDataByID(ctx, dataID)
	if err != nil {
		return data, http.StatusNotFound, err
	}
	if data.UserID != userID {
		return data, http.StatusUnauthorized, err
	}
	if data.Type == model.BinaryType {
		url, errMinio := s.Minio.GetOne(ctx, s.Config.Minio.Bucket, data.ObjectID)
		if errMinio != nil {
			return data, http.StatusNotFound, errMinio
		}
		data.Data = url
	}
	return data, 0, nil
}

// Добавление текстовых данных.
func (s *KeeperService) AddTextData(
	ctx context.Context,
	name string,
	textData string,
	userID int64,
) (model.Keeper, int, error) {
	keeper := model.AddKeeper{
		Name: name,
		Data: textData,
		Type: model.TextType,
	}
	data, err := s.Storage.AddData(ctx, keeper, userID)
	if err != nil {
		return data, http.StatusInternalServerError, err
	}
	return data, 0, nil
}

// Добавление бинарных данных.
func (s *KeeperService) AddBinaryData(
	ctx context.Context,
	file model.UploadedFile,
	name string,
	userID int64,
) (model.Keeper, int, error) {
	objectID, err := s.Minio.CreateOne(ctx, s.Config.Minio.Bucket, file)
	if err != nil {
		return model.Keeper{}, http.StatusInternalServerError, err
	}
	keeper := model.AddKeeper{
		Name:     name,
		ObjectID: objectID,
		FileName: file.FileName,
		Type:     model.BinaryType,
	}
	data, err := s.Storage.AddData(ctx, keeper, userID)
	if err != nil {
		return data, http.StatusInternalServerError, err
	}
	return data, 0, nil
}

// Обновление текстовых данных по ID.
func (s *KeeperService) UpdateTextData(
	ctx context.Context,
	name string,
	textData string,
	userID int64,
	dataID int64,
) (model.Keeper, int, error) {
	foundData, status, err := s.GetDataByID(ctx, userID, dataID)
	if err != nil {
		return foundData, status, err
	}
	keeper := model.AddKeeper{
		Name: name,
		Data: textData,
		Type: model.TextType,
	}
	data, err := s.Storage.UpdateData(ctx, keeper, dataID)
	if err != nil {
		return data, http.StatusInternalServerError, err
	}
	return data, 0, nil
}

// Обновление бинарных данных по ID.
func (s *KeeperService) UpdateBinaryData(
	ctx context.Context,
	name string,
	file model.UploadedFile,
	userID int64,
	dataID int64,
) (model.Keeper, int, error) {
	foundData, status, err := s.GetDataByID(ctx, userID, dataID)
	if err != nil {
		return foundData, status, err
	}
	err = s.Minio.DeleteOne(ctx, s.Config.Minio.Bucket, foundData.ObjectID)
	if err != nil {
		return model.Keeper{}, http.StatusInternalServerError, err
	}
	objectID, err := s.Minio.CreateOne(ctx, s.Config.Minio.Bucket, file)
	if err != nil {
		return model.Keeper{}, http.StatusInternalServerError, err
	}
	keeper := model.AddKeeper{
		Name:     name,
		ObjectID: objectID,
		FileName: file.FileName,
		Type:     model.BinaryType,
	}
	data, err := s.Storage.UpdateData(ctx, keeper, dataID)
	if err != nil {
		return data, http.StatusInternalServerError, err
	}
	return data, 0, nil
}

// Удаление данных по ID.
func (s *KeeperService) DeleteData(ctx context.Context, userID int64, dataID int64) (model.Keeper, int, error) {
	foundData, status, err := s.GetDataByID(ctx, userID, dataID)
	if err != nil {
		return foundData, status, err
	}
	if foundData.Type == model.BinaryType {
		err = s.Minio.DeleteOne(ctx, s.Config.Minio.Bucket, foundData.ObjectID)
		if err != nil {
			return model.Keeper{}, http.StatusInternalServerError, err
		}
	}
	data, err := s.Storage.DeleteData(ctx, dataID)
	if err != nil {
		return data, http.StatusInternalServerError, err
	}
	return data, 0, nil
}
