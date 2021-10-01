package repository

import (
	"github.com/ditointernet/tradulab-service/internal/core/domain"
)

type FileRepository interface {
	SaveFile(file *domain.File) error
	GetFiles() ([]domain.File, error)
}
