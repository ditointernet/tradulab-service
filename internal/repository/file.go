package repository

import (
	"github.com/ditointernet/tradulab-service/adapters"
	"github.com/ditointernet/tradulab-service/driven"
	"github.com/ditointernet/tradulab-service/internal/core/domain"
)

type File struct {
	cli *adapters.Database
}

func MustNewFile(db *adapters.Database) *File {
	return &File{
		cli: db,
	}
}

func (d *File) SaveFile(file *domain.File) error {

	client := d.cli.GetDatabase()

	dto := &driven.File{
		ID:        file.ID,
		ProjectID: file.ProjectID,
		FilePath:  file.FilePath,
	}

	query := client.Exec(
		"INSERT into files (id, project_id, file_path) values (?,?,?)",
		dto.ID,
		dto.ProjectID,
		dto.FilePath,
	)

	return query.Error
}
