package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/joosejunsheng/jhft/cmd"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	cmd.Run()
}
