package services

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/pkg/errors"

	"github.com/ditointernet/tradulab-service/internal/core/domain"
	"github.com/ditointernet/tradulab-service/internal/repository"
	"github.com/ditointernet/tradulab-service/internal/storage"
	"github.com/google/uuid"
)

type File struct {
	repo    repository.FileRepository
	storage storage.FileStorage
}

func MustNewFile(repo repository.FileRepository, storage storage.FileStorage) *File {
	return &File{
		repo:    repo,
		storage: storage,
	}
}

func (f File) CheckExtension(extension string) error {
	if extension != ".json" {
		return errors.New("file not supported. Must be .json")
	}

	return nil
}

func (f *File) CreateFile(ctx context.Context, entry *domain.File) (domain.File, error) {
	extension := filepath.Ext(entry.FileName)
	err := f.CheckExtension(extension)
	if err != nil {
		newFile := domain.File{
			Id:        uuid.New().String(),
			ProjectId: entry.ProjectId,
		}
		f.repo.SetStatusFailed(ctx, newFile)

		return domain.File{}, errors.New("error trying to create the file, status changed to failed")
	}

	id := uuid.New().String()
	fileName := fmt.Sprintf("%s%s", id, extension)
	url, err := f.storage.CreateSignedURL(ctx, fileName)
	if err != nil {
		return domain.File{}, errors.Wrap(err, "couldn't create SignedURL")
	}

	newFile := domain.File{
		Id:        id,
		ProjectId: entry.ProjectId,
		FilePath:  url,
		FileName:  fileName,
	}
	err = f.repo.CreateFile(ctx, newFile)
	if err != nil {
		return domain.File{}, err
	}

	return newFile, nil
}

func (f File) findFile(ctx context.Context, id string) (domain.File, error) {
	file, err := f.repo.FindFile(ctx, id)
	if err != nil {
		return domain.File{}, err
	}

	return file, nil
}

func (f *File) GetProjectFiles(ctx context.Context, projectId string) ([]domain.File, error) {
	files, err := f.repo.GetProjectFiles(ctx, projectId)
	if err != nil {
		return nil, err
	}

	return files, nil
}

func (f *File) SetUploadSuccessful(ctx context.Context, entry *domain.File) error {
	_, err := f.findFile(ctx, entry.Id)
	if err != nil {
		return err
	}

	err = f.repo.SetUploadSuccessful(ctx, entry)
	if err != nil {
		return err
	}

	return nil
}

func (f *File) CreateSignedURL(ctx context.Context, entry *domain.File) (string, error) {
	file, err := f.findFile(ctx, entry.Id)
	if err != nil {
		return "", err
	}
	extension := filepath.Ext(entry.FileName)
	err = f.CheckExtension(extension)
	if err != nil {
		return "", err
	}

	fileName := fmt.Sprintf("%s%s", file.Id, extension)

	url, err := f.storage.CreateSignedURL(ctx, fileName)
	if err != nil {
		return "", errors.Wrap(err, "couldn't create SignedURL")
	}

	return url, nil
}
