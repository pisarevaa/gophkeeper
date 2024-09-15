package keeper

import (
	"context"

	"github.com/pisarevaa/gophkeeper/internal/server/model"
)

type KeeperServicer interface {
	GetData(ctx context.Context, userID int64) ([]model.Keeper, int, error)
	GetDataByID(ctx context.Context, userID int64, dataID int64) (model.Keeper, int, error)
	AddTextData(ctx context.Context, name string, textData string, userID int64) (model.Keeper, int, error)
	AddBinaryData(ctx context.Context, file model.UploadedFile, name string, userID int64) (model.Keeper, int, error)
	UpdateTextData(
		ctx context.Context,
		name string,
		textData string,
		userID int64,
		dataID int64,
	) (model.Keeper, int, error)
	UpdateBinaryData(
		ctx context.Context,
		file model.UploadedFile,
		name string,
		userID int64,
		dataID int64,
	) (model.Keeper, int, error)
	DeleteData(ctx context.Context, userID int64, dataID int64) (model.Keeper, int, error)
}
