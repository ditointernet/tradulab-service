package handler

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	"cloud.google.com/go/pubsub"
	"github.com/ditointernet/tradulab-service/internal/core/domain"
	"github.com/ditointernet/tradulab-service/internal/rest"
)

type Message struct {
	Message *pubsub.Message
	rFile   rest.File
}

type FileName struct {
	Name string
}

func MustNewMessage(message *pubsub.Message, rFile rest.File) *Message {
	return &Message{
		Message: message,
		rFile:   rFile,
	}
}

func (m Message) HandleMessage(ctx context.Context) error {
	log.Println("new message received")
	var fileName FileName

	data := m.Message.Data

	err := json.Unmarshal(data, &fileName)
	if err != nil {
		return err
	}

	filename := strings.Split(fileName.Name, ".")

	file := &domain.File{
		ID: filename[0],
	}

	err = m.rFile.EditFile(ctx, file.ID)
	if err != nil {
		return err
	}

	log.Println("file uploaded")

	m.Message.Ack()
	return nil
}
