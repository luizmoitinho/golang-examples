package main

import (
	"log"
	"net/http"
)

func main() {

	//URI: /home
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Página Inicial"))
	})

	//URI: /usuarios
	http.HandleFunc("/usuarios", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Carregar página de usuarios"))
	})

	log.Fatal(http.ListenAndServe(":5000", nil)) //up server

}
