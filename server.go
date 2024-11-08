package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// http.HandleFunc("/", handler)
	fmt.Println("Server started on http://localhost:8080")

	CreateDB()

	router := mux.NewRouter()
	router.HandleFunc("/", handler).Methods("GET")

	router.HandleFunc("/writeups", GetWriteups).Methods("GET")
	router.HandleFunc("/writeups", AddWriteup).Methods("POST")

	router.HandleFunc("/users", GetUser).Methods("GET")
	router.HandleFunc("/users", AddUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}
