package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	User     string
	Host     string
	Password string
	DbName   string
	Port     string
}

func GoDotEnvVariable() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	c := &Config{
		User:     os.Getenv("POSTGRES_USER"),
		Host:     os.Getenv("HOST"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DbName:   os.Getenv("POSTGRES_DB"),
		Port:     os.Getenv("PORTPOSTGRES"),
	}

	return c, nil
}
