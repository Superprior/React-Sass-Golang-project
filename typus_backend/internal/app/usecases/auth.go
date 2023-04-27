package usecases

import (
	"backend/internal/app/models"
	"backend/internal/app/repos"
	"backend/pkg/jwt_funcs"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
)

// Auth usecases.
// Contains its repo.
type AuthUsecase struct {
	repo *repos.AuthRepo
}

// NewAuthUsercase
// This function creates a new AuthUsecase.
func NewAuthUsecase() *AuthUsecase {
	r, err := repos.NewAuthRepo()
	if err != nil {
		log.Fatal("Could not create the AuthRepo")
	}

	return &AuthUsecase{
		repo: r,
	}
}

// RegisterUser
// This function implements usecase (inner logic) for user registration.
func (u *AuthUsecase) RegisterUser(creds models.RegisterCredentials) (id int, err error) {

	// Encrypt the provided password

	h := sha256.New()
	h.Write([]byte(creds.Password))
	encrypted_pwd := hex.EncodeToString(h.Sum(nil))

	// Create a new User instance
	user := &models.User{
		Username:     creds.Username,
		Email:        creds.Email,
		EncryptedPwd: encrypted_pwd,
		Role:         1,
	}

	// Call the repo method to insert the instance into the database
	user, err = u.repo.CreateInstance(user)

	if err != nil {
		return 0, err
	}

	// Return the ID of the created user
	return user.ID, nil
}

// LoginUser
// This function implements usecase (inner logic) for logging users in.
func (u *AuthUsecase) LoginUser(creds models.LoginCredentials) (token string, err error) {
	var user *models.User

	// Determine if username of email was provided
	if creds.Email == "" {
		user, err = u.repo.GetInstanceByUsername(creds.Username)
	} else {
		user, err = u.repo.GetInstanceByEmail(creds.Email)
	}

	if err != nil {
		return "", err
	}

	// Encrypt the provided password
	h := sha256.New()
	h.Write([]byte(creds.Password))
	creds_encrypted_pwd := hex.EncodeToString(h.Sum(nil))

	// Compare the provided' and database's password hashes
	if user.EncryptedPwd == creds_encrypted_pwd {
		token, err := jwt_funcs.GenerateJWT(user.Username, user.Email, int8(user.Role))
		if err != nil {
			return "", err
		}

		return token, nil
	} else {
		return "", fmt.Errorf("Incorrect password provided")
	}
}
