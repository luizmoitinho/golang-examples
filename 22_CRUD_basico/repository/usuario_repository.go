package repository

import (
	"database/sql"
	"fmt"

	//"repository/DbConnection"

	"crud/models"

	_ "github.com/go-sql-driver/mysql" // Driver de conexao com o Mysql
)

func getConnection() (*sql.DB, error) {
	db, err := Connect()
	if err != nil {
		return db, fmt.Errorf("Erro ao estabelecer conexao com o banco de dados: %v", err)
	}
	return db, nil
}

func InsertUsuario(usuario models.Usuario) (models.Usuario, error) {
	db, err := getConnection()
	if err != nil {
		return usuario, fmt.Errorf(err.Error())
	}

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

	db, err := getConnection()
	if err != nil {
		return usuario, fmt.Errorf(err.Error())
	}

	query, err := db.Query("SELECT id, nome, email FROM tb_usuarios WHERE id = ?", ID)
	if err != nil {
		return usuario, fmt.Errorf("usuario_repository.getUserById(): erro ao criar o statement")
	}

	defer query.Close()
	defer db.Close()

	if query.Next() {
		if err := query.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); err != nil {
			return usuario, fmt.Errorf("usuario_repository.GetUsers(): %v", err)
		}
	}

	return usuario, nil
}

func GetUsers() ([]models.Usuario, error) {
	var usuarios []models.Usuario

	db, err := getConnection()
	if err != nil {
		return usuarios, fmt.Errorf(err.Error())
	}

	query, err := db.Query("SELECT id, nome, email FROM tb_usuarios")
	if err != nil {
		return usuarios, fmt.Errorf("usuario_repository.GetUsers(): erro ao criar a query: %v", err)
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

func UpdateUser(usuario models.Usuario) (models.Usuario, error) {

	db, err := getConnection()
	if err != nil {
		return usuario, fmt.Errorf(err.Error())
	}

	statement, err := db.Prepare("UPDATE tb_usuarios set nome = ?, email = ? WHERE id = ?")
	if err != nil {
		return usuario, fmt.Errorf("usuario_repository.UpdateUser(): erro ao criar o statement")
	}

	defer statement.Close()
	defer db.Close()

	if _, err := statement.Exec(usuario.Nome, usuario.Email, usuario.ID); err != nil {
		return usuario, fmt.Errorf("usuario_repository.UpdateUser(): erro ao atualizar o usuario")
	}
	return usuario, nil

}

func DeleteUser(ID int64) error {

	db, err := getConnection()
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	statement, err := db.Prepare("DELETE FROM tb_usuarios WHERE id = ?")
	if err != nil {
		return fmt.Errorf("usuario_repository.DeleteUser(): erro ao criar o statement")
	}

	defer db.Close()
	defer statement.Close()
	if _, err := statement.Exec(ID); err != nil {
		return fmt.Errorf("usuario_repository.DeleteUser(): Erro ao remover o usuário: %v", err)
	}

	return nil
}
