package command

import (
	"log/slog"

	"github.com/urfave/cli/v2"

	"github.com/pisarevaa/gophkeeper/internal/shared/model"
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
			service, err := NewCommand()
			if err != nil {
				return err
			}
			err = service.RegisterUser(user)
			if err != nil {
				return err
			}
			slog.Info("You are successfully registered into Gophkeeper")
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
			service, err := NewCommand()
			if err != nil {
				return err
			}
			err = service.LoginUser(user)
			if err != nil {
				return err
			}
			slog.Info("You are successfully login into Gophkeeper")
			return nil
		},
	}
	return &command
}
