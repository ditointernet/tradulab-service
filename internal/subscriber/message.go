package subscriber

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"cloud.google.com/go/pubsub"
	"github.com/ditointernet/tradulab-service/internal/core/domain"
	"github.com/ditointernet/tradulab-service/internal/core/services"
)

type Subscriber struct {
	sFile services.File
}

type FileName struct {
	Name string
}

func MustNewSubscriber(sFile services.File) *Subscriber {
	return &Subscriber{
		sFile: sFile,
	}
}

func (s Subscriber) HandleMessage(ctx context.Context, m *pubsub.Message) error {
	log.Println("new message received")
	var fileName FileName

	data := m.Data
	fmt.Println(string(data))
	err := json.Unmarshal(data, &fileName)
	if err != nil {
		return err
	}

	filename := strings.Split(fileName.Name, ".")

	file := &domain.File{
		ID: filename[0],
	}

	err = s.sFile.EditFile(ctx, file)
	if err != nil {
		return err
	}

	log.Println("file uploaded")

	m.Ack()
	return nil
}
