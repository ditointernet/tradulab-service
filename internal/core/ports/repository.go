package ports

import (
	"github.com/ditointernet/tradulab-service/internal/core/domain"
)

type FileRepository interface {
	SaveFile(*domain.File) error
}
