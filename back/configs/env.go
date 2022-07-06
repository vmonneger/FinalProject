// This file check if environment variables are loaded.
package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Return env variables to connect DB.
func EnvDB() (string, string) {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("DB_USERNAME"), os.Getenv("DB_PASS")
}

func EnvJwtSecret() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("JWT_SECRET")
}

func EnvServerPort() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("SERVER_PORT")
}
