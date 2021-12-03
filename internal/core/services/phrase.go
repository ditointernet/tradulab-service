package services

import (
	"context"

	"github.com/ditointernet/tradulab-service/internal/core/domain"
	"github.com/ditointernet/tradulab-service/internal/repository"
	"github.com/ditointernet/tradulab-service/internal/storage"
	"github.com/google/uuid"
)

type Phrase struct {
	repo repository.PhraseRepository
}

func MustNewPhrase(repo repository.PhraseRepository, storage storage.FileStorage) *Phrase {
	return &Phrase{
		repo: repo,
	}
}

func (p *Phrase) CreateOrUpdatePhrase(ctx context.Context, entry *domain.Phrase) (domain.Phrase, error) {

	newPhrase := domain.Phrase{
		ID:      uuid.New().String(),
		FileID:  entry.FileID,
		Key:     entry.Key,
		Content: entry.Content,
	}
	err := p.repo.CreateOrUpdatePhrase(ctx, newPhrase)
	if err != nil {
		return domain.Phrase{}, err
	}

	return newPhrase, nil
}

func (p *Phrase) CleanDB(ctx context.Context, phrasesKey []string, fileId string) error {
	err := p.repo.DeletePhrases(ctx, phrasesKey, fileId)
	if err != nil {
		return err
	}

	return nil
}
