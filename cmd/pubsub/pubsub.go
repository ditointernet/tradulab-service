package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
	"github.com/ditointernet/tradulab-service/adapters"
	"github.com/ditointernet/tradulab-service/internal/core/services"
	"github.com/ditointernet/tradulab-service/internal/repository"
	"github.com/ditointernet/tradulab-service/internal/storage"
	"github.com/ditointernet/tradulab-service/internal/subscriber"
	"google.golang.org/api/option"
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

	cred := &adapters.Config{
		Credentials:  env.Credentials,
		ProjectID:    env.ProjectID,
		Subscription: env.Subscription,
	}

	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, cred.ProjectID, option.WithCredentialsFile(cred.Credentials))
	if err != nil {
		panic(err)
	}

	log.Println("Listening to subscription")
	sub := client.Subscription(cred.Subscription)
	sub.ReceiveSettings.Synchronous = true
	sub.ReceiveSettings.MaxOutstandingMessages = 1
	message := subscriber.MustNewSubscriber(*fService, storage)
	err = sub.Receive(ctx, func(c context.Context, m *pubsub.Message) {
		err := message.HandleMessage(c, m)
		if err != nil {
			fmt.Println("Couldn't handle message", err.Error())
		}
	})
	if err != nil {
		fmt.Println("Error receiving message", err.Error())
	}
}
