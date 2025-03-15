package main

import (
	"context"
	"fmt"
	"gemini/ai"
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

	ctx := context.Background()
	prompt := "Explain how AI works"

	response, err := ai.GenerateContent(ctx, apiKey, prompt)
	if err != nil {
		log.Fatalf("Error generating content: %v", err)
	}

	if len(response.Candidates) > 0 && len(response.Candidates[0].Content.Parts) > 0 {
		fmt.Println(response.Candidates[0].Content.Parts[0].Text)
	} else {
		fmt.Println("No response from the API.")
	}
}
