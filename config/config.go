package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Init(configFilePath string) string {
	err := godotenv.Load(configFilePath)
	if err != nil {
		log.Fatal("error with finding .env file")
	}
	port := os.Getenv("PORT")

	return port
}
