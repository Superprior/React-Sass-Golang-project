package repos

import (
	"backend/internal/app/models"
	"backend/internal/app/store"
	"database/sql"
	"fmt"
)

type AuthRepo struct {
	store *store.Store
}

func NewAuthRepo() (*AuthRepo, error) {
	sConf := store.NewConfig()
	s := store.New(sConf)

	if err := s.Open(); err != nil {
		return nil, err
	}

	return &AuthRepo{
		store: s,
	}, nil
}

func (r *AuthRepo) CreateUser(user *models.User) (*models.User, error) {

	query := fmt.Sprintf(
		"INSERT INTO users (username, email, role, encrypted_pwd) VALUES ('%s', '%s', %d, '%s') RETURNING id;",
		user.Username, user.Email, user.Role, user.EncryptedPwd,
	)

	err := r.store.DB.QueryRow(query).Scan(&user.ID)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *AuthRepo) GetUserByEmail(email string) (*models.User, error) {

	query := fmt.Sprintf("SELECT id, username, email, role, encrypted_pwd FROM users WHERE email='%s'", email)
	var user models.User

	err := r.store.DB.QueryRow(query).Scan(&user.ID, &user.Username, &user.Email, &user.Role, &user.EncryptedPwd)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("No user with such email")
		}
		return nil, err
	}
	return &user, nil
}

func (r *AuthRepo) GetUserByUsername(username string) (*models.User, error) {

	query := fmt.Sprintf("SELECT id, username, email, role, encrypted_pwd FROM users WHERE username='%s'", username)
	var user models.User

	err := r.store.DB.QueryRow(query).Scan(&user.ID, &user.Username, &user.Email, &user.Role, &user.EncryptedPwd)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("No user with such username")
		}
		return nil, err
	}
	return &user, nil
}
