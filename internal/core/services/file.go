package services

import (
	"errors"
	"path/filepath"

	"github.com/ditointernet/tradulab-service/internal/core/domain"
)

type FileHandler interface {
	CheckFile(*domain.File) error
}
type File struct {
}

func MustNewFile() *File {
	return &File{}
}

func (f File) CheckFile(entry *domain.File) error {

	extension := filepath.Ext(entry.FilePath)
	if extension != ".csv" {
		return errors.New("file not supported. Must be .csv")
	}

	return nil
}
