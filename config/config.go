package config

import (
	"os"
)

type Config struct {
	DatabaseURL  string
	DatabaseName string
}

func LoadConfig() (*Config, error) {
	dbURL := os.Getenv("DATABASE_URL")
	dbName := os.Getenv("DATABASE_NAME")

	if dbURL == "" || dbName == "" {
		return nil, errors.New("missing required environment variables")
	}

	return &Config{
		DatabaseURL:  dbURL,
		DatabaseName: dbName,
	}, nil
}