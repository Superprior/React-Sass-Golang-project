package store

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Store struct {
	db *sql.DB
}

// func (s *Store) Open() error {
// 	dbConnectionString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
// 		s.config.DatabaseHost, s.config.DatabasePort, s.config.DatabaseName,
// 		s.config.DatabaseUser, s.config.DatabasePassword, s.config.DatabaseSSLMode)

// 	db, err := sql.Open("postgres", dbConnectionString)
// 	if err != nil {
// 		return err
// 	}

// 	if err := db.Ping(); err != nil {
// 		return err
// 	}

// 	s.db = db
// 	return nil
// }

func (s *Store) Close() {
	s.db.Close()
}
