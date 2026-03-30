package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_Host     string
	DB_Port     string
	DB_User     string
	DB_Password string
	Uri         string
}

func LoadConfig() *Config {
	errEnv := godotenv.Load()
	if errEnv != nil {
		fmt.Println("Error loading .env file")
		return nil
	}

	uri := os.Getenv("MONGODB_URI")
	fmt.Println("MONGODB_URI :", uri)

	if uri == "" {
		log.Fatal("MongoDB URI is empty")
	}

	config := &Config{
		DB_Host:     os.Getenv("DB_HOST"),
		DB_Port:     os.Getenv("DB_PORT"),
		DB_User:     os.Getenv("DB_USER"),
		DB_Password: os.Getenv("DB_PASSWORD"),
		Uri:         os.Getenv("MONGODB_URI"),
	}

	if config.DB_Host == "" {
		log.Println("WARNING: DB_HOST is empty! Check your .env file and path")
	}

	return config
}
