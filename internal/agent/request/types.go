package request

import (
	"bytes"

	"github.com/pisarevaa/gophkeeper/internal/shared/model"
)

type Requester interface {
	SetToken(token string)
	RegisterUser(user model.RegisterUser) (createdUser model.UserResponse, err error)
	LoginUser(user model.RegisterUser) (tokenResponse model.TokenResponse, err error)
	GetData() (dataResponse []model.DataResponse, err error)
	GetDataByID(dataID int64) (dataResponse model.DataResponse, err error)
	AddTextData(textData model.AddTextData) (dataResponse model.DataResponse, err error)
	UpdateTextData(textData model.AddTextData, dataID int64) (dataResponse model.DataResponse, err error)
	AddBinaryData(formData *bytes.Buffer) (dataResponse model.DataResponse, err error)
	UpdateBinaryData(formData *bytes.Buffer, dataID int64) (dataResponse model.DataResponse, err error)
	DeleteData(dataID int64) (dataResponse model.DataResponse, err error)
}
