package command

import (
	"github.com/pisarevaa/gophkeeper/internal/agent/config"
	"github.com/pisarevaa/gophkeeper/internal/agent/request"
	"github.com/pisarevaa/gophkeeper/internal/agent/service"
	sharedUtils "github.com/pisarevaa/gophkeeper/internal/shared/utils"
)

func NewCommand() (*service.Service, error) {
	config, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	service := service.NewService(
		service.WithClient(request.NewClient(config.ServerHost)),
		service.WithValidator(sharedUtils.NewValidator()),
		service.WithConfig(config),
	)
	return service, nil
}
