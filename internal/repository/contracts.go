package repository

import (
	"context"

	"github.com/ditointernet/tradulab-service/internal/core/domain"
)

type FileRepository interface {
	CreateFile(ctx context.Context, file domain.File) error
	FindFile(ctx context.Context, id string) (domain.File, error)
	SetUploadSuccessful(ctx context.Context, file *domain.File) error
	GetFiles(ctx context.Context, projectId string) ([]domain.File, error)
	FindProject(ctx context.Context, projectId string) (string, error)
}

type PhraseRepository interface {
	CreatePhrase(ctx context.Context, entry domain.Phrase) error
	GetPhrase(ctx context.Context, entry domain.Phrase) (domain.Phrase, error)
	UpdatePhrase(ctx context.Context, entry domain.Phrase) error
	GetByFileId(ctx context.Context, id string) (domain.Phrase, error)
	DeletePhrases(ctx context.Context, phrasesKey []string, projectId string) error
}
