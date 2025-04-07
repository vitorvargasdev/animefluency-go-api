package main

import (
	"log"

	"github.com/vitorvargasdev/animefluency-go-api/cmd/cli"
)

func main() {
	if err := cli.StartCLI(); err != nil {
		log.Fatal(err)
	}
}
