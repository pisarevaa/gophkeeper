package request

import (
	"github.com/pisarevaa/gophkeeper/internal/shared/model"
)

type Requester interface {
	SetToken() (err error)
	RegisterUser(user model.RegisterUser) (createdUser model.UserResponse, err error)
	LoginUser(user model.RegisterUser) (tokenResponse model.TokenResponse, err error)
	GetData() (dataResponse []model.DataResponseShort, err error)
	GetDataByID(dataID int64) (dataResponse model.DataResponse, err error)
	AddTextData(textData model.AddTextData) (dataResponse model.DataResponse, err error)
	UpdateTextData(textData model.AddTextData, dataID int64) (dataResponse model.DataResponse, err error)
	AddBinaryData(filepath string, name string) (dataResponse model.DataResponse, err error)
	UpdateBinaryData(filepath string, name string, dataID int64) (dataResponse model.DataResponse, err error)
	DeleteData(dataID int64) (dataResponse model.DataResponse, err error)
	DownloadFile(url string, filename string) (err error)
}
