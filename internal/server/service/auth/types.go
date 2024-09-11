package auth

import (
	"context"

	"github.com/pisarevaa/gophkeeper/internal/server/model"
)

type Service interface {
	RegisterUser(ctx context.Context, user model.RegisterUser) (model.User, int, error)
	Login(ctx context.Context, user model.RegisterUser) (string, int, error)
}
