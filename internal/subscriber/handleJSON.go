package subscriber

import (
	"context"
	"encoding/json"
	"fmt"
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
	var phrasesInFile []string
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

		phrasesInFile = append(phrasesInFile, fmt.Sprintf("'%s'", phrase.Key))
		h.sPhrase.HandlePhrase(ctx, phrase)
	}

	h.sPhrase.CleanDB(ctx, phrasesInFile, fileID)

	return nil
}
