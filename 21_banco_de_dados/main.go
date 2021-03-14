package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// config, err := ini.Load("db_local.config")
	// if err != nil {
	// 	log.Printf("Erro ao ler arquivo: %v", err)
	// }

	// fmt.Println(config)

	//usuario:senha@/banco
	strCon := "devbook:123456@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", strCon)
	if err != nil {
		log.Fatal("Erro ao tentar abrir conexão com o banco de dados: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Erro ao tentar conectar com o banco de dados: %v", err)
	}

	fmt.Println("Conexao com banco de dados estabelecida.")
}
