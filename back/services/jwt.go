package services

import (
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/vmonneger/FinalProject/configs"
)

type Token struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

// CreateToken : takes userId as parameter,
// generates JWT token and
// Return JWT token string
func CreateToken(id, email string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Token{
		ID:    id,
		Email: email,
	})

	tokenString, err := token.SignedString([]byte(configs.EnvJwtSecret()))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
