package jwt_funcs

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	authorized bool
	username   string
	email      string
	role       int8
}

func ValidateJWT(tokenString string) (*Claims, error) {

	/*
		This function parses a token, checks its signature and returns Claims object.
	*/

	claims := jwt.MapClaims{}
	secretKey := os.Getenv("JWT_KEY")
	if secretKey == "" {
		return nil, fmt.Errorf("Could not get access the secret key.")
	}

	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("Could not parse the JWT.")
	}

	// Assert the provided variable's types

	authorized, ok := claims["authorized"].(bool)
	if !ok {
		fmt.Println("AUTHORIZED")
		return nil, fmt.Errorf("AUTHORIZED claim did not pass the type assertion.")
	}
	username, ok := claims["username"].(string)
	if !ok {
		fmt.Println("USERNAME")
		return nil, fmt.Errorf("USERNAME claim did not pass the type assertion.")
	}
	email, ok := claims["email"].(string)
	if !ok {
		fmt.Println("EMAIL")
		return nil, fmt.Errorf("EMAIL claim did not pass the type assertion.")
	}

	// Convert ROLE claim into int8

	extractedRole, ok := claims["role"].(int)
	role := int8(extractedRole)

	// Construct an object and return it

	extractedClaims := Claims{
		authorized: authorized,
		username:   username,
		email:      email,
		role:       role,
	}

	return &extractedClaims, nil
}
