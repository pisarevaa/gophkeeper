package storage

import (
	"context"

	"github.com/pisarevaa/gophkeeper/internal/server/model"
)

type Storage interface {
	GetUserByEmail(ctx context.Context, email string) (user model.User, err error)
	GetUserByID(ctx context.Context, userID int64) (user model.User, err error)
	RegisterUser(ctx context.Context, email string, passwordHash string) (user model.User, err error)
	CloseConnection()
}
