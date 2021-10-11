package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/storage"
	"github.com/ditointernet/tradulab-service/adapters"
	"github.com/ditointernet/tradulab-service/internal/core/services"
	"github.com/ditointernet/tradulab-service/internal/repository"
	"github.com/ditointernet/tradulab-service/internal/rest"
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

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		fmt.Println("Could not connect to the Google Cloud Storage")
	}

	if err := client.Close(); err != nil {
		fmt.Println("Could not close the GCS")
	}
	store22 := adapters.MustNewGCS(client, ctx)
	store22.ListBucket()

	fRepository := repository.MustNewFile(sql)
	fService := services.MustNewFile(fRepository)

	router := server.Listen()
	// rPhrase := rest.MustNewPhrase()
	rFile, err := rest.NewFile(rest.ServiceInput{
		File: fService,
	})
	if err != nil {
		panic(err)
	}

	// router.GET("/:id", rPhrase.FindByID)
	router.POST("/file", rFile.CreateFile)
	router.GET("/file", rFile.GetAllFiles)
	router.PUT("/file/:id", rFile.EditFile)

	router.Run()
}
