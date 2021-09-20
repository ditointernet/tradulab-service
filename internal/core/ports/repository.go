package ports

import "github.com/ditointernet/tradulab-service/drivers"

type Repository interface {
	SaveFile(*drivers.File) error
}
