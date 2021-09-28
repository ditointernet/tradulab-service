package services

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/ditointernet/tradulab-service/driven"
	"github.com/ditointernet/tradulab-service/internal/core/domain"
	"github.com/ditointernet/tradulab-service/internal/core/ports"
	"github.com/google/uuid"
)

type FileHandler interface {
	CheckFile(*domain.File) error
	SaveFile(*domain.File) error
}

type File struct {
	repo ports.Repository
}

func MustNewFile(repo ports.Repository) *File {
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

	dto := &driven.File{
		ID:        entry.ID,
		ProjectID: entry.ProjectID,
		FilePath:  entry.FilePath,
	}

	err := f.repo.SaveFile(dto)

	if err != nil {
		fmt.Println(err)
	}

	return nil
}
