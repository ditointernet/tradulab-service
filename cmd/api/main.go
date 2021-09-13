package main

import (
	"github.com/ditointernet/tradulab-service/internal/rest"
)

func main() {
	server := MustNewServer()

	router := server.Listen()
	rPhrase := rest.MustNewPhrase()

	router.GET("/:id", rPhrase.FindByID)

	router.Run()
}
