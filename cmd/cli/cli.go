package cli

import (
    "log"
    "os"

    "github.com/vitorvargasdev/animefluency-go-api/internal/cli/commands" 
    "github.com/urfave/cli/v2"
)

func StartCLI() {
    app := &cli.App{
        Name:  "AnimeFluency API",
        Usage: "A CLI tool to manage the AnimeFluency API",
        Commands: []*cli.Command{
            commands.NewStartCommand(),
        },
    }

    if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}

