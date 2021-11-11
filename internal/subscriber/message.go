package subscriber

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	"cloud.google.com/go/pubsub"
	googleStorage "cloud.google.com/go/storage"
	"github.com/ditointernet/tradulab-service/internal/core/domain"
	"github.com/ditointernet/tradulab-service/internal/core/services"
	"github.com/ditointernet/tradulab-service/internal/storage"
)

type Subscriber struct {
	sFile   services.File
	handler Handler
}

type FileName struct {
	Name string
}

func MustNewSubscriber(sFile services.File, handler Handler) *Subscriber {
	return &Subscriber{
		sFile:   sFile,
		handler: handler,
	}
}

func (s Subscriber) HandleMessage(ctx context.Context, m *pubsub.Message, strg storage.Storage) error {
	log.Println("new message received")
	var fileName FileName

	data := m.Data
	err := json.Unmarshal(data, &fileName)
	if err != nil {
		m.Ack()
		return err
	}

	filename := strings.Split(fileName.Name, ".")

	file := &domain.File{
		ID: filename[0],
	}

	err = s.sFile.SetUploadSuccessful(ctx, file)
	if err != nil {
		m.Nack()
		return err
	}
	log.Println("file uploaded")

	rc, err := DownloadDoc(ctx, fileName.Name, strg)
	if err != nil {
		return err
	}
	s.handler.Process(ctx, rc, file.ID)

	m.Ack()
	return nil
}

func DownloadDoc(ctx context.Context, docName string, strg storage.Storage) (*googleStorage.Reader, error) {
	rc, err := strg.BucketHandle.Object(docName).NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	return rc, nil
}
