package request

import (
	"bytes"
	"strconv"

	"github.com/pisarevaa/gophkeeper/internal/shared/model"
)

// SetError().  !!!!!!!!!!!!!!!!!

// Получение всех данных пользователя.
func (c *Client) GetData() ([]model.DataResponse, int, error) {
	var dataResponse []model.DataResponse
	resp, err := c.Client.R().
		SetResult(&dataResponse).
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+c.Token).
		Get(c.ServerHost + "/api/data")
	return dataResponse, resp.StatusCode(), err
}

// Получение данных пользователя по ID.
func (c *Client) GetDataByID(dataID int64) (model.DataResponse, int, error) {
	var dataResponse model.DataResponse
	resp, err := c.Client.R().
		SetResult(&dataResponse).
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+c.Token).
		Get(c.ServerHost + "/api/data/" + strconv.FormatInt(dataID, 10))
	return dataResponse, resp.StatusCode(), err
}

// Добавление текстовых данных пользователем.
func (c *Client) AddTextData(textData model.AddTextData) (model.DataResponse, int, error) {
	var dataResponse model.DataResponse
	resp, err := c.Client.R().
		SetBody(textData).
		SetResult(&dataResponse).
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+c.Token).
		Post(c.ServerHost + "/api/data/text")
	return dataResponse, resp.StatusCode(), err
}

// Обновление текстовых данных пользователем.
func (c *Client) UpdateTextData(textData model.AddTextData, dataID int64) (model.DataResponse, int, error) {
	var dataResponse model.DataResponse
	resp, err := c.Client.R().
		SetBody(textData).
		SetResult(&dataResponse).
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+c.Token).
		Put(c.ServerHost + "/api/data/text/" + strconv.FormatInt(dataID, 10))
	return dataResponse, resp.StatusCode(), err
}

// Добавление бинарных данных пользователем.
func (c *Client) AddBinaryData(formData *bytes.Buffer) (model.DataResponse, int, error) {
	var dataResponse model.DataResponse
	resp, err := c.Client.R().
		SetBody(formData).
		SetResult(&dataResponse).
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+c.Token).
		Post(c.ServerHost + "/api/data/binary")
	return dataResponse, resp.StatusCode(), err
}

// Обновление бинарных данных пользователем.
func (c *Client) UpdateBinaryData(formData *bytes.Buffer, dataID int64) (model.DataResponse, int, error) {
	var dataResponse model.DataResponse
	resp, err := c.Client.R().
		SetBody(formData).
		SetResult(&dataResponse).
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+c.Token).
		Put(c.ServerHost + "/api/data/binary/" + strconv.FormatInt(dataID, 10))
	return dataResponse, resp.StatusCode(), err
}

// Удаление данных пользователя по ID.
func (c *Client) DeleteData(dataID int64) (model.DataResponse, int, error) {
	var dataResponse model.DataResponse
	resp, err := c.Client.R().
		SetResult(&dataResponse).
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+c.Token).
		Delete(c.ServerHost + "/api/data/" + strconv.FormatInt(dataID, 10))
	return dataResponse, resp.StatusCode(), err
}
