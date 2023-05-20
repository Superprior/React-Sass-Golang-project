package store

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Store struct {
	DB     *sql.DB
	Config *Config
}

func New(config *Config) *Store {
	return &Store{
		Config: config,
	}
}

func (s *Store) Open() error {
	dbConnectionString := fmt.Sprintf("postgres:///%s?host=%s&sslmode=%s&user=%s&password=%s",
		s.Config.Name, s.Config.Host, s.Config.SSLMode,
		s.Config.User, s.Config.Password)

	db, err := sql.Open("postgres", dbConnectionString)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.DB = db
	return nil
}

func (s *Store) Close() {
	s.DB.Close()
}
