package models

import "github.com/dgrijalva/jwt-go"

type User struct {
	Username string `json:"username"`
	Password string	`json:"password"`
	jwt.StandardClaims
}

type Exception struct {
	Message string `json:"message"`
}