package store

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Store struct {
	config             *Config
	db                 *sql.DB
	languageRepository *LanguageRepository
	sampleRepository   *SampleRepository
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	dbConnectionString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		s.config.DatabaseHost, s.config.DatabasePort, s.config.DatabaseName,
		s.config.DatabaseUser, s.config.DatabasePassword, s.config.DatabaseSSLMode)

	db, err := sql.Open("postgres", dbConnectionString)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db
	return nil
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) Language() *LanguageRepository {
	if s.languageRepository != nil {
		return s.languageRepository
	}

	s.languageRepository = &LanguageRepository{
		store: s,
	}

	return s.languageRepository
}

func (s *Store) Sample() *SampleRepository {
	if s.sampleRepository != nil {
		return s.sampleRepository
	}

	s.sampleRepository = &SampleRepository{
		store: s,
	}

	return s.sampleRepository
}
