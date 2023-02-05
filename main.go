package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"secapi/models"

	"github.com/gorilla/mux"
)

type User models.User

// public endpoint
func landing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to this public endpoint"))
}

// endpoint to log in
func login(w http.ResponseWriter, r *http.Request) {
	var user User

	json.NewDecoder(r.Body).Decode(&user)

	if user.Username == "admin" && user.Password == "password123" {
		// correct credentials

	} else {
		// invalid credentials
		w.Write([]byte("invalid creds"))
	}
}

// private endpoint
func private(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("[*] This is a highly classified endpoint"))
}

func main() {
	fmt.Println("Hello friend")
}