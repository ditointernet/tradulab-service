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

func (d *File) CreateFile(ctx context.Context, file domain.File) error {
	dto := &driven.File{
		ID:        file.ID,
		ProjectID: file.ProjectID,
		Status:    driven.CREATED,
	}

	_, err := d.cli.ExecContext(
		ctx,
		"INSERT into files (id, project_id, status) values ($1, $2, $3)",
		dto.ID,
		dto.ProjectID,
		dto.Status,
	)

	return err
}

func (d *File) GetFiles(ctx context.Context) ([]domain.File, error) {
	var files []domain.File

	allFiles, err := d.cli.QueryContext(ctx, "SELECT id, project_id, status FROM files") // tem que arrrumar esse filePath
	if err != nil {
		return nil, err
	}
	defer allFiles.Close()

	for allFiles.Next() {
		var file domain.File

		err = allFiles.Scan(&file.ID, &file.ProjectID, &file.Status)
		if err != nil {
			return nil, err
		}

		files = append(files, file)
	}

	return files, nil
}

func (d *File) FindFile(ctx context.Context, id string) (domain.File, error) {
	var file domain.File

	err := d.cli.QueryRowContext(
		ctx,
		"SELECT id, project_id, status FROM files WHERE id = $1",
		id).Scan(&file.ID, &file.ProjectID, &file.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.File{}, errors.New("file not found")
		}
		return domain.File{}, err

	}

	return file, nil
}

func (d *File) SetUploadSuccessful(ctx context.Context, file *domain.File) error {
	dto := &driven.File{
		ID:     file.ID,
		Status: driven.SUCCESS,
	}

	_, err := d.cli.ExecContext(
		ctx,
		"UPDATE files SET status = $2 WHERE id = $1",
		dto.ID,
		dto.Status,
	)

	return err
}
