package store

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Standard store structure.
type Store struct {
	DB     *sql.DB
	Config *Config
}

// Create a new store instance.
func New(config *Config) *Store {
	return &Store{
		Config: config,
	}
}

// Open a new store connection with parameters provided in its Config field.
func (s *Store) Open() (err error) {
	dbConnectionString := fmt.Sprintf("postgres:///%s?host=%s&sslmode=%s&user=%s&password=%s",
		s.Config.Name, s.Config.Host, s.Config.SSLMode, s.Config.User, s.Config.Password)

	db, err := sql.Open("postgres", dbConnectionString)
	if err != nil {
		return
	}

	if err = db.Ping(); err != nil {
		return
	}

	s.DB = db
	return
}

// Close the connection.
func (s *Store) Close() {
	s.DB.Close()
}
