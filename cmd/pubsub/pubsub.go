package main

import (
	"context"
	"encoding/json"
	"fmt"

	"cloud.google.com/go/pubsub"
	"github.com/ditointernet/tradulab-service/adapters"
	"github.com/ditointernet/tradulab-service/internal/core/domain"
	"github.com/ditointernet/tradulab-service/internal/core/services"
	"github.com/ditointernet/tradulab-service/internal/repository"
	"github.com/ditointernet/tradulab-service/internal/rest"
	"github.com/ditointernet/tradulab-service/internal/storage"
	"google.golang.org/api/option"
)

type FileName struct {
	Name string
}

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

	sql, err := db.DB()
	if err != nil {
		panic(err)
	}

	fRepository := repository.MustNewFile(sql)
	storage := storage.MustNewStorage(
		context.Background(),
		env.ProjectID,
		env.BucketName,
		env.ExpirationTime,
		env.AllowedType,
	)
	fService := services.MustNewFile(fRepository, storage)

	rFile, err := rest.NewFile(rest.ServiceInput{
		File: fService,
	})
	if err != nil {
		panic(err)
	}

	cred := &adapters.Config{
		Credentials: env.Credentials,
		ProjectID:   env.ProjectID,
	}

	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, cred.ProjectID, option.WithCredentialsFile(cred.Credentials))
	if err != nil {
		panic(err)
	}

	fmt.Println("Listening to subscription")
	sub := client.Subscription("files-topic-sub")
	sub.ReceiveSettings.Synchronous = true
	sub.ReceiveSettings.MaxOutstandingMessages = 1
	err = sub.Receive(ctx, func(c context.Context, m *pubsub.Message) {
		fmt.Println("new message received")
		var fileName FileName

		data := m.Data

		err := json.Unmarshal(data, &fileName)
		if err != nil {
			panic(err)
		}

		file := &domain.File{
			ID: fileName.Name,
		}

		rFile.EditFile(ctx, file.ID)
		fmt.Println("file uploaded")
		m.Ack()
	})
	if err != nil {
		panic(err)
	}
}
