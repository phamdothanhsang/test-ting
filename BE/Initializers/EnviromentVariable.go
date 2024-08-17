package Initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvironmentVariables() {
	//Load environment variable
	var err error = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
