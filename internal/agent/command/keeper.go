package command

import (
	"log/slog"

	"github.com/urfave/cli/v2"

	"github.com/pisarevaa/gophkeeper/internal/agent/config"
	"github.com/pisarevaa/gophkeeper/internal/agent/request"
	"github.com/pisarevaa/gophkeeper/internal/agent/service"
	sharedUtils "github.com/pisarevaa/gophkeeper/internal/shared/utils"
)

func GetDataCommand() *cli.Command { //nolint:dupl // it's ok
	command := cli.Command{
		Name:  "get_all_data",
		Usage: "get all user's data",
		Args:  false,
		Action: func(_ *cli.Context) error {
			config := config.NewConfig()
			service := service.NewService(
				service.WithClient(request.NewClient(config.ServerHost)),
				service.WithValidator(sharedUtils.NewValidator()),
				service.WithConfig(config),
			)
			err := service.GetData()
			if err != nil {
				slog.Error("Error: " + err.Error())
			}
			return nil
		},
	}
	return &command
}

func GetDataByIDCommand() *cli.Command { //nolint:dupl // it's ok
	command := cli.Command{
		Name:  "get_data_by_id",
		Usage: "get all user's data",
		Args:  true,
		Flags: []cli.Flag{
			&cli.Int64Flag{
				Name:     "dataID",
				Required: true,
				Usage:    "data ID",
			},
		},
		Action: func(cCtx *cli.Context) error {
			dataID := cCtx.Int64("dataID")
			config := config.NewConfig()
			service := service.NewService(
				service.WithClient(request.NewClient(config.ServerHost)),
				service.WithValidator(sharedUtils.NewValidator()),
				service.WithConfig(config),
			)
			err := service.GetDataByID(dataID)
			if err != nil {
				slog.Error("Error: " + err.Error())
			}
			return nil
		},
	}
	return &command
}
