package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

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
		FileID:  phrase.FileID,
		Key:     phrase.Key,
		Content: phrase.Content,
	}

	_, err := p.cli.ExecContext(
		ctx,
		"UPDATE phrases SET content = $2 WHERE key = $1 AND file_id = $3",
		dto.Key,
		dto.Content,
		dto.FileID,
	)

	return err
}

func (p *Phrase) GetByFileId(ctx context.Context, id string) (domain.Phrase, error) {
	var phrase domain.Phrase

	err := p.cli.QueryRowContext(
		ctx,
		"SELECT key FROM phrases WHERE file_id = $1",
		id).Scan(&phrase.Key)
	if err != nil {
		return domain.Phrase{}, err
	}

	return phrase, nil
}

func (p *Phrase) DeletePhrases(ctx context.Context, phrasesKey []string, fileId string) error {
	list := strings.Join(phrasesKey[:], ", ")

	query := fmt.Sprintf("DELETE FROM phrases WHERE file_id = $1 AND key NOT IN (%s)", list)

	_, err := p.cli.ExecContext(
		ctx,
		query,
		fileId,
	)

	if err != nil {
		return err
	}

	return nil
}
