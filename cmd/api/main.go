package main

import (
	"context"
	"fmt"

	"github.com/ditointernet/tradulab-service/adapters"
	"github.com/ditointernet/tradulab-service/internal/core/services"
	"github.com/ditointernet/tradulab-service/internal/repository"
	"github.com/ditointernet/tradulab-service/internal/rest"
	"github.com/ditointernet/tradulab-service/internal/storage"
)

func main() {

	env, err := adapters.GoDotEnvVariable()
	if err != nil {
		fmt.Println("Error during environment variables build", err.Error())
		return
	}

	db := adapters.NewDatabase(&adapters.Config{
		User:     env.User,
		Host:     env.Host,
		Password: env.Password,
		DbName:   env.DbName,
		Port:     env.Port,
	})

	server := MustNewServer()

	sql, err := db.DB()
	if err != nil {
		panic(err)
	}

	storage := storage.MustNewStorage(
		context.Background(),
		env.ProjectID,
		env.BucketName,
		env.ExpirationTime,
		env.AllowedType,
	)
	pRepository := repository.MustNewPhrase(sql)
	pService := services.MustNewPhrase(pRepository, storage)

	rPhrase := rest.MustNewPhrase(pService)
	fRepository := repository.MustNewFile(sql)

	fService := services.MustNewFile(fRepository, storage)
	rFile, err := rest.NewFile(rest.ServiceInput{
		File: fService,
	})
	if err != nil {
		panic(err)
	}
	router := server.Listen()

	router.POST("/files", rFile.CreateFile)
	router.GET("/files", rFile.GetProjectFiles)
	router.POST("/files/:id/signed-url", rFile.CreateSignedURL)
	router.GET("/phrases/:id", rPhrase.GetPhrasesById)
	router.GET("/phrases", rPhrase.GetFilePhrases)

	router.Run()
}
