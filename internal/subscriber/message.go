package subscriber

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"cloud.google.com/go/pubsub"
	"github.com/ditointernet/tradulab-service/internal/core/domain"
	"github.com/ditointernet/tradulab-service/internal/core/services"
	"github.com/ditointernet/tradulab-service/internal/storage"
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

	err = DownloadDoc(ctx, fileName.Name, strg)

	if err != nil {
		return err
	}

	m.Ack()
	return nil
}

func DownloadDoc(ctx context.Context, docName string, strg storage.Storage) error {
	rc, err := strg.BucketHandle.Object(docName).NewReader(ctx)
	if err != nil {
		return err
	}
	defer rc.Close()

	d, err := ioutil.ReadAll(rc)
	if err != nil {
		return err
	}
	fmt.Printf("Blob %s downloaded.\n", docName)
	fmt.Println(string(d))

	return nil
}
