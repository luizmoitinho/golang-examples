package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"crud/models"

	"crud/repository"
)

//CreateUser ... insere um usuario no banco de dados
func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	var usuario models.Usuario

	if err = json.Unmarshal(bodyRequest, &usuario); err != nil {
		w.Write([]byte("Erro ao converter o usuário para struct"))
		return
	}

	usuario, err = repository.InsertUsuario(usuario)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso, Id: %d", usuario.ID)))

}

//GetUSer ... retorna todos os usuarios
func GetUsers(w http.ResponseWriter, r *http.Request) {
	usuarios, err := repository.GetUsers()
	fmt.Println(usuarios)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	if err := json.NewEncoder(w).Encode(usuarios); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

}

//GetUSer ... retorna um usuário especifico
// func GetUser(w http.ResponseWriter, r *http.Request)
// {

// }
