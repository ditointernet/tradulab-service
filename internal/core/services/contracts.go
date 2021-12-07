package services

import (
	"context"

	"github.com/ditointernet/tradulab-service/internal/core/domain"
)

type FileHandler interface {
	CreateFile(ctx context.Context, file *domain.File) (domain.File, error)
	SetUploadSuccessful(ctx context.Context, file *domain.File) error
	GetProjectFiles(ctx context.Context, projectId string) ([]domain.File, error)
	CreateSignedURL(ctx context.Context, file *domain.File) (string, error)
}

type PhraseHandler interface {
	GetPhrasesById(ctx context.Context, phraseId string) (domain.Phrase, error)
}
