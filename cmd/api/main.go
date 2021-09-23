package main

import (
	"github.com/ditointernet/tradulab-service/internal/core/services"
	"github.com/ditointernet/tradulab-service/internal/rest"
)

func main() {
	server := MustNewServer()
	fService := services.MustNewFile()

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
