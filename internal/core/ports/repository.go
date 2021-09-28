package ports

import "github.com/ditointernet/tradulab-service/driven"

type Repository interface {
	SaveFile(*driven.File) error
}
