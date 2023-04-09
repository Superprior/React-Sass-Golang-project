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
