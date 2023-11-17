package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// User構造体
type User struct {
    ID      int
    Name 	string
    Score   int
}

func main() {
    router := mux.NewRouter().StrictSlash(true)
    
	router.HandleFunc("/", home)
    router.HandleFunc("/users", findAllUsers).Methods("GET")
    router.HandleFunc("/users/{id}", findById).Methods("GET")
    router.HandleFunc("/users", createUser).Methods("POST")
	router.HandleFunc("/users", updateUser).Methods("PUT")
	router.HandleFunc("/users", deleteUser).Methods("DELETE")
    
    log.Fatal(http.ListenAndServe(":8080", router))
}

