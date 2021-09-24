package main

import (
	"log"
	"os"

	"github.com/ditointernet/tradulab-service/database"
	"github.com/ditointernet/tradulab-service/internal/core/services"
	"github.com/ditointernet/tradulab-service/internal/rest"
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
		Port:     os.Getenv("PORT_POSTGRES"),
	}

	return c, nil
}

func main() {

	server := MustNewServer()
	db := database.MustNewDB()

	db.StartPostgres()
	fService := services.MustNewFile(db)

	router := server.Listen()
	rPhrase := rest.MustNewPhrase()
	rFile, err := rest.MustNewFile(rest.ServiceInput{
		File: fService,
	})
	if err != nil {
		panic(err)
	}

	router.GET("/:id", rPhrase.FindByID)
	router.POST("/file", rFile.CreateFile)

	router.Run()
}
