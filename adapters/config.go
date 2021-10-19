package adapters

import (
	"os"
	"strconv"
)

type Config struct {
	User     string
	Host     string
	Password string
	DbName   string
	Port     string

	ProjectID      string
	BucketName     string
	AllowedType    string
	ExpirationTime int
}

func GoDotEnvVariable() (*Config, error) {

	expTime, _ := strconv.Atoi(os.Getenv("EXPIRATION_TIME"))

	c := &Config{
		User:           os.Getenv("POSTGRES_USER"),
		Host:           os.Getenv("HOST"),
		Password:       os.Getenv("POSTGRES_PASSWORD"),
		DbName:         os.Getenv("POSTGRES_DB"),
		Port:           os.Getenv("PORTPOSTGRES"),
		ProjectID:      os.Getenv("PROJECT_ID"),
		BucketName:     os.Getenv("BUCKET_NAME"),
		AllowedType:    os.Getenv("ALLOWED_TYPE"),
		ExpirationTime: expTime,
	}

	return c, nil
}
