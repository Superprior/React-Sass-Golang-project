package store

import "backend/internal/app/model"

type LanguageRepository struct {
	store *Store
}

func (r *LanguageRepository) Create(l *model.Language) (*model.Language, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO programming_languages (slug, title) VALUES ($1, $2) RETURNING id;",
		l.Slug, l.Title,
	).Scan(&l.ID); err != nil {
		return nil, err
	}
	return l, nil
}

func (r *LanguageRepository) FindBySlug(slug string) (*model.Language, error) {
	l := &model.Language{}
	if err := r.store.db.QueryRow(
		"SELECT id, slug, title FROM programming_languages WHERE slug=$1",
		slug,
	).Scan(
		&l.ID,
		&l.Slug,
		&l.Title,
	); err != nil {
		return nil, err
	}
	return l, nil
}

func (r *LanguageRepository) GetList() (*[]model.Language, error) {
	rows, err := r.store.db.Query("SELECT * FROM programming_languages;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var langs []model.Language

	for rows.Next() {
		var lang model.Language
		if err := rows.Scan(&lang.ID, &lang.Slug, &lang.Title); err != nil {
			return &langs, nil
		}
		langs = append(langs, lang)
	}

	if err = rows.Err(); err != nil {
		return &langs, err
	}
	return &langs, nil
}
