package repository

import (
	"context"

	"github.com/ditointernet/tradulab-service/internal/core/domain"
)

type FileRepository interface {
	SaveFile(ctx context.Context, file *domain.File) error
	FindFile(ctx context.Context, id string) error
	EditFile(ctx context.Context, file *domain.File) error
}
