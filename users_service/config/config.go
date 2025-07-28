package config

import (
	"errors"
	"log"
	"os"
)

type Config struct {
	DatabaseName     string
	DatabasePassword string
	DatabaseUser     string
	DatabaseHost     string
	DatabasePort     string
	DatabaseAddress  string
}

func LoadConfig() (*Config, error) {

	dbName := os.Getenv("DATABASE_NAME")
	dbPassword := os.Getenv("DATABASE_PASSWORD")
	dbUser := os.Getenv("DATABASE_USER")
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")

	if dbName == "" || dbPassword == "" || dbUser == "" || dbHost == "" || dbPort == "" {
		return nil, errors.New("missing database configuration variables")

	}
	if dbHost == "" {
		dbHost = "localhost"
	}
	if dbPort == "" {
		dbPort = "3306"
	}

	return &Config{
		DatabaseName:     dbName,
		DatabasePassword: dbPassword,
		DatabaseUser:     dbUser,
		DatabaseHost:     dbHost,
		DatabasePort:     dbPort,
		DatabaseAddress:  dbHost + ":" + dbPort,
	}, nil
}

func GetDatabaseUser() string {
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	return cfg.DatabaseUser
}

func GetDatabasePassword() string {
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	return cfg.DatabasePassword
}

func GetDatabaseAddress() string {
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	return cfg.DatabaseAddress
}

func GetDatabaseName() string {
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	return cfg.DatabaseName
}
