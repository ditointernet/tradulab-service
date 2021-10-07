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

<<<<<<< HEAD
func (d *File) FindFile(ctx context.Context, id string) error {
	var file domain.File

	err := d.cli.QueryRowContext(
		ctx,
		"SELECT id, project_id, file_path FROM files WHERE id = $1",
		id).Scan(&file.ID, &file.ProjectID, &file.FilePath)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("file not found")
=======
func (d *File) GetFiles(ctx context.Context) ([]domain.File, error) {
	var files []domain.File

	allFiles, err := d.cli.QueryContext(ctx, "SELECT id, project_id, file_path FROM files")
	if err != nil {
		return nil, err
	}
	defer allFiles.Close()

	for allFiles.Next() {
		var file domain.File

		err = allFiles.Scan(&file.ID, &file.ProjectID, &file.FilePath)
		if err != nil {
			return nil, err
>>>>>>> 8d961fd809c092abb5081eb723f4d222b23777b8
		}
		return err

<<<<<<< HEAD
	}

	return nil
}

func (d *File) EditFile(ctx context.Context, file *domain.File) error {
	dto := &driven.File{
		ID:       file.ID,
		FilePath: file.FilePath,
	}

	_, err := d.cli.ExecContext(
		ctx,
		"UPDATE files SET file_path = $2 WHERE id = $1",
		dto.ID,
		dto.FilePath,
	)

	return err
=======
		files = append(files, file)
	}

	return files, nil
>>>>>>> 8d961fd809c092abb5081eb723f4d222b23777b8
}
