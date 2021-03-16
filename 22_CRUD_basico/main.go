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
	router.HandleFunc("/usuarios", server.CreateUser).Methods(http.MethodPost)

	fmt.Println("Escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
