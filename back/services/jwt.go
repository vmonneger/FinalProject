package services

import (
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/vmonneger/FinalProject/configs"
)

// Token ,contains data that will enrypted in JWT token
// When jwt token will decrypt, token model will returns
// Need this model to authenticate and validate resources access by loggedIn user
type Token struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

// CreateToken : takes userId as parameter,
// generates JWT token and
// Return JWT token string
func CreateToken(id, email string) (map[string]string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Token{
		ID:    id,
		Email: email,
	})
	// token -> string. Only server knows this secret (foobar).
	tokenString, err := token.SignedString([]byte(configs.EnvJwtSecret()))
	if err != nil {
		return nil, err
	}
	m := make(map[string]string)
	m["token"] = tokenString // set response data
	return m, nil
}
