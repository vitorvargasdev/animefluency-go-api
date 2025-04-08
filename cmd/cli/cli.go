package cli

import (
	"os"

	"github.com/urfave/cli/v2"
	"github.com/vitorvargasdev/animefluency-go-api/internal/cli/commands"
)

func StartCLI() error {
	app := &cli.App{
		Name:  "AnimeFluency API",
		Usage: "A CLI tool to manage the AnimeFluency API",
		Commands: []*cli.Command{
			commands.NewStartCommand(),
			commands.NewMigrateCommand(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		return err
	}

	return nil
}
