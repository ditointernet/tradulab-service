package repository

import (
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

func (d *File) SaveFile(file *domain.File) error {
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

func (d *File) GetFiles() ([]domain.File, error) {
	var files []domain.File

	allFiles, err := d.cli.Query(
		"SELECT * FROM files",
	)
	if err != nil {
		return nil, err
	}

	for allFiles.Next() {
		var id, project_id, file_path string

		err = allFiles.Scan(&id, &project_id, &file_path)
		if err != nil {
			return nil, err
		}

		f := domain.File{
			ID:        id,
			ProjectID: project_id,
			FilePath:  file_path,
		}

		files = append(files, f)
	}

	defer allFiles.Close()

	return files, nil
}
