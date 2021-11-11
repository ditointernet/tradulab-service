package subscriber

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	"cloud.google.com/go/storage"
	"github.com/ditointernet/tradulab-service/internal/core/domain"
	"github.com/ditointernet/tradulab-service/internal/core/services"
)

type Handler struct {
	sPhrase services.Phrase
}

func MustNewHandler(sPhrase services.Phrase) *Handler {
	return &Handler{
		sPhrase: sPhrase,
	}
}

func (h Handler) Process(ctx context.Context, rc *storage.Reader, fileID string) error {
	d, err := ioutil.ReadAll(rc)
	if err != nil {
		return err
	}
	m := map[string]interface{}{}
	err = json.Unmarshal(d, &m)
	if err != nil {
		log.Fatal(err)
	}

	for index := range m {
		phrase := &domain.Phrase{
			FileID:  fileID,
			Key:     index,
			Content: m[index].(string),
		}

		h.sPhrase.CreatePhrase(ctx, phrase)

	}

	return nil
}
