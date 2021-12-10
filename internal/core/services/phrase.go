package services

import (
	"context"
	"strconv"

	"github.com/ditointernet/tradulab-service/internal/core/domain"
	"github.com/ditointernet/tradulab-service/internal/repository"
	"github.com/ditointernet/tradulab-service/internal/storage"
	"github.com/pkg/errors"
)

type Phrase struct {
	repo repository.PhraseRepository
}

func MustNewPhrase(repo repository.PhraseRepository, storage storage.FileStorage) *Phrase {
	return &Phrase{
		repo: repo,
	}
}

func (p *Phrase) CreateOrUpdatePhraseTx(ctx context.Context, entries []*domain.Phrase) error {

	err := p.repo.CreateOrUpdatePhraseTx(ctx, entries)
	if err != nil {
		return err
	}

	return nil
}

func (p *Phrase) CleanDB(ctx context.Context, phrasesKey []string, fileId string) error {
	err := p.repo.DeletePhrases(ctx, phrasesKey, fileId)
	if err != nil {
		return err
	}

	return nil
}

func (p *Phrase) GetPhrasesById(ctx context.Context, phraseId string) (domain.Phrase, error) {
	phrase, err := p.repo.GetPhrasesById(ctx, phraseId)

	if err != nil {
		return domain.Phrase{}, err
	}

	return phrase, nil
}

func (p *Phrase) GetFilePhrases(ctx context.Context, fileId, page string) ([]domain.Phrase, error) {
	numberPage, err := strconv.Atoi(page)
	if err != nil {
		return nil, err
	}

	if numberPage <= 0 {
		return nil, errors.New("page must be bigger than zero")
	}

	phrases, err := p.repo.GetFilePhrases(ctx, fileId, numberPage)
	if err != nil {
		return nil, err
	}

	return phrases, nil
}
