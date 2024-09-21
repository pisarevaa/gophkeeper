package command

import (
	"github.com/urfave/cli/v2"

	"github.com/pisarevaa/gophkeeper/internal/shared/model"
)

// Получение всех данных пользователя.
func GetDataCommand() *cli.Command {
	command := cli.Command{
		Name:  "get_all_data",
		Usage: "get all user's data",
		Args:  false,
		Action: func(_ *cli.Context) error {
			service, err := NewCommand()
			if err != nil {
				return err
			}
			err = service.GetData()
			if err != nil {
				return err
			}
			return nil
		},
	}
	return &command
}

// Получение данных пользователя по ID.
func GetDataByIDCommand() *cli.Command {
	command := cli.Command{
		Name:  "get_data_by_id",
		Usage: "get user's data by ID",
		Args:  true,
		Flags: []cli.Flag{
			&cli.Int64Flag{
				Name:     "data-id",
				Required: true,
				Usage:    "data ID",
			},
		},
		Action: func(cCtx *cli.Context) error {
			dataID := cCtx.Int64("data-id")
			service, err := NewCommand()
			if err != nil {
				return err
			}
			err = service.GetDataByID(dataID)
			if err != nil {
				return err
			}
			return nil
		},
	}
	return &command
}

// Добавление текстовых данных пользователем.
func AddTextDataCommand() *cli.Command {
	command := cli.Command{
		Name:  "add_text_data",
		Usage: "add text data",
		Args:  true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "name",
				Required: true,
				Usage:    "data name",
			},
			&cli.StringFlag{
				Name:     "data",
				Required: true,
				Usage:    "user's secured data",
			},
		},
		Action: func(cCtx *cli.Context) error {
			textData := model.AddTextData{
				Name: cCtx.String("name"),
				Data: cCtx.String("data"),
			}
			service, err := NewCommand()
			if err != nil {
				return err
			}
			err = service.AddTextData(textData)
			if err != nil {
				return err
			}
			return nil
		},
	}
	return &command
}

// Обновление текстовых данных пользователем.
func UpdateTextDataCommand() *cli.Command {
	command := cli.Command{
		Name:  "update_text_data",
		Usage: "update text data",
		Args:  true,
		Flags: []cli.Flag{
			&cli.Int64Flag{
				Name:     "data-id",
				Required: true,
				Usage:    "data ID",
			},
			&cli.StringFlag{
				Name:     "name",
				Required: true,
				Usage:    "data name",
			},
			&cli.StringFlag{
				Name:     "data",
				Required: true,
				Usage:    "user's secured data",
			},
		},
		Action: func(cCtx *cli.Context) error {
			dataID := cCtx.Int64("data-id")
			textData := model.AddTextData{
				Name: cCtx.String("name"),
				Data: cCtx.String("data"),
			}
			service, err := NewCommand()
			if err != nil {
				return err
			}
			err = service.UpdateTextData(textData, dataID)
			if err != nil {
				return err
			}
			return nil
		},
	}
	return &command
}

// Добавление бинарных данных пользователем.
func AddBinaryDataCommand() *cli.Command {
	command := cli.Command{
		Name:  "add_binary_data",
		Usage: "add binary data",
		Args:  true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "name",
				Required: true,
				Usage:    "data name",
			},
			&cli.StringFlag{
				Name:     "filepath",
				Required: true,
				Usage:    "binary data filepath",
			},
		},
		Action: func(cCtx *cli.Context) error {
			name := cCtx.String("name")
			filepath := cCtx.String("filepath")
			service, err := NewCommand()
			if err != nil {
				return err
			}
			err = service.AddBinaryData(filepath, name)
			if err != nil {
				return err
			}
			return nil
		},
	}
	return &command
}

// Обновление бинарных данных пользователем.
func UpdateBinaryData() *cli.Command {
	command := cli.Command{
		Name:  "update_binary_data",
		Usage: "update binary data",
		Args:  true,
		Flags: []cli.Flag{
			&cli.Int64Flag{
				Name:     "data-id",
				Required: true,
				Usage:    "data ID",
			},
			&cli.StringFlag{
				Name:     "name",
				Required: true,
				Usage:    "data name",
			},
			&cli.StringFlag{
				Name:     "filepath",
				Required: true,
				Usage:    "binary data filepath",
			},
		},
		Action: func(cCtx *cli.Context) error {
			dataID := cCtx.Int64("data-id")
			name := cCtx.String("name")
			filepath := cCtx.String("filepath")
			service, err := NewCommand()
			if err != nil {
				return err
			}
			err = service.UpdateBinaryData(filepath, name, dataID)
			if err != nil {
				return err
			}
			return nil
		},
	}
	return &command
}

// Удаление данных пользователя по ID.
func DeleteDataCommand() *cli.Command {
	command := cli.Command{
		Name:  "delete_data",
		Usage: "delete data",
		Args:  true,
		Flags: []cli.Flag{
			&cli.Int64Flag{
				Name:     "data-id",
				Required: true,
				Usage:    "data ID",
			},
		},
		Action: func(cCtx *cli.Context) error {
			dataID := cCtx.Int64("data-id")
			service, err := NewCommand()
			if err != nil {
				return err
			}
			err = service.DeleteData(dataID)
			if err != nil {
				return err
			}
			return nil
		},
	}
	return &command
}
