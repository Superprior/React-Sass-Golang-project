package store

import (
	"backend/internal/app/model"
	"database/sql"
	"fmt"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) CreateInstance(user *model.User) (int, error) {
	query := fmt.Sprintf(
		"INSERT INTO users (username, email, role, encrypted_pwd) VALUES ('%s', '%s', '%d', '%s') RETURNING id;",
		user.Username, user.Email, user.Role, user.EncryptedPwd,
	)

	var id int
	err := r.store.db.QueryRow(query).Scan(&id)

	if err != nil {
		return 0, nil
	}

	return id, nil
}

func (r *UserRepository) GetInstance(id int) (*model.User, error) {

	query := fmt.Sprintf(
		"SELECT id, username, email, role FROM users WHERE id=%d;",
		id,
	)

	var user model.User
	err := r.store.db.QueryRow(query).Scan(&user.ID, &user.Username, &user.Email, &user.Role)

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) CheckUniqueValue(name string, value string) (bool, error) {
	query := fmt.Sprintf(
		"SELECT id FROM users WHERE %s='%s';",
		name, value,
	)
	var id int
	err := r.store.db.QueryRow(query).Scan(&id)
	if err != nil {
		if err != sql.ErrNoRows {
			return false, err
		} else {
			return true, nil
		}
	}
	return false, nil
}
