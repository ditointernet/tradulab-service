package services

import (
	"errors"
	"path/filepath"

	"github.com/ditointernet/tradulab-service/drivers"
)

type FileHandler interface {
	CheckFile(*drivers.File) error
}
type File struct {
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
