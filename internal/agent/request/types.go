package request

import (
	"bytes"
	"github.com/pisarevaa/gophkeeper/internal/shared/model"
)

type Requester interface {
	RegisterUser(user model.RegisterUser) (createdUser model.UserResponse, status int, err error)
	LoginUser(user model.RegisterUser) (tokenResponse model.TokenResponse, status int, err error)
	GetData() (dataResponse []model.DataResponse, status int, err error)
	GetDataByID(dataID int64) (dataResponse model.DataResponse, status int, err error)
	AddTextData(textData model.AddTextData) (dataResponse model.DataResponse, status int, err error)
	UpdateTextData(textData model.AddTextData, dataID int64) (dataResponse model.DataResponse, status int, err error)
	AddBinaryData(formData *bytes.Buffer) (dataResponse model.DataResponse, status int, err error)
	UpdateBinaryData(formData *bytes.Buffer, dataID int64) (dataResponse model.DataResponse, status int, err error)
	DeleteData(dataID int64) (dataResponse model.DataResponse, status int, err error)
}
