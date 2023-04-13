package store

import (
	"backend/internal/app/model"
	"backend/pkg/parsers"
	"fmt"
)

type SampleRepository struct {
	store *Store
}

func (r *SampleRepository) GetList() (*[]model.Sample, error) {
	rows, err := r.store.db.Query("SELECT id, title, lang_slug FROM code_samples;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var samples []model.Sample

	for rows.Next() {
		var sample model.Sample
		var codeLinesStr string
		if err := rows.Scan(&sample.ID, &sample.Title, &sample.LangSlug); err != nil {
			return &samples, nil
		}

		err := r.store.db.QueryRow(
			fmt.Sprintf("SELECT array_to_string(content, '\\\\') AS lines FROM code_samples WHERE id=%d;", sample.ID),
		).Scan(&codeLinesStr)

		sample.Content = parsers.ParsePostgresArray(codeLinesStr)

		if err != nil {
			return nil, err
		}
		samples = append(samples, sample)
	}

	if err = rows.Err(); err != nil {
		return &samples, err
	}
	return &samples, nil
}

func (r *SampleRepository) GetInstance(id int) (*model.Sample, error) {
	var (
		sample       model.Sample
		codeLinesStr string
	)

	err := r.store.db.QueryRow(fmt.Sprintf("SELECT id, title, lang_slug FROM code_samples WHERE id=%d;", id)).Scan(
		&sample.ID, &sample.Title, &sample.LangSlug)

	if err != nil {
		return nil, err
	}

	err = r.store.db.QueryRow(
		fmt.Sprintf("SELECT array_to_string(content, '\\\\') AS lines FROM code_samples WHERE id=%d;", sample.ID),
	).Scan(&codeLinesStr)

	sample.Content = parsers.ParsePostgresArray(codeLinesStr)

	if err != nil {
		return nil, err
	}
	return &sample, nil
}

func (r *SampleRepository) CreateInstance(title string, langSlug string, context string) (int, error) {
	codeLines := parsers.RawStringToLinesArray(context)
	arrayToInsert := parsers.LinesArrayToPostgresArray(codeLines)
	var id int

	query := fmt.Sprintf(
		"INSERT INTO code_samples (title, content, lang_slug) VALUES ('%s', ARRAY %s, '%s') RETURNING id;",
		title, arrayToInsert, langSlug,
	)

	err := r.store.db.QueryRow(query).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *SampleRepository) DeleteInstance(id int) error {
	err := r.store.db.QueryRow(fmt.Sprintf("DELETE FROM code_samples WHERE id=%d;", id))

	if err != nil {
		println(err)
	}
	return nil
}
