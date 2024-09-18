package request

import (
	"github.com/pisarevaa/gophkeeper/internal/agent/model"
)

type Requester interface {
	RegisterUser(model.RegisterUser) (status int, err error)
}
