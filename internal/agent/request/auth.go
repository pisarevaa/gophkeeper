package request

import (
	"github.com/pisarevaa/gophkeeper/internal/shared/model"
)

// Запрос на создание пользователя.
func (c *Client) RegisterUser(user model.RegisterUser) (model.UserResponse, int, error) {
	var createdUser model.UserResponse
	resp, err := c.Client.R().
		SetResult(&createdUser).
		SetBody(user).
		SetHeader("Content-Type", "application/json").
		Post(c.ServerHost + "/auth/register")
	return createdUser, resp.StatusCode(), err
}

// Запрос на авторизацию пользователя.
func (c *Client) LoginUser(user model.RegisterUser) (model.TokenResponse, int, error) {
	var tokenResponse model.TokenResponse
	resp, err := c.Client.R().
		SetResult(&tokenResponse).
		SetBody(user).
		SetHeader("Content-Type", "application/json").
		Post(c.ServerHost + "/auth/register")
	return tokenResponse, resp.StatusCode(), err
}
