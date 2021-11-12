package services

import (
	"context"
	"database/sql"

	"github.com/ditointernet/tradulab-service/internal/core/domain"
	"github.com/ditointernet/tradulab-service/internal/repository"
	"github.com/ditointernet/tradulab-service/internal/storage"
	"github.com/google/uuid"
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

func (p *Phrase) VerifyInDB(ctx context.Context, entry *domain.Phrase) error {
	result, err := p.repo.GetPhrase(ctx, *entry)

	if err == nil {
		if result.Content != entry.Content {
			p.repo.UpdatePhrase(ctx, *entry)
		}
		return nil
	}

	return err
}

func (p *Phrase) CreatePhrase(ctx context.Context, entry *domain.Phrase) (domain.Phrase, error) {
	err := p.VerifyInDB(ctx, entry)

	id := uuid.New().String()

	newPhrase := domain.Phrase{
		ID:      id,
		FileID:  entry.FileID,
		Key:     entry.Key,
		Content: entry.Content,
	}

	if err != nil && err == sql.ErrNoRows {
		err = p.repo.CreatePhrase(ctx, newPhrase)
		if err != nil {
			return domain.Phrase{}, err
		}
	}

	return newPhrase, nil
}
