package services

import (
	"errors"
	"path/filepath"

	"github.com/ditointernet/tradulab-service/internal/core/domain"
	"github.com/ditointernet/tradulab-service/internal/repository"
	"github.com/google/uuid"
)

type File struct {
	repo repository.FileRepository
}

func MustNewFile(repo repository.FileRepository) *File {
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
	err := f.CheckFile(entry)

	if err != nil {
		return err
	}

	entry.ID = uuid.New().String()

	err = f.repo.SaveFile(entry)
	if err != nil {
		return err
	}

	return nil
}

func (f File) findFile(id string) error {
	err := f.repo.FindFile(id)
	if err != nil {
		return err
	}

	return nil
}

func (f *File) EditFile(entry *domain.File) error {
	err := f.findFile(entry.ID)
	if err != nil {
		return err
	}

	err = f.repo.EditFile(entry)
	if err != nil {
		return err
	}

	return nil
}
