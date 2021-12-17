package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ditointernet/tradulab-service/driven"
	"github.com/ditointernet/tradulab-service/internal/core/domain"
	"github.com/google/uuid"
)

type Phrase struct {
	cli *sql.DB
}

func MustNewPhrase(db *sql.DB) *Phrase {
	return &Phrase{
		cli: db,
	}
}

func (p *Phrase) CreateOrUpdatePhraseTx(ctx context.Context, phrases []*domain.Phrase) error {
	tx, err := p.cli.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	for _, value := range phrases {
		dto := &driven.Phrase{
			Id:      uuid.New().String(),
			FileId:  value.FileId,
			Key:     value.Key,
			Content: value.Content,
		}

		_, err := tx.ExecContext(
			ctx,
			`INSERT into phrases (id, file_id, key, content)
			VALUES ($1, $2, $3, $4)
			ON CONFLICT (key, file_id)
			DO UPDATE SET content = $4`,
			dto.Id,
			dto.FileId,
			dto.Key,
			dto.Content,
		)
		if err != nil {
			tx.Rollback()
			return err
		}

	}
	err = tx.Commit()
	if err != nil {
		return err
	}

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

func (p *Phrase) GetPhrasesById(ctx context.Context, phraseId string) (domain.Phrase, error) {
	var phrase domain.Phrase

	err := p.cli.QueryRowContext(
		ctx,
		"SELECT id, file_id, content, key FROM phrases WHERE id = $1",
		phraseId).Scan(&phrase.Id, &phrase.FileId, &phrase.Content, &phrase.Key)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Phrase{}, errors.New("phrase not found")
		}
		return domain.Phrase{}, err

	}

	return phrase, nil
}

func (p *Phrase) GetFilePhrases(ctx context.Context, fileId string, page int) ([]domain.Phrase, error) {
	if page <= 0 {
		return nil, errors.New("must be bigger zero")
	}

	limit, err := strconv.Atoi(os.Getenv("PAGINATION_LIMIT"))
	// numberPage, err := strconv.Atoi(limit)

	offset := limit * (page - 1)

	var phrases []domain.Phrase

	allPhrases, err := p.cli.QueryContext(ctx, "SELECT id, file_id, key, content FROM phrases WHERE file_id = $1 OFFSET $2 LIMIT $3", fileId, offset, limit)
	if err != nil {
		return nil, err
	}
	defer allPhrases.Close()

	for allPhrases.Next() {
		var phrase domain.Phrase

		err = allPhrases.Scan(&phrase.Id, &phrase.FileId, &phrase.Key, &phrase.Content)
		if err != nil {
			return nil, err
		}

		phrases = append(phrases, phrase)
	}

	return phrases, nil
}
