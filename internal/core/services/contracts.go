package services

import (
	"context"

	"github.com/ditointernet/tradulab-service/internal/core/domain"
)

type FileHandler interface {
	SaveFile(ctx context.Context, file *domain.File) error
	EditFile(ctx context.Context, file *domain.File) error
}
