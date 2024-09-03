package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/sickodev/grandmaster-llm/cmd"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading environment file")
	}

	cmd.Execute()
}
