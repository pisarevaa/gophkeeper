package request

import (
	"github.com/pisarevaa/gophkeeper/internal/shared/model"
)

type Requester interface {
	RegisterUser(user model.RegisterUser) (createdUser model.UserResponse, status int, err error)
	LoginUser(user model.RegisterUser) (tokenResponse model.TokenResponse, status int, err error)
}
