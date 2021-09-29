package ports

import (
	"github.com/ditointernet/tradulab-service/internal/core/domain"
)

type Repository interface {
	SaveFile(*domain.File) error
}
