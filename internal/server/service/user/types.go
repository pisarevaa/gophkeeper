package user

import (
	"context"

	"github.com/pisarevaa/gophkeeper/internal/server/model"
)

type Service interface {
	RegisterUser(ctx context.Context, user model.RegisterUser) (model.User, error)
}
