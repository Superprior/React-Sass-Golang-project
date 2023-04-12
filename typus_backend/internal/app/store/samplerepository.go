package store

import (
	"backend/internal/app/model"
	"backend/pkg/parsers"
	"fmt"
	"log"
)

type SampleRepository struct {
	store *Store
}

func (r *SampleRepository) FindById(id int) (*model.Sample, error) {
	log.Fatal("Not implemented")
	return nil, nil
}

func (r *SampleRepository) GetList() (*[]model.Sample, error) {
	rows, err := r.store.db.Query("SELECT * FROM code_samples;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var samples []model.Sample

	for rows.Next() {
		var sample model.Sample
		var codeLinesStr string
		if err := rows.Scan(&sample.ID, &sample.Title, &codeLinesStr, &sample.LangSlug); err != nil {
			return &samples, nil
		}

		sample.Content = parsers.ParsePostgresArray(codeLinesStr)
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

	err := r.store.db.QueryRow(fmt.Sprintf("SELECT * FROM code_samples WHERE id=%d;", id)).Scan(
		&sample.ID, &sample.Title, &codeLinesStr, &sample.LangSlug)

	sample.Content = parsers.ParsePostgresArray(codeLinesStr)

	if err != nil {
		return nil, err
	}
	return &sample, nil
}
