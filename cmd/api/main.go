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
	rFile, err := rest.MustNewFile(rest.ServiceInput{
		File: fService,
	})
	if err != nil {
		panic(err)
	}

	// phrase := services.Phrase{}
	phrase := services.PhraseBackward{}
	rPhrase := rest.MustNewPhrase(&phrase)
	router.GET("/:id", rPhrase.FindByID)
	router.POST("/file", rFile.CreateFile)

	router.Run()
}
