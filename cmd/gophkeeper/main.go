package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/pisarevaa/gophkeeper/internal/agent/command"
)

func main() {
	app := &cli.App{
		Name:        "gophkeeper",
		Usage:       "Keep your data safely",
		Description: "Gophkeeper helps to keep your data safely",
		Commands: []*cli.Command{
			command.RegisterCommand(),
			command.LoginCommand(),
			command.GetDataCommand(),
			command.GetDataByIDCommand(),
			command.AddTextDataCommand(),
			command.UpdateTextDataCommand(),
			command.AddBinaryDataCommand(),
			command.UpdateBinaryData(),
			command.DeleteDataCommand(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
