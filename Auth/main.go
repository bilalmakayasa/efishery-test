package main

import (
	"net/http"

	"Training/Auth/src"
)

func main() {
	http.HandleFunc("/", src.Login)
	http.HandleFunc("/register", src.Register)
	http.HandleFunc("/welcome", src.Credential)
	http.ListenAndServe(":8081", nil)
}
