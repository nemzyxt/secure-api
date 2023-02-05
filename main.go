package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// public endpoint
func landing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to this public endpoint"))
}

func main() {
	fmt.Println("Hello friend")
}