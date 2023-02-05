package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"secapi/middleware"
	"secapi/models"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type User models.User

// extract and return the signing key
func getSecretKey() string {
	key := os.Getenv("SIGNING_KEY")
	return key
}

// public endpoint
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to this public endpoint"))
}

// endpoint to log in
func login(w http.ResponseWriter, r *http.Request) {
	var user User

	json.NewDecoder(r.Body).Decode(&user)

	if user.Username == "admin" && user.Password == "password123" {
		// correct credentials
		token, _ := middleware.GenerateToken()
		w.Header().Set("Authorization", token)
		w.Write([]byte("Authenticated"))

	} else {
		// invalid credentials
		w.Write([]byte("Invalid Creds"))
	}
}

// private endpoint
func reserved(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("[*] This is a highly classified endpoint"))
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", home).Methods("GET")
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/secure", middleware.Authenticate()).Methods("POST")
}