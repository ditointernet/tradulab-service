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

func (d *File) SaveFile(ctx context.Context, file *domain.File) error {
	dto := &driven.File{
		ID:        file.ID,
		ProjectID: file.ProjectID,
		FilePath:  file.FilePath,
	}

	_, err := d.cli.ExecContext(
		ctx,
		"INSERT into files (id, project_id, file_path) values ($1, $2, $3)",
		dto.ID,
		dto.ProjectID,
		dto.FilePath,
	)

	return err
}

func (d *File) FindFile(ctx context.Context, id string) error {
	var file domain.File

	err := d.cli.QueryRowContext(
		ctx,
		"SELECT id, project_id, file_path FROM files WHERE id = $1",
		id).Scan(&file.ID, &file.ProjectID, &file.FilePath)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("file not found")
		}
		return err

	}

	return nil
}

func (d *File) EditFile(ctx context.Context, file *domain.File) error {
	dto := &driven.File{
		ID:        file.ID,
		ProjectID: file.ProjectID,
		FilePath:  file.FilePath,
	}

	_, err := d.cli.ExecContext(
		ctx,
		"UPDATE files SET project_id = $2, file_path = $3 WHERE id = $1",
		dto.ID,
		dto.ProjectID,
		dto.FilePath,
	)

	return err
}
