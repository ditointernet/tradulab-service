package services

import (
	"errors"
	"path/filepath"

	"github.com/ditointernet/tradulab-service/internal/core/domain"
	"github.com/ditointernet/tradulab-service/internal/core/ports"
	"github.com/google/uuid"
)

type File struct {
	repo ports.FileRepository
}

func MustNewFile(repo ports.FileRepository) *File {
	return &File{repo: repo}
}

func (f File) CheckFile(entry *domain.File) error {

	extension := filepath.Ext(entry.FilePath)
	if extension != ".csv" {
		return errors.New("file not supported. Must be .csv")
	}

	return nil
}

func (f *File) SaveFile(entry *domain.File) error {

	entry.ID = uuid.New().String()

	err := f.repo.SaveFile(entry)

	if err != nil {
		return errors.New("error in database")
	}

	return nil
}
