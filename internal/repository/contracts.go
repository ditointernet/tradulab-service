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
