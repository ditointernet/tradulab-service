package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
	"github.com/ditointernet/tradulab-service/adapters"
	"google.golang.org/api/option"
)

func main() {
	env, err := adapters.GoDotEnvVariable()
	if err != nil {
		fmt.Println("Error during environment variables build", err.Error())
		return
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
		fmt.Printf("%s", string(m.Data))
		fmt.Println("new message received")
		m.Ack()
	})
	if err != nil {
		panic(err)
	}
}
