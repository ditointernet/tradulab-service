package repository

import (
	"context"

	"github.com/ditointernet/tradulab-service/internal/core/domain"
)

type FileRepository interface {
	CreateFile(ctx context.Context, file domain.File) error
	FindFile(ctx context.Context, id string) (domain.File, error)
	SetUploadSuccessful(ctx context.Context, file *domain.File) error
	GetProjectFiles(ctx context.Context, projectId string) ([]domain.File, error)
	SetCreateFail(ctx context.Context, file domain.File) error
}

type PhraseRepository interface {
	CreateOrUpdatePhraseTx(ctx context.Context, entries []*domain.Phrase) error
	GetByFileId(ctx context.Context, id string) (domain.Phrase, error)
	DeletePhrases(ctx context.Context, phrasesKey []string, projectId string) error
	GetPhrasesById(ctx context.Context, phraseId string) (domain.Phrase, error)
	GetFilePhrases(ctx context.Context, fileId string, page int) ([]domain.Phrase, error)
	CountPhrases(ctx context.Context, fileId string) (int, error)
}
