package middlewares

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/vmonneger/FinalProject/configs"
	"github.com/vmonneger/FinalProject/responses"
	"github.com/vmonneger/FinalProject/services"
)

// Check if user got token
func IsAuthorized(h http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Auth route don't check token
		if strings.Contains(r.URL.Path, "/auth") {
			h.ServeHTTP(w, r)
		} else {
			reqToken := r.Header.Get("Authorization")
			tokenString := strings.Split(reqToken, "Bearer ")[1]

			t := services.Token{}
			token, err := jwt.ParseWithClaims(tokenString, &t, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("there was an error in parsing")
				}
				return []byte(configs.EnvJwtSecret()), nil
			})

			if err != nil || !token.Valid {
				w.WriteHeader(http.StatusUnauthorized)
				response := responses.RequestResponse{
					Status:  http.StatusUnauthorized,
					Message: "error",
					Data:    map[string]interface{}{"data": err.Error()}}
				json.NewEncoder(w).Encode(response)
				return
			} else {
				h.ServeHTTP(w, r)
			}
		}
	})
}
