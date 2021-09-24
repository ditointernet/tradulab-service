package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ditointernet/tradulab-service/database"
	"github.com/ditointernet/tradulab-service/internal/core/services"
	"github.com/ditointernet/tradulab-service/internal/rest"
	"github.com/joho/godotenv"
)

type config struct {
	user     string
	host     string
	password string
	dbName   string
	port     string
}

func GoDotEnvVariable() (*config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	c := &config{
		user:     os.Getenv("POSTGRES_USER"),
		host:     os.Getenv("HOST"),
		password: os.Getenv("POSTGRES_PASSWORD"),
		dbName:   os.Getenv("POSTGRES_DB"),
		port:     os.Getenv("PORT_POSTGRES"),
	}

	return c, nil
}

func main() {

	env, err := GoDotEnvVariable()
	if err != nil {
		fmt.Println("Error during environment variables build", err.Error())
		return
	}

	postgres := database.NewConfig(&database.ConfigDB{
		User:     env.user,
		Host:     env.host,
		Password: env.password,
		DbName:   env.dbName,
		Port:     env.port,
	})

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
