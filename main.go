package main

import (
	"fmt"
	"log"
	"net/http"

	handler "Login-API/handler"
)

func main() {
	http.HandleFunc("/login", handler.Login)

	fmt.Println("Server started at localhost or 127.0.0.1 : 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
