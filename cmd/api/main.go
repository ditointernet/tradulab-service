package main

import (
	"github.com/ditointernet/tradulab-service/internal/rest"
)

func main() {
	server := MustNewServer()

	router := server.Listen()
	rPhrase := rest.MustNewPhrase()
	rFile := rest.MustNewFile()

	router.GET("/:id", rPhrase.FindByID)
	router.POST("/file", rFile.CreateFile)

	router.Run()
}
