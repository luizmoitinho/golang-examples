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

	fmt.Println(usuario)
	usuario, err = repository.InsertUsuario(usuario)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso, Id: %d", usuario.ID)))

}
