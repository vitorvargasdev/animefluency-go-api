package commands

import (
	"github.com/urfave/cli/v2"
	"github.com/vitorvargasdev/animefluency-go-api/internal/server"
)

func NewStartCommand() *cli.Command {
	return &cli.Command{
		Name:  "start",
		Usage: "Starts the API",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "port",
				Value:   3000,
				Aliases: []string{"p"},
				Usage:   "Port to run the server on",
			},
		},
		Action: func(ctx *cli.Context) error {
			port := ctx.Int("port")

			if err := server.StartServer(port); err != nil {
				return err
			}

			return nil
		},
	}
}
