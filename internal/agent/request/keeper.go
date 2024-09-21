package request

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/pisarevaa/gophkeeper/internal/shared/model"
)

// Получение всех данных пользователя.
func (c *Client) GetData() ([]model.DataResponseShort, error) {
	var dataResponse []model.DataResponseShort
	resp, err := c.Client.R().
		SetResult(&dataResponse).
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+c.Token).
		Get(c.ServerHost + "/api/data")
	if err != nil {
		return dataResponse, err
	}
	if resp.StatusCode() != http.StatusOK {
		return dataResponse, errors.New(string(resp.Body()))
	}
	if err = json.Unmarshal(resp.Body(), &dataResponse); err != nil {
		return dataResponse, err
	}
	return dataResponse, nil
}

// Получение данных пользователя по ID.
func (c *Client) GetDataByID(dataID int64) (model.DataResponse, error) {
	var dataResponse model.DataResponse
	resp, err := c.Client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+c.Token).
		Get(c.ServerHost + "/api/data/" + strconv.FormatInt(dataID, 10))

	if err != nil {
		return dataResponse, err
	}
	if resp.StatusCode() != http.StatusOK {
		return dataResponse, errors.New(string(resp.Body()))
	}
	if err = json.Unmarshal(resp.Body(), &dataResponse); err != nil {
		return dataResponse, err
	}
	return dataResponse, nil
}

// Добавление текстовых данных пользователем.
func (c *Client) AddTextData(textData model.AddTextData) (model.DataResponse, error) {
	var dataResponse model.DataResponse
	resp, err := c.Client.R().
		SetBody(textData).
		SetResult(&dataResponse).
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+c.Token).
		Post(c.ServerHost + "/api/data/text")

	if err != nil {
		return dataResponse, err
	}
	if resp.StatusCode() != http.StatusOK {
		return dataResponse, errors.New(string(resp.Body()))
	}
	if err = json.Unmarshal(resp.Body(), &dataResponse); err != nil {
		return dataResponse, err
	}
	return dataResponse, nil
}

// Обновление текстовых данных пользователем.
func (c *Client) UpdateTextData(textData model.AddTextData, dataID int64) (model.DataResponse, error) {
	var dataResponse model.DataResponse
	resp, err := c.Client.R().
		SetBody(textData).
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+c.Token).
		Put(c.ServerHost + "/api/data/text/" + strconv.FormatInt(dataID, 10))
	if err != nil {
		return dataResponse, err
	}
	if resp.StatusCode() != http.StatusOK {
		return dataResponse, errors.New(string(resp.Body()))
	}
	if err = json.Unmarshal(resp.Body(), &dataResponse); err != nil {
		return dataResponse, err
	}
	return dataResponse, nil
}

// Добавление бинарных данных пользователем.
func (c *Client) AddBinaryData(filepath string, name string) (model.DataResponse, error) {
	var dataResponse model.DataResponse
	resp, err := c.Client.R().
		SetFile("file", filepath).
		SetFormData(map[string]string{
			"name": name,
		}).
		SetHeader("Authorization", "Bearer "+c.Token).
		Post(c.ServerHost + "/api/data/binary")

	if err != nil {
		return dataResponse, err
	}
	if resp.StatusCode() != http.StatusOK {
		return dataResponse, errors.New(string(resp.Body()))
	}
	if err = json.Unmarshal(resp.Body(), &dataResponse); err != nil {
		return dataResponse, err
	}
	return dataResponse, nil
}

// Обновление бинарных данных пользователем.
func (c *Client) UpdateBinaryData(filepath string, name string, dataID int64) (model.DataResponse, error) {
	var dataResponse model.DataResponse
	resp, err := c.Client.R().
		SetFile("file", filepath).
		SetFormData(map[string]string{
			"name": name,
		}).
		SetHeader("Authorization", "Bearer "+c.Token).
		Put(c.ServerHost + "/api/data/binary/" + strconv.FormatInt(dataID, 10))

	if err != nil {
		return dataResponse, err
	}
	if resp.StatusCode() != http.StatusOK {
		return dataResponse, errors.New(string(resp.Body()))
	}
	if err = json.Unmarshal(resp.Body(), &dataResponse); err != nil {
		return dataResponse, err
	}
	return dataResponse, nil
}

// Удаление данных пользователя по ID.
func (c *Client) DeleteData(dataID int64) (model.DataResponse, error) {
	var dataResponse model.DataResponse
	resp, err := c.Client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+c.Token).
		Delete(c.ServerHost + "/api/data/" + strconv.FormatInt(dataID, 10))

	if err != nil {
		return dataResponse, err
	}
	if resp.StatusCode() != http.StatusOK {
		return dataResponse, errors.New(string(resp.Body()))
	}
	if err = json.Unmarshal(resp.Body(), &dataResponse); err != nil {
		return dataResponse, err
	}
	return dataResponse, nil
}
