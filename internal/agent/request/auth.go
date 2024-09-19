package request

import (
	"github.com/pisarevaa/gophkeeper/internal/shared/model"
)

// Запрос на создание пользователя.
func (c *Client) RegisterUser(user model.RegisterUser) (int, error) {
	resp, err := c.Client.R().
		SetBody(user).
		SetHeader("Content-Type", "application/json").
		Post(c.ServerHost + "/auth/register")
	return resp.StatusCode(), err
}

// Запрос на авторизацию пользователя.
func (c *Client) LoginUser(user model.RegisterUser) (int, error) {
	resp, err := c.Client.R().
		SetBody(user).
		SetHeader("Content-Type", "application/json").
		Post(c.ServerHost + "/auth/register")
	return resp.StatusCode(), err
}
