package repos

import (
	"backend/internal/app/models"
	"backend/internal/app/store"
	"database/sql"
	"fmt"
)

// Auth repository.
// Contains the store in order to hit the database using provided connection.
type AuthRepo struct {
	store *store.Store
}

// NewAuthRepo
// This function creates a new AuthRepo.
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

// CreateUser
// This function creates a new user in the database using provided data.
func (r *AuthRepo) CreateUser(user *models.User) (*models.User, error) {

	// Create the query
	query := fmt.Sprintf(
		"INSERT INTO users (username, email, role, encrypted_pwd) VALUES ('%s', '%s', %d, '%s') RETURNING id;",
		user.Username, user.Email, user.Role, user.EncryptedPwd,
	)

	// Perform the query and get the id
	err := r.store.DB.QueryRow(query).Scan(&user.ID)

	if err != nil {
		return nil, err
	}

	// Return the same user but with the id
	return user, nil
}

// GetUserByEmail
// This function get the user from the database by the provided email.
func (r *AuthRepo) GetUserByEmail(email string) (*models.User, error) {

	// Create the query
	query := fmt.Sprintf("SELECT id, username, email, role, encrypted_pwd FROM users WHERE email='%s'", email)
	var user models.User

	// Perform the query and get the user's data
	err := r.store.DB.QueryRow(query).Scan(&user.ID, &user.Username, &user.Email, &user.Role, &user.EncryptedPwd)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("No user with such email")
		}
		return nil, err
	}
	return &user, nil
}

// GetUserByEmail
// This function get the user from the database by the provided email.
func (r *AuthRepo) GetUserByUsername(username string) (*models.User, error) {

	// Create the query
	query := fmt.Sprintf("SELECT id, username, email, role, encrypted_pwd FROM users WHERE username='%s'", username)
	var user models.User

	// Perform the query and get the user's data
	err := r.store.DB.QueryRow(query).Scan(&user.ID, &user.Username, &user.Email, &user.Role, &user.EncryptedPwd)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("No user with such username")
		}
		return nil, err
	}
	return &user, nil
}
