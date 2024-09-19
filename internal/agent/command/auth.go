package command

import (
	"log/slog"

	"github.com/urfave/cli/v2"

	"github.com/pisarevaa/gophkeeper/internal/agent/config"
	"github.com/pisarevaa/gophkeeper/internal/shared/model"
	"github.com/pisarevaa/gophkeeper/internal/agent/request"
	"github.com/pisarevaa/gophkeeper/internal/agent/service"
	"github.com/pisarevaa/gophkeeper/internal/agent/utils"
)

func RegisterCommand() *cli.Command {
	command := cli.Command{
		Name:  "register",
		Usage: "register an account",
		Args:  true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "email",
				Required: true,
				Usage:    "user's email",
			},
			&cli.StringFlag{
				Name:     "password",
				Required: true,
				Usage:    "user's password with greater than 5 symbols",
			},
		},
		Action: func(cCtx *cli.Context) error {
			user := model.RegisterUser{
				Email:    cCtx.String("email"),
				Password: cCtx.String("password"),
			}
			config := config.NewConfig()
			service := service.NewService(
				service.WithClient(request.NewClient(config.ServerHost)),
				service.WithValidator(utils.NewValidator()),
				service.WithConfig(config),
			)
			err := service.RegisterUser(user)
			if err == nil {
				slog.Info("You are successfully registered in Gophkeeper")
			} else {
				slog.Error("Error: " + err.Error())
			}
			return nil
		},
	}
	return &command
}

func LoginCommand() *cli.Command {
	command := cli.Command{
		Name:  "login",
		Usage: "login into account",
		Args:  true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "email",
				Required: true,
				Usage:    "user's email",
			},
			&cli.StringFlag{
				Name:     "password",
				Required: true,
				Usage:    "user's password",
			},
		},
		Action: func(cCtx *cli.Context) error {
			user := model.RegisterUser{
				Email:    cCtx.String("email"),
				Password: cCtx.String("password"),
			}
			config := config.NewConfig()
			service := service.NewService(
				service.WithClient(request.NewClient(config.ServerHost)),
				service.WithValidator(utils.NewValidator()),
				service.WithConfig(config),
			)
			err := service.RegisterUser(user)
			if err == nil {
				slog.Info("You are successfully registered in Gophkeeper")
			} else {
				slog.Error("Error: " + err.Error())
			}
			return nil
		},
	}
	return &command
}