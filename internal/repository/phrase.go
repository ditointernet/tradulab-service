package repository

import (
	"context"
	"database/sql"

	"github.com/ditointernet/tradulab-service/driven"
	"github.com/ditointernet/tradulab-service/internal/core/domain"
)

type Phrase struct {
	cli *sql.DB
}

func MustNewPhrase(db *sql.DB) *Phrase {
	return &Phrase{
		cli: db,
	}
}

func (d *Phrase) CreatePhrase(ctx context.Context, phrase domain.Phrase) error {
	dto := &driven.Phrase{
		FileID:  phrase.FileID,
		Key:     phrase.Key,
		Content: phrase.Content,
	}

	_, err := d.cli.ExecContext(
		ctx,
		"INSERT into phrases (file_id, key, content) values ($1, $2, $3)",
		dto.FileID,
		dto.Key,
		dto.Content,
	)

	return err
}
