package main

import (
	"fmt"
	"net/http"

	"github.com/bilalmakayasa/efishery-test/Auth/src"
)

func main() {
	port := ":8081"
	http.HandleFunc("/login", src.Login)
	http.HandleFunc("/register", src.Register)
	http.HandleFunc("/welcome", src.Credential)
	fmt.Printf("Authentication services listening on port:%v", port)
	http.ListenAndServe(port, nil)
}
