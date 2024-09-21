package request

import (
	"time"

	"github.com/go-resty/resty/v2"
)

const timeoutSec = 30

type Client struct {
	Client     *resty.Client
	ServerHost string
	Token      string
}

// Создание клиента для запросов.
func NewClient(serverHost string) *Client {
	restyClient := resty.New()
	restyClient.SetTimeout(timeoutSec * time.Second)
	return &Client{Client: restyClient, ServerHost: serverHost}
}

func (c *Client) SetToken(token string) {
	c.Client.Token = token
}
