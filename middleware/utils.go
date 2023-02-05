package middleware

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

func getSigningKey() string {
	if err := godotenv.Load("../.env"); err != nil {
		panic(err)
	}
	return os.Getenv("SIGNING_KEY")
}

func generateToken(username string) (string, error) {
	key := getSigningKey()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": username,
	})

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Minute * 60).Unix()

	tokenString, err := token.SignedString(key)
	return tokenString, err
}

func verifyToken(tokenString string) (jwt.Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return getSigningKey(), nil
	})
	if err != nil {
		return nil, err
	}

	return token.Claims, err
}