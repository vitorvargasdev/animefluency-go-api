package commands

import (
	"github.com/urfave/cli/v2"
	"github.com/vitorvargasdev/animefluency-go-api/internal/database"
)

func NewMigrateCommand() *cli.Command {
	return &cli.Command{
		Name:  "migrate",
		Usage: "Handle database migrations",
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
				Usage: "Run all pending migrations",
				Action: func(ctx *cli.Context) error {
					return database.Up()
				},
			},
			{
				Name:  "down",
				Usage: "Revert the last migration",
				Action: func(ctx *cli.Context) error {
					return database.Down()
				},
			},
			{
				Name: "drop",
				Usage: "Drop the database",
				Action: func(ctx *cli.Context) error {
					return database.Drop()
				},
			},
			{
				Name:   "reset",
				Usage:  "Drop and run all migrations",
				Action: func(ctx *cli.Context) error { 
					return database.Reset()
				},
			},
		},
	}
}
