package services

import (
	"context"

	"github.com/ditointernet/tradulab-service/internal/core/domain"
	"github.com/ditointernet/tradulab-service/internal/repository"
	"github.com/ditointernet/tradulab-service/internal/storage"
)

type Phrase struct {
	repo repository.PhraseRepository
	//storage storage.FileStorage
}

func MustNewPhrase(repo repository.PhraseRepository, storage storage.FileStorage) *Phrase {
	return &Phrase{
		repo: repo,
		//storage: storage,
	}
}

func (p *Phrase) CreatePhrase(ctx context.Context, entry *domain.Phrase) (domain.Phrase, error) {
	newPhrase := domain.Phrase{
		FileID:  entry.FileID,
		Key:     entry.Key,
		Content: entry.Content,
	}
	err := p.repo.CreatePhrase(ctx, newPhrase)
	if err != nil {
		return domain.Phrase{}, err
	}

	return newPhrase, nil
}
