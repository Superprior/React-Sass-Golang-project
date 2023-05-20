package usecases

import (
	"backend/internal/errors"
	"backend/internal/models"
	"backend/internal/repos"
	"backend/pkg/jwt_funcs"
	"crypto/sha256"
	"encoding/hex"
	"log"
)

// Auth usecase definition.
// Contains its repo to perform db queries.
type AuthUsecase struct {
	repo *repos.AuthRepo
}

// Create a new AuthUsecase.
func NewAuthUsecase() (uc *AuthUsecase) {
	r, err := repos.NewAuthRepo()
	if err != nil {
		log.Fatal("Could not create the AuthRepo")
	}

	uc = &AuthUsecase{repo: r}
	return
}

// Usecase (inner logic) for user registration.
func (u *AuthUsecase) RegisterUser(creds models.RegisterCredentials) (id int, err error) {

	// Encrypt the provided password
	// TODO: make a separate function for that

	h := sha256.New()
	h.Write([]byte(creds.Password))
	encrypted_pwd := hex.EncodeToString(h.Sum(nil))

	// Create a new User instance
	user := &models.User{
		Username:     creds.Username,
		Email:        creds.Email,
		EncryptedPwd: encrypted_pwd,
		Role:         models.USER,
	}

	// Call the repo method to insert the instance into the database
	createdUser, err := u.repo.CreateInstance(user)
	id = createdUser.ID

	return
}

// Usecase (inner logic) for logging users in.
func (u *AuthUsecase) LoginUser(creds models.LoginCredentials) (token string, err error) {
	var user *models.User

	// Determine if username or email was provided
	if creds.Email == "" {
		user, err = u.repo.GetInstanceByUsername(creds.Username)
	} else {
		user, err = u.repo.GetInstanceByEmail(creds.Email)
	}

	if err != nil {
		return
	}

	// Encrypt the provided password
	h := sha256.New()
	h.Write([]byte(creds.Password))
	creds_encrypted_pwd := hex.EncodeToString(h.Sum(nil))

	// Compare the provided' and database's password hashes
	if user.EncryptedPwd == creds_encrypted_pwd {
		token, err = jwt_funcs.GenerateJWT(user.Username, user.Email, int8(user.Role))
		if err != nil {
			err = errors.ErrServerError
			return
		}
		return
	} else {
		err = errors.ErrInvalidCredentials
		return
	}
}
