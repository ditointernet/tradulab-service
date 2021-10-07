package services

import (
	"context"

	"github.com/ditointernet/tradulab-service/internal/core/domain"
)

type FileHandler interface {
	CreateFile(ctx context.Context, file *domain.File) error
	EditFile(ctx context.Context, file *domain.File) error
	GetFiles(ctx context.Context) ([]domain.File, error)
}
