package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func GetEnvVar(key string) (string, error) {
	token := os.Getenv(key)
	if token == "" {
		// Token not found, could be running locally
		err := godotenv.Load(".env")
		if err != nil {
			return "", fmt.Errorf("Error loading .env file")
		}
	}
	token = os.Getenv(key)

	if token == "" {
		return "", fmt.Errorf("token not found")
	}

	return token, nil
}
