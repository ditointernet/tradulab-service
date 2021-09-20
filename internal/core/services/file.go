package services

import (
	"errors"
	"path/filepath"

	"github.com/ditointernet/tradulab-service/database"
	"github.com/ditointernet/tradulab-service/drivers"
	"github.com/google/uuid"
)

type FileHandler interface {
	CheckFile(*drivers.File) error
	SaveFile(*drivers.File) error
}
type File struct {
	repo database.Database
}

func MustNewFile() *File {
	return &File{}
}

func (f File) CheckFile(entry *drivers.File) error {

	extension := filepath.Ext(entry.FilePath)
	if extension != ".csv" {
		return errors.New("file not supported. Must be .csv")
	}

	return nil
}

func (f *File) SaveFile(*drivers.File) error {
	content := &drivers.File{}
	content.ID = uuid.New().String()

	f.repo.SaveFile(content)

	return nil
}
