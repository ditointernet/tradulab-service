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
	handler Handler
	strg    storage.Storage
	sFile   services.File
}

type FileName struct {
	Name string
}

func MustNewSubscriber(sFile services.File, strg storage.Storage, handler Handler) *Subscriber {
	return &Subscriber{
		sFile:   sFile,
		strg:    strg,
		handler: handler,
	}
}

func (s Subscriber) HandleMessage(ctx context.Context, m *pubsub.Message) error {
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

	rc, err := s.DownloadDoc(ctx, fileName.Name)
	if err != nil {
		return err
	}
	s.handler.Process(ctx, rc, file.ID)

	m.Ack()
	return nil
}

func (s Subscriber) DownloadDoc(ctx context.Context, docName string) (*googleStorage.Reader, error) {
	rc, err := s.strg.BucketHandle.Object(docName).NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	return rc, nil
}
