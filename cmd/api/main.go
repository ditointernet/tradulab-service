package main

import (
	"fmt"

	"github.com/ditointernet/tradulab-service/database"
	"github.com/ditointernet/tradulab-service/internal/core/services"
	"github.com/ditointernet/tradulab-service/internal/rest"
)

func main() {

	env, err := database.GoDotEnvVariable()
	if err != nil {
		fmt.Println("Error during environment variables build", err.Error())
		return
	}

	db := database.NewConfig(&database.ConfigDB{
		User:     env.User,
		Host:     env.Host,
		Password: env.Password,
		DbName:   env.DbName,
		Port:     env.Port,
	})

	server := MustNewServer()

	db.StartPostgres()
	fService := services.MustNewFile(*db)

	router := server.Listen()
	rPhrase := rest.MustNewPhrase()
	rFile, err := rest.NewFile(rest.ServiceInput{
		File: fService,
	})
	if err != nil {
		panic(err)
	}

	router.GET("/:id", rPhrase.FindByID)
	router.POST("/file", rFile.CreateFile)

	router.Run()
}
