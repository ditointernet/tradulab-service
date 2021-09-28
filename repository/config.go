package repository

import (
	"os"
)

// refatorar
type Config struct {
	User     string
	Host     string
	Password string
	DbName   string
	Port     string
}

func GoDotEnvVariable() (*Config, error) {

	c := &Config{
		User:     os.Getenv("POSTGRES_USER"),
		Host:     os.Getenv("HOST"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DbName:   os.Getenv("POSTGRES_DB"),
		Port:     os.Getenv("PORTPOSTGRES"),
	}

	return c, nil
}
