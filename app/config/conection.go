package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Conection() *sql.DB {
	db, err := sql.Open("mysql", "joao-geraldo:senha@tcp(127.0.0.1:3306)/gerenciador_tarefas?parseTime=true")

	if err != nil {
		log.Fatal("Erro na conexão com o banco de dados!", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Erro ao verificar conexão com o banco de dados!", err)
	}
	fmt.Println("Conexão Realizada!")
	return db
}
