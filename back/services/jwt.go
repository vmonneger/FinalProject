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

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &Token{
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

// // Check if user got token
// func IsAuthorized(h http.Handler) http.Handler {

// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 		// Auth route don't check token
// 		if strings.Contains(r.URL.Path, "/auth") {
// 			h.ServeHTTP(w, r)
// 		} else {
// 			tokenString := r.Header.Get("Authorization")
// 			// Decode from the struct
// 			t := Token{}
// 			token, err := jwt.ParseWithClaims(tokenString, &t, func(token *jwt.Token) (interface{}, error) {
// 				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 					return nil, fmt.Errorf("there was an error in parsing")
// 				}
// 				return []byte(configs.EnvJwtSecret()), nil
// 			})

// 			if err != nil || !token.Valid {
// 				w.WriteHeader(http.StatusUnauthorized)
// 				response := responses.UserResponse{
// 					Status:  http.StatusUnauthorized,
// 					Message: "error",
// 					Data:    map[string]interface{}{"data": err.Error()}}
// 				json.NewEncoder(w).Encode(response)
// 				return
// 			} else {
// 				h.ServeHTTP(w, r)
// 			}
// 		}
// 	})
// }
