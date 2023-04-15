package jwt

import (
	"backend/pkg/loggers"
	"os"

	"github.com/golang-jwt/jwt"
)

func GenerateJWT(username string, email string, role int8) (string, error) {
	token := jwt.New(jwt.SigningMethodEdDSA)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["username"] = username
	claims["email"] = email
	claims["role"] = role

	secretKey := os.Getenv("JWT_KEY")
	if secretKey == "" {
		loggers.LogEnvError("JWT_KEY")
		return "", nil
	}

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
