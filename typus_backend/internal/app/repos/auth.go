package repos

import (
	"backend/internal/app/models"
	"backend/internal/app/store"
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
