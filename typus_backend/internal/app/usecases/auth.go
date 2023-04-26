package usecases

import (
	"backend/internal/app/models"
	"backend/internal/app/repos"
	"crypto/sha256"
	"encoding/hex"
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

func (u *AuthUsecase) LoginUser(creds models.LoginCredentials) (role int8, err error) {
	return 1, nil
}
