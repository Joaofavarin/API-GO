package main

import (
	"fmt"
	"log"
	"net/http"

	"example.com/myapp/handlers"

	"github.com/gorilla/mux"
)

func main() {
	log.Println(2)

	router := mux.NewRouter()
	router.HandleFunc("/users", handlers.GetUsers).Methods("GET")

	fmt.Println("API is on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
