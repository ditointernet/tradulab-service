package services

import (
	"github.com/ditointernet/tradulab-service/drivers"
	"github.com/google/uuid"
)

func (file Path) CheckFile() bool {
	dto := drivers.File{}

	extenssion := file[len(file)-4:]

	if extenssion != ".csv" {
		return false
	}

	id := uuid.New().String()

	return true
}

type File interface {
	CheckFile() bool
}

type Path struct {
	P string
}
