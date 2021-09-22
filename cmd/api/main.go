package main

import (
	"github.com/ditointernet/tradulab-service/database"
	"github.com/ditointernet/tradulab-service/internal/core/services"
	"github.com/ditointernet/tradulab-service/internal/rest"
)

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
