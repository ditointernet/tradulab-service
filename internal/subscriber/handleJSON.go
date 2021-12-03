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
	d, err := ioutil.ReadAll(rc)
	if err != nil {
		return err
	}
	m := make(map[string]string)
	err = json.Unmarshal(d, &m)
	if err != nil {
		return errors.New("fail in unmarshal or json format")
	}

	var phrasesKeys []string
	var allPhrases []*domain.Phrase

	for key, value := range m {
		phrase := &domain.Phrase{
			FileID:  fileID,
			Key:     key,
			Content: value,
		}

		phrasesKeys = append(phrasesKeys, phrase.Key)
		allPhrases = append(allPhrases, phrase)
	}
	err = h.sPhrase.CreateOrUpdatePhraseTx(ctx, allPhrases)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	err = h.sPhrase.CleanDB(ctx, phrasesKeys, fileID)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
