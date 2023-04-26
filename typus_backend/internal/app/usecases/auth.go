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

type AuthUsecase struct {
	repo *repos.AuthRepo
}

func NewAuthUsecase() *AuthUsecase {
	r, err := repos.NewAuthRepo()
	if err != nil {
		log.Fatal("Could not create the AuthRepo")
	}

	return &AuthUsecase{
		repo: r,
	}
}

func (u *AuthUsecase) RegisterUser(creds models.RegisterCredentials) (id int, err error) {

	h := sha256.New()
	h.Write([]byte(creds.Password))
	encrypted_pwd := hex.EncodeToString(h.Sum(nil))

	user := &models.User{
		Username:     creds.Username,
		Email:        creds.Email,
		EncryptedPwd: encrypted_pwd,
		Role:         1,
	}

	user, err = u.repo.CreateUser(user)

	if err != nil {
		return 0, err
	}

	return user.ID, nil
}

func (u *AuthUsecase) LoginUser(creds models.LoginCredentials) (token string, err error) {
	var user *models.User

	if creds.Email == "" {
		user, err = u.repo.GetUserByUsername(creds.Username)
	} else {
		user, err = u.repo.GetUserByEmail(creds.Email)
	}

	if err != nil {
		return "", err
	}

	h := sha256.New()
	h.Write([]byte(creds.Password))
	creds_encrypted_pwd := hex.EncodeToString(h.Sum(nil))

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
