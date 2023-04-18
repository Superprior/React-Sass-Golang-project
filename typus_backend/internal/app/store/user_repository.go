package store

import (
	"backend/internal/app/model"
	"backend/pkg/jwt_funcs"
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

func (r *UserRepository) Login(user *model.User) (string, error) {

	/*
		Find the requested user in db, check his/her credentials, make a JWT if everything's fine.
		If the password is incorrect or user doesn't exist, return error.
	*/

	// Check if username or email was provided, query the database

	var (
		query            string
		err              error
		encrypted_pwd_db string
	)

	if user.Username == "" {
		query = fmt.Sprintf("SELECT id, username, role, encrypted_pwd FROM users WHERE email='%s';", user.Email)
		err = r.store.db.QueryRow(query).Scan(&user.ID, &user.Username, &user.Role, &encrypted_pwd_db)
	} else {
		query = fmt.Sprintf("SELECT id, email, role, encrypted_pwd FROM users WHERE username='%s';", user.Username)
		err = r.store.db.QueryRow(query).Scan(&user.ID, &user.Email, &user.Role, &encrypted_pwd_db)
	}

	if err != nil {
		return "", err
	}

	// Check if the password is correct

	if user.EncryptedPwd != encrypted_pwd_db {
		return "", fmt.Errorf("Incorrect password provided")
	}

	// Make and return a JWT

	new_token, err := jwt_funcs.GenerateJWT(user.Username, user.Email, int8(user.Role))

	if err != nil {
		return "", err
	}

	return new_token, nil
}
