package ports

import "github.com/ditointernet/tradulab-service/internal/core/domain"

type FileHandler interface {
	CheckFile(*domain.File) error
	SaveFile(*domain.File) error
}
