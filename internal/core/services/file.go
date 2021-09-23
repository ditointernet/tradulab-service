package services

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/ditointernet/tradulab-service/database"
	"github.com/ditointernet/tradulab-service/drivers"
	"github.com/ditointernet/tradulab-service/internal/core/domain"
	"github.com/google/uuid"
)

type FileHandler interface {
	CheckFile(*domain.File) error
	SaveFile(*drivers.File) error
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

// corrigir pra camada de domain
func (f *File) SaveFile(*drivers.File) error {
	content := &drivers.File{}
	content.ID = uuid.New().String()

	err := f.repo.SaveFile(content)

	if err != nil {
		fmt.Println(err)
	}

	return nil
}
