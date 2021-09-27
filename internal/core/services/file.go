package services

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/ditointernet/tradulab-service/database"
	"github.com/ditointernet/tradulab-service/internal/core/domain"
	"github.com/google/uuid"
)

type FileHandler interface {
	CheckFile(*domain.File) error
	SaveFile(*domain.File) error
}

type File struct {
	repo database.Database
}

func MustNewFile(repo database.Database) *File {
	return &File{repo: repo}
}

func (f File) CheckFile(entry *domain.File) error {

	extension := filepath.Ext(entry.FilePath)
	if extension != ".csv" {
		return errors.New("file not supported. Must be .csv")
	}

	return nil
}

<<<<<<< HEAD
// corrigir pra camada de domain
func (f *File) SaveFile(entry *domain.File) error {
	content := &domain.File{
		ID: uuid.New().String(),
=======
func (f *File) SaveFile(entry *domain.File) error {

	content := &domain.File{
		ID:        uuid.New().String(),
		ProjectID: entry.ProjectID,
		FilePath:  entry.FilePath,
>>>>>>> 0ceae2d6338a43dbcf7d20f0a41503cd42613423
	}

	err := f.repo.SaveFile(content)

	if err != nil {
		fmt.Println(err)
	}

	return nil
}
