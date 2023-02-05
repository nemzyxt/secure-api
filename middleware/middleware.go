package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"secapi/models"

	"github.com/dgrijalva/jwt-go"
)

type Exception models.Exception

func GenerateToken(username string, password string, key string) string {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["password"] = password
	claims["exp"] = time.Now().Add(time.Minute * 60).Unix()

	tokenString, _ := token.SignedString([]byte(key))
	return tokenString
}

func ValidateEndpoint(next http.HandlerFunc, key string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader != "" {
			bearerToken := strings.Split(authHeader, " ")
			if len(bearerToken) == 2 {
				token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("there was an error")
					} 
					return []byte(key), nil
				})
				if err != nil {
					json.NewEncoder(w).Encode(Exception{Message: err.Error()})
					return
				}

				if token.Valid {
					next(w, r)
				} else {
					json.NewEncoder(w).Encode(Exception{Message: "invalid auth header"})
					return
				}
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("auth header required"))
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("auth header required"))
		}
	})
}