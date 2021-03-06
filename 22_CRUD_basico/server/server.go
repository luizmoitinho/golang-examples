package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"crud/models"

	"crud/repository"

	"github.com/gorilla/mux"
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
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso, Id: %d", usuario.ID)))

}

//GetUSer ... retorna todos os usuarios
func GetUsers(w http.ResponseWriter, r *http.Request) {
	usuarios, err := repository.GetUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if err := json.NewEncoder(w).Encode(usuarios); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

}

//GetUSer ... retorna um usuário especifico
func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Falha ao tentar receber o valor do parâmetro id"))
		return
	}

	usuario, err := repository.GetUserById(ID)
	if usuario.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Usuário nao foi encontrado"))
		return
	}
	if err := json.NewEncoder(w).Encode(usuario); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

}

//UpdateUser ... atualiza um usuario especifico
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Falha ao tentar receber o valor do parâmetro id"))
		return
	}

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

	usuario.ID = ID
	usuario, err = repository.UpdateUser(usuario)
	if err = json.Unmarshal(bodyRequest, &usuario); err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	if err := json.NewEncoder(w).Encode(usuario); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Falha ao tentar receber o valor do parâmetro id"))
		return
	}

	err = repository.DeleteUser(ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
