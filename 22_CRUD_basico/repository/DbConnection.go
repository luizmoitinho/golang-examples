package repository

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // Driver de conexao com o Mysql
)

//Connect .. abre conexao com o banco de dados
func Connect() (*sql.DB, error) {

	strCon := "devbook:123456@/devbook?charset=utf8&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", strCon)

	if err != nil {
		fmt.Println("erro: DbConnection.Connect(), %v", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Erro ao tentar conectar com o banco de dados: %v", err)
	}

	return db, nil

}
