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

func (p *Phrase) GetPhrase(ctx context.Context, entry domain.Phrase) (domain.Phrase, error) {
	var phrase domain.Phrase

	err := p.cli.QueryRowContext(
		ctx,
		"SELECT key FROM phrases WHERE key = $1 and file_id = $2",
		entry.Key, entry.FileID).Scan(&phrase.Key)
	if err != nil {
		return domain.Phrase{}, err
	}

	return phrase, nil
}

func (p *Phrase) CreatePhrase(ctx context.Context, phrase domain.Phrase) error {
	dto := &driven.Phrase{
		ID:      phrase.ID,
		FileID:  phrase.FileID,
		Key:     phrase.Key,
		Content: phrase.Content,
	}

	_, err := p.cli.ExecContext(
		ctx,
		"INSERT into phrases (id, file_id, key, content) values ($1, $2, $3, $4)",
		dto.ID,
		dto.FileID,
		dto.Key,
		dto.Content,
	)

	return err
}

func (p *Phrase) UpdatePhrase(ctx context.Context, phrase domain.Phrase) error {
	dto := &driven.Phrase{
		Key:     phrase.Key,
		Content: phrase.Content,
	}

	_, err := p.cli.ExecContext(
		ctx,
		"UPDATE phrases SET content = $2 WHERE key = $1",
		dto.Key,
		dto.Content,
	)

	return err
}
