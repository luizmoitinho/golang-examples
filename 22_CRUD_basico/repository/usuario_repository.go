package repository

import (
	"database/sql"
	"fmt"

	//"repository/DbConnection"

	"crud/models"

	_ "github.com/go-sql-driver/mysql" // Driver de conexao com o Mysql
)

var db *sql.DB
var err error

func init() {

	db, err = Connect()
	if err != nil {
		fmt.Println("Erro ao estabelecer conexao com o banco de dados: %v", err)
	}
}

func InsertUsuario(usuario models.Usuario) (models.Usuario, error) {
	statement, err := db.Prepare("INSERT INTO tb_usuarios (nome, email) VALUES (?, ?)")
	if err != nil {
		return usuario, fmt.Errorf("usuario_repository.Insert(): erro ao criar o statement")
	}

	insert, err := statement.Exec(usuario.Nome, usuario.Email)
	if err != nil {
		return usuario, fmt.Errorf("usuario_repository.Insert(): erro ao executar o statement")
	}

	lastId, err := insert.LastInsertId()
	usuario.ID = lastId
	if err != nil {
		return usuario, fmt.Errorf("usuario_repository.Insert(): erro ao obter o Id inserido  ")
	}
	defer db.Close()
	defer statement.Close()

	return usuario, nil

}
