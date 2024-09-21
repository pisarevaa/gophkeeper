package request

import (
	"bytes"
	"errors"
	"net/http"
	"strconv"

	"github.com/pisarevaa/gophkeeper/internal/shared/model"
)

// Получение всех данных пользователя.
func (c *Client) GetData() ([]model.DataResponse, error) {
	var dataResponse []model.DataResponse
	var errMsg string
	resp, err := c.Client.R().
		SetResult(&dataResponse).
		SetError(&errMsg).
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+c.Token).
		Get(c.ServerHost + "/api/data")

	if err != nil {
		return dataResponse, err
	}
	if resp.StatusCode() != http.StatusOK {
		return dataResponse, errors.New(errMsg)
	}
	return dataResponse, nil
}

// Получение данных пользователя по ID.
func (c *Client) GetDataByID(dataID int64) (model.DataResponse, error) {
	var dataResponse model.DataResponse
	var errMsg string
	resp, err := c.Client.R().
		SetResult(&dataResponse).
		SetError(&errMsg).
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+c.Token).
		Get(c.ServerHost + "/api/data/" + strconv.FormatInt(dataID, 10))

	if err != nil {
		return dataResponse, err
	}
	if resp.StatusCode() != http.StatusOK {
		return dataResponse, errors.New(errMsg)
	}
	return dataResponse, nil
}

// Добавление текстовых данных пользователем.
func (c *Client) AddTextData(textData model.AddTextData) (model.DataResponse, error) {
	var dataResponse model.DataResponse
	var errMsg string
	resp, err := c.Client.R().
		SetBody(textData).
		SetResult(&dataResponse).
		SetError(&errMsg).
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+c.Token).
		Post(c.ServerHost + "/api/data/text")

	if err != nil {
		return dataResponse, err
	}
	if resp.StatusCode() != http.StatusOK {
		return dataResponse, errors.New(errMsg)
	}
	return dataResponse, nil
}

// Обновление текстовых данных пользователем.
func (c *Client) UpdateTextData(textData model.AddTextData, dataID int64) (model.DataResponse, error) {
	var dataResponse model.DataResponse
	var errMsg string
	resp, err := c.Client.R().
		SetBody(textData).
		SetResult(&dataResponse).
		SetError(&errMsg).
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+c.Token).
		Put(c.ServerHost + "/api/data/text/" + strconv.FormatInt(dataID, 10))
	if err != nil {
		return dataResponse, err
	}
	if resp.StatusCode() != http.StatusOK {
		return dataResponse, errors.New(errMsg)
	}
	return dataResponse, nil
}

// Добавление бинарных данных пользователем.
func (c *Client) AddBinaryData(formData *bytes.Buffer) (model.DataResponse, error) {
	var dataResponse model.DataResponse
	var errMsg string
	resp, err := c.Client.R().
		SetBody(formData).
		SetResult(&dataResponse).
		SetError(&errMsg).
		SetHeader("Content-Type", "multipart/form-data").
		SetHeader("Authorization", "Bearer "+c.Token).
		Post(c.ServerHost + "/api/data/binary")

	if err != nil {
		return dataResponse, err
	}
	if resp.StatusCode() != http.StatusOK {
		return dataResponse, errors.New(errMsg)
	}
	return dataResponse, nil
}

// Обновление бинарных данных пользователем.
func (c *Client) UpdateBinaryData(formData *bytes.Buffer, dataID int64) (model.DataResponse, error) {
	var dataResponse model.DataResponse
	var errMsg string
	resp, err := c.Client.R().
		SetBody(formData).
		SetResult(&dataResponse).
		SetError(&errMsg).
		SetHeader("Content-Type", "multipart/form-data").
		SetHeader("Authorization", "Bearer "+c.Token).
		Put(c.ServerHost + "/api/data/binary/" + strconv.FormatInt(dataID, 10))

	if err != nil {
		return dataResponse, err
	}
	if resp.StatusCode() != http.StatusOK {
		return dataResponse, errors.New(errMsg)
	}
	return dataResponse, nil
}

// Удаление данных пользователя по ID.
func (c *Client) DeleteData(dataID int64) (model.DataResponse, error) {
	var dataResponse model.DataResponse
	var errMsg string
	resp, err := c.Client.R().
		SetResult(&dataResponse).
		SetError(&errMsg).
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+c.Token).
		Delete(c.ServerHost + "/api/data/" + strconv.FormatInt(dataID, 10))

	if err != nil {
		return dataResponse, err
	}
	if resp.StatusCode() != http.StatusOK {
		return dataResponse, errors.New(errMsg)
	}
	return dataResponse, nil
}
