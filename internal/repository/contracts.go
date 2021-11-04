package repository

import (
	"context"

	"github.com/ditointernet/tradulab-service/internal/core/domain"
)

type FileRepository interface {
	CreateFile(ctx context.Context, file domain.File) error
	FindFile(ctx context.Context, id string) (domain.File, error)
	UpdateFileStatus(ctx context.Context, file *domain.File) error
	GetFiles(ctx context.Context) ([]domain.File, error)
}
