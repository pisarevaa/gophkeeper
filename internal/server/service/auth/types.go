package auth

import (
	"context"

	"github.com/pisarevaa/gophkeeper/internal/shared/model"
)

type AuthServicer interface { //nolint:revive // it's okey
	RegisterUser(ctx context.Context, user model.RegisterUser) (model.User, int, error)
	Login(ctx context.Context, user model.RegisterUser) (string, int, error)
}
