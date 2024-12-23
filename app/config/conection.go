package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Conection() *sql.DB {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	entrada := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("SERVER"), os.Getenv("DATABASE"))

	db, err := sql.Open("mysql", entrada)

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
