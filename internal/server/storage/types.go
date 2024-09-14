package storage

import (
	"context"

	"github.com/pisarevaa/gophkeeper/internal/server/model"
)

type AuthStorage interface {
	GetUserByEmail(ctx context.Context, email string) (user model.User, err error)
	GetUserByID(ctx context.Context, userID int64) (user model.User, err error)
	RegisterUser(ctx context.Context, email string, passwordHash string) (user model.User, err error)
}

type KeeperStorage interface {
	GetDataByUserID(ctx context.Context, userID int64) ([]model.Keeper, error)
	GetDataByID(ctx context.Context, dataID int64) (model.Keeper, error)
	AddData(ctx context.Context, keeper model.AddKeeper, userID int64) (model.Keeper, error)
	UpdateData(ctx context.Context, keeper model.AddKeeper, dataID int64) (model.Keeper, error)
	DeleteData(ctx context.Context, dataID int64) (model.Keeper, error)
}
