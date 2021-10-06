package repository

import (
	"context"
	"database/sql"

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

	_, err := d.cli.Exec(
		"INSERT into files (id, project_id, file_path) values ($1, $2, $3)",
		dto.ID,
		dto.ProjectID,
		dto.FilePath,
	)

	return err
}

func (d *File) GetFiles(ctx context.Context) ([]domain.File, error) {
	var files []domain.File

	allFiles, err := d.cli.Query(
		"SELECT id, project_id, file_path FROM files",
	)
	if err != nil {
		return nil, err
	}
	defer allFiles.Close()

	for allFiles.Next() {
		var file domain.File

		err = allFiles.Scan(&file.ID, &file.ProjectID, &file.FilePath)
		if err != nil {
			return nil, err
		}

		files = append(files, file)
	}

	return files, nil
}
