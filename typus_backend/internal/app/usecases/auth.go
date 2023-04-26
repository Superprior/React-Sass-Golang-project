package usecases

import (
	"backend/internal/app/models"
)

type AuthUsecase struct{}

func (u *AuthUsecase) RegisterUser(creds models.RegisterCredentials) (id int, err error) {
	return 0, nil
}

func (u *AuthUsecase) LoginUser(creds models.LoginCredentials) (role int8, err error) {
	return 1, nil
}
