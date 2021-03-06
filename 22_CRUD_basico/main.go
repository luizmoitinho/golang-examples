package main

import (
	"fmt"
	"log"
	"net/http"

	"crud/server"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	//Post
	router.HandleFunc("/usuario", server.CreateUser).Methods(http.MethodPost)
	//Get
	router.HandleFunc("/usuarios", server.GetUsers).Methods(http.MethodGet)
	router.HandleFunc("/usuario/{id}", server.GetUser).Methods(http.MethodGet)
	//Put
	router.HandleFunc("/usuario/{id}", server.UpdateUser).Methods(http.MethodPut)
	//Delete
	router.HandleFunc("/usuario/{id}", server.DeleteUser).Methods(http.MethodDelete)

	fmt.Println("Escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
