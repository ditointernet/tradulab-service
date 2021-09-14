package services

import (
	"errors"

	"github.com/ditointernet/tradulab-service/drivers"
	"github.com/google/uuid"
)

func (file Path) CheckFile() (drivers.File, error) {

	extenssion := file.P[len(file.P)-4:]

	if extenssion != ".csv" {
		return drivers.File{}, errors.New("file not supported. Must be .csv")
	}

	id := uuid.New().String()
	projectId := uuid.New().String()
	dto := drivers.File{ID: id, ProjectID: projectId, FilePath: file.P}

	return dto, nil
}

type File interface {
	CheckFile() (drivers.File, error)
}

type Path struct {
	P string
}
