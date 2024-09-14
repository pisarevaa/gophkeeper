package keeper

import (
	"context"
	"net/http"

	"github.com/pisarevaa/gophkeeper/internal/server/model"
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
	return data, 0, nil
}

// Добавление новых данных.
func (s *KeeperService) AddData(ctx context.Context, keeper model.AddKeeper, userID int64) (model.Keeper, int, error) {
	data, err := s.Storage.AddData(ctx, keeper, userID)
	if err != nil {
		return data, http.StatusInternalServerError, err
	}
	return data, 0, nil
}

// Обновление данных по ID.
func (s *KeeperService) UpdateData(
	ctx context.Context,
	keeper model.AddKeeper,
	userID int64,
	dataID int64,
) (model.Keeper, int, error) {
	foundData, status, err := s.GetDataByID(ctx, userID, dataID)
	if err != nil {
		return foundData, status, err
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
	data, err := s.Storage.DeleteData(ctx, dataID)
	if err != nil {
		return data, http.StatusInternalServerError, err
	}
	return data, 0, nil
}
