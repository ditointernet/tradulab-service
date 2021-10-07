package services

import (
	"context"
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

func (f *File) CreateFile(ctx context.Context, entry *domain.File) error {
	err := f.CheckFile(entry)
	if err != nil {
		return err
	}

	entry.ID = uuid.New().String()

	err = f.repo.CreateFile(ctx, entry)
	if err != nil {
		return err
	}

	return nil
}

func (f File) findFile(ctx context.Context, id string) error {
	err := f.repo.FindFile(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (f *File) GetFiles(ctx context.Context) ([]domain.File, error) {
	files, err := f.repo.GetFiles(ctx)
	if err != nil {
		return nil, err
	}

	return files, nil
}

func (f *File) EditFile(ctx context.Context, entry *domain.File) error {
	err := f.findFile(ctx, entry.ID)
	if err != nil {
		return err
	}

	err = f.repo.EditFile(ctx, entry)
	if err != nil {
		return err
	}

	return nil
}
