package repos

import (
	"backend/internal/errors"
	"backend/internal/models"
	"backend/pkg/store"
	"database/sql"
	"log"

	"github.com/lib/pq"
)

// Default representation of the auth repository.
// Contains the store in order to query the database.
type AuthRepo struct {
	store *store.Store
}

// Create a new AuthRepo.
func NewAuthRepo() (repo *AuthRepo, err error) {
	sConf := store.NewConfig()
	s := store.New(sConf)

	if err = s.Open(); err != nil {
		log.Fatal("Could not create AuthRepo")
	}

	repo = &AuthRepo{store: s}
	return
}

// Create a new user in the database using provided data.
func (r *AuthRepo) CreateInstance(user *models.User) (*models.User, error) {

	// Construct the query and query the database
	query := `
		INSERT INTO users (username, email, role, encrypted_pwd) 
		VALUES ($1, $2, $3, $4) 
		RETURNING id;`

	err := r.store.DB.
		QueryRow(query, user.Username, user.Email, user.Role, user.EncryptedPwd).
		Scan(&user.ID)

	// Check for insertion errors
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				if pqErr.Constraint == "users_username_key" {
					err = errors.ErrNonUniqueEmail
					return nil, err
				} else if pqErr.Constraint == "users_email_key" {
					err = errors.ErrNonUniqueEmail
					return nil, err
				}
			default:
				{
					err = errors.ErrServerError
					return nil, err
				}
			}
		}
	}

	// Return the same user but with the id
	return user, nil
}

// Get the user from the database using provided email.
func (r *AuthRepo) GetInstanceByEmail(email string) (user *models.User, err error) {

	// Construct the query and query the database
	query := `
		SELECT id, username, email, role, encrypted_pwd 
		FROM users 
		WHERE email=$1;`

	// Perform the query and get the user's data
	err = r.store.DB.
		QueryRow(query, email).
		Scan(&user.ID, &user.Username, &user.Email, &user.Role, &user.EncryptedPwd)

	// Check for errors and return if everything's fine
	if err != nil {
		if err == sql.ErrNoRows {
			err = errors.ErrNoUserWithEmail
			return
		}
		err = errors.ErrServerError
		return
	}
	return
}

// GetUserByEmail
// This function get the user from the database by the provided email.
func (r *AuthRepo) GetInstanceByUsername(username string) (user *models.User, err error) {

	// Construct the query and query the database
	query := `
		SELECT id, username, email, role, encrypted_pwd 
		FROM users 
		WHERE username=$1`

	// Perform the query and get the user's data
	err = r.store.DB.
		QueryRow(query, username).
		Scan(&user.ID, &user.Username, &user.Email, &user.Role, &user.EncryptedPwd)

	// Check for errors and return if everything's fine
	if err != nil {
		if err == sql.ErrNoRows {
			err = errors.ErrNoUserWithUsername
			return
		}
		err = errors.ErrServerError
		return
	}
	return
}
