package services

import (
	"context"
	"errors"

	"github.com/ditointernet/tradulab-service/internal/core/domain"
)

var phrases map[string]domain.Phrase = map[string]domain.Phrase{
	"1": {
		FileID: "1",
		Key:    "pipoca doce",
	},
}

type Phrase struct{}

func (p *Phrase) GetByID(ctx context.Context, id string) (domain.Phrase, error) {
	if phrase, ok := phrases[id]; ok {
		return phrase, nil
	}
	return domain.Phrase{}, errors.New("phrase not found")
}
