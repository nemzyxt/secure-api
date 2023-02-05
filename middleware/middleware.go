package middleware

import (
	"time"


	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(username string, password string, key string) string {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["password"] = password
	claims["exp"] = time.Now().Add(time.Minute * 60).Unix()

	tokenString, _ := token.SignedString(key)
	return tokenString
}

func AuthEndpoint() {
	
}