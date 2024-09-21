package request

import (
	"errors"
	"net/http"

	"github.com/pisarevaa/gophkeeper/internal/shared/model"
)

// Запрос на создание пользователя.
func (c *Client) RegisterUser(user model.RegisterUser) (model.UserResponse, error) {
	var createdUser model.UserResponse
	var errMsg string
	resp, err := c.Client.R().
		SetResult(&createdUser).
		SetBody(user).
		SetError(&errMsg).
		SetHeader("Content-Type", "application/json").
		Post(c.ServerHost + "/auth/register")

	if err != nil {
		return createdUser, err
	}
	if resp.StatusCode() != http.StatusOK {
		return createdUser, errors.New(errMsg)
	}
	return createdUser, nil
}

// Запрос на авторизацию пользователя.
func (c *Client) LoginUser(user model.RegisterUser) (model.TokenResponse, error) {
	var tokenResponse model.TokenResponse
	var errMsg string
	resp, err := c.Client.R().
		SetResult(&tokenResponse).
		SetBody(user).
		SetError(&errMsg).
		SetHeader("Content-Type", "application/json").
		Post(c.ServerHost + "/auth/register")

	if err != nil {
		return tokenResponse, err
	}
	if resp.StatusCode() != http.StatusOK {
		return tokenResponse, errors.New(errMsg)
	}
	return tokenResponse, nil
}
