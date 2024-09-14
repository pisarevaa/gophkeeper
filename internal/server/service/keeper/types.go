package keeper

import (
	"context"

	"github.com/pisarevaa/gophkeeper/internal/server/model"
)

type Service interface {
	GetData(ctx context.Context, userID int64) ([]model.Keeper, int, error)
	GetDataByID(ctx context.Context, userID int64, dataID int64) (model.Keeper, int, error)
	AddData(ctx context.Context, keeper model.AddKeeper, userID int64) (model.Keeper, int, error)
	UpdateData(ctx context.Context, keeper model.AddKeeper, userID int64, dataID int64) (model.Keeper, int, error)
	DeleteData(ctx context.Context, userID int64, dataID int64) (model.Keeper, int, error)
}
