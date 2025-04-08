package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/vitorvargasdev/animefluency-go-api/cmd/cli"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("error loading .env file: %w", err)
	}

	if err := cli.StartCLI(); err != nil {
		log.Fatal(err)
	}
}
