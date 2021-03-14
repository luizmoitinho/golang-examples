package main

import (
	"log"
	"net/http"
)

func main() {

	//URI: /home
	http.HandleFunc("/home", home)

	//URI: /usuarios
	http.HandleFunc("/usuarios", usuarios)

	log.Fatal(http.ListenAndServe(":5000", nil)) //up server

}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Página Inicial"))
}

func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Carregar página de usuarios"))
}
