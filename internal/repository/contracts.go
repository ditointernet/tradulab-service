package repository

import (
	"context"

	"github.com/ditointernet/tradulab-service/internal/core/domain"
)

type FileRepository interface {
	SaveFile(ctx context.Context, file *domain.File) error
	GetFiles(ctx context.Context) ([]domain.File, error)
}
