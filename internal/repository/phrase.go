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

func (p *Phrase) CreateOrUpdatePhrase(ctx context.Context, phrase domain.Phrase) error {
	dto := &driven.Phrase{
		ID:      phrase.ID,
		FileID:  phrase.FileID,
		Key:     phrase.Key,
		Content: phrase.Content,
	}

	_, err := p.cli.ExecContext(
		ctx,
		`INSERT into phrases (id, file_id, key, content)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (key, file_id)
		DO UPDATE SET content = $4`,
		dto.ID,
		dto.FileID,
		dto.Key,
		dto.Content,
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
	var phrasesFormat []string

	for _, valueKey := range phrasesKey {
		phrasesFormat = append(phrasesFormat, fmt.Sprintf("'%s'", valueKey))
	}

	list := strings.Join(phrasesFormat[:], ", ")

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
