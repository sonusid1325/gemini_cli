package main

import (
	"gemini/cli"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal("GEMINI_API_KEY not set in .env")
	}

	args := os.Args[1:]

	if len(args) == 0 {
		for {
			cli.Init(apiKey)
		}
	}
}
