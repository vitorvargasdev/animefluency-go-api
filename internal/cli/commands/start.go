package commands

import (
	"github.com/urfave/cli/v2"
)

func NewStartCommand() *cli.Command {
	return &cli.Command{
		Name:  "start",
		Usage: "Starts the API",
		Action: func(ctx *cli.Context) error {
			return nil
		},
	}
}
