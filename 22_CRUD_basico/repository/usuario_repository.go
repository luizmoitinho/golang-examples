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

func GetUserById(ID int64) (models.Usuario, error) {
	var usuario models.Usuario

	query, err := db.Query("SELECT id, nome, email FROM tb_usuarios WHERE id = ?", ID)
	if err != nil {
		return usuario, fmt.Errorf("usuario_repository.getUserById(): erro ao criar o statement")
	}

	if query.Next() {
		if err := query.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); err != nil {
			return usuario, fmt.Errorf("usuario_repository.GetUsers(): %v", err)
		}
	}

	return usuario, nil
}

func GetUsers() ([]models.Usuario, error) {

	var usuarios []models.Usuario

	query, err := db.Query("SELECT id, nome, email FROM tb_usuarios")
	if err != nil {
		return usuarios, fmt.Errorf("usuario_repository.GetUsers(): erro ao criar a query")
	}

	defer query.Close()
	defer db.Close()

	for query.Next() {

		var usuario models.Usuario

		if err := query.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); err != nil {
			return usuarios, fmt.Errorf("usuario_repository.GetUsers(): %v", err)
		}

		usuarios = append(usuarios, usuario)

	}

	return usuarios, nil
}
