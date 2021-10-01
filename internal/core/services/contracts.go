package services

import "github.com/ditointernet/tradulab-service/internal/core/domain"

type FileHandler interface {
	SaveFile(*domain.File) error
	GetFiles() ([]*domain.File, error)
}
