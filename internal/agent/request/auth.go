package request

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/pisarevaa/gophkeeper/internal/agent/utils"
	"github.com/pisarevaa/gophkeeper/internal/shared/model"
)

// Запрос на создание пользователя.
func (c *Client) RegisterUser(user model.RegisterUser) (model.UserResponse, error) {
	var createdUser model.UserResponse
	resp, err := c.Client.R().
		SetBody(user).
		SetHeader("Content-Type", "application/json").
		Post(c.ServerHost + "/auth/register")

	if err != nil {
		return createdUser, err
	}
	if resp.StatusCode() != http.StatusOK {
		return createdUser, errors.New(string(resp.Body()))
	}
	if err = json.Unmarshal(resp.Body(), &createdUser); err != nil {
		return createdUser, err
	}
	return createdUser, nil
}

// Запрос на авторизацию пользователя.
func (c *Client) LoginUser(user model.RegisterUser) (model.TokenResponse, error) {
	var tokenResponse model.TokenResponse
	resp, err := c.Client.R().
		SetBody(user).
		SetHeader("Content-Type", "application/json").
		Post(c.ServerHost + "/auth/login")

	if err != nil {
		return tokenResponse, err
	}
	if resp.StatusCode() != http.StatusOK {
		return tokenResponse, errors.New(string(resp.Body()))
	}
	if err = json.Unmarshal(resp.Body(), &tokenResponse); err != nil {
		return tokenResponse, err
	}
	return tokenResponse, nil
}

func (c *Client) SetToken() error {
	tokenResponse, err := utils.LoadUserDataFromDosk()
	if err != nil {
		return err
	}
	if tokenResponse.Token == "" {
		return errors.New("token is not found in " + utils.UserFilename + ", login again please")
	}
	c.Client.Token = tokenResponse.Token
	return nil
}
