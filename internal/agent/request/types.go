package request

import (
	"github.com/pisarevaa/gophkeeper/internal/shared/model"
)

type Requester interface {
	RegisterUser(model.RegisterUser) (status int, err error)
	LoginUser(model.RegisterUser) (status int, err error)
}
