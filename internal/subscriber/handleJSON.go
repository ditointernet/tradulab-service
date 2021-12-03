package subscriber

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"

	"cloud.google.com/go/storage"
	"github.com/ditointernet/tradulab-service/internal/core/domain"
	"github.com/ditointernet/tradulab-service/internal/core/services"
)

type Handler struct {
	sPhrase services.Phrase
}

func MustNewHandlerJSON(sPhrase services.Phrase) *Handler {
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
	m := make(map[string]string)
	err = json.Unmarshal(d, &m)
	if err != nil {
		return errors.New("fail in unmarshal or json format")
	}

	for key, value := range m {
		phrase := &domain.Phrase{
			FileID:  fileID,
			Key:     key,
			Content: value,
		}

		phrasesInFile = append(phrasesInFile, phrase.Key)
		_, err := h.sPhrase.CreateOrUpdatePhrase(ctx, phrase)
		if err != nil {
			log.Println(err.Error())
			return err
		}
	}

	err = h.sPhrase.CleanDB(ctx, phrasesInFile, fileID)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
