package commands

import (
	"github.com/urfave/cli/v2"
	"github.com/vitorvargasdev/animefluency-go-api/internal/database"
)

func NewMigrateCommand() *cli.Command {
	return &cli.Command{
		Name:  "migrate",
		Usage: "Migrate the database",
		Subcommands: []*cli.Command{
			{
				Name:  "create",
				Usage: "Create a new migration",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "name",
						Value:   "migration",
						Aliases: []string{"n"},
						Usage:   "Name of the migration",
					},
				},
				Action: func(ctx *cli.Context) error {
					name := ctx.String("name")
					return database.NewMigration(name)
				},
			},
			{
				Name:  "up",
				Usage: "Migrate the database",
				Action: func(ctx *cli.Context) error {
					return database.Up()
				},
			},
			{
				Name:  "down",
				Usage: "Migrate the database",
				Action: func(ctx *cli.Context) error {
					return database.Down()
				},
			},
			{
				Name: "Drop",
				Usage: "Migrate the database",
				Action: func(ctx *cli.Context) error {
					return database.Drop()
				},
			},
			{
				Name:   "reset",
				Usage:  "Migrate the database",
				Action: func(ctx *cli.Context) error { 
					return database.Reset()
				},
			},
		},
	}
}
