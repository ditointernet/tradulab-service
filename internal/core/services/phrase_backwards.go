package services

import (
	"context"
	"errors"

	"github.com/ditointernet/tradulab-service/internal/core/domain"
)

type PhraseBackward struct{}

func (p *PhraseBackward) GetByID(ctx context.Context, id string) (domain.Phrase, error) {
	if phrase, ok := phrases[id]; ok {
		runes := []rune(phrase.Key)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return domain.Phrase{
			FileID: phrase.FileID,
			Key:    string(runes),
		}, nil
	}
	return domain.Phrase{}, errors.New("phrase not found")
}
