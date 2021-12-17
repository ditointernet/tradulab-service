package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ditointernet/tradulab-service/driven"
	"github.com/ditointernet/tradulab-service/internal/core/domain"
)

type File struct {
	cli *sql.DB
}

func MustNewFile(db *sql.DB) *File {
	return &File{
		cli: db,
	}
}

func (f *File) CreateFile(ctx context.Context, file domain.File) error {
	dto := &driven.File{
		Id:        file.Id,
		ProjectId: file.ProjectId,
		Status:    driven.CREATED,
	}

	_, err := f.cli.ExecContext(
		ctx,
		"INSERT into files (id, project_id, status) values ($1, $2, $3)",
		dto.Id,
		dto.ProjectId,
		dto.Status,
	)

	return err
}

func (f *File) GetProjectFiles(ctx context.Context, projectId string) ([]domain.File, error) {
	var files []domain.File

	allFiles, err := f.cli.QueryContext(ctx, "SELECT id, project_id, status FROM files WHERE project_id = $1", projectId)
	if err != nil {
		return nil, err
	}
	defer allFiles.Close()

	for allFiles.Next() {
		var file domain.File

		err = allFiles.Scan(&file.Id, &file.ProjectId, &file.Status)
		if err != nil {
			return nil, err
		}

		files = append(files, file)
	}

	return files, nil
}

func (f *File) FindFile(ctx context.Context, id string) (domain.File, error) {
	var file domain.File

	err := f.cli.QueryRowContext(
		ctx,
		"SELECT id, project_id, status FROM files WHERE id = $1",
		id).Scan(&file.Id, &file.ProjectId, &file.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.File{}, errors.New("file not found")
		}
		return domain.File{}, err

	}

	return file, nil
}

func (f *File) SetUploadSuccessful(ctx context.Context, file *domain.File) error {
	dto := &driven.File{
		Id:     file.Id,
		Status: driven.SUCCESS,
	}

	_, err := f.cli.ExecContext(
		ctx,
		"UPDATE files SET status = $2 WHERE id = $1",
		dto.Id,
		dto.Status,
	)

	return err
}

func (f *File) SetStatusFailed(ctx context.Context, file domain.File) error {
	dto := &driven.File{
		Id:        file.Id,
		ProjectId: file.ProjectId,
		Status:    driven.FAILED,
	}

	_, err := f.cli.ExecContext(
		ctx,
		"INSERT into files (id, project_id, status) values ($1, $2, $3)",
		dto.Id,
		dto.ProjectId,
		dto.Status,
	)

	return err
}
