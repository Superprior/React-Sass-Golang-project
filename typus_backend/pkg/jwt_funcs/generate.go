package jwt_funcs

import (
	"os"

	"github.com/golang-jwt/jwt"
)

func GenerateJWT(username string, email string, role int8) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["username"] = username
	claims["email"] = email
	claims["role"] = role

	secretKey := os.Getenv("JWT_KEY")
	if secretKey == "" {
		return "", nil
	}

	tokenString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
