package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatal(err)
	}
	database_env := os.Getenv("MYSQL_DATABASE")
	usr_env := os.Getenv("MYSQL_USER")
	pass_env := os.Getenv("MYSQL_PASSWORD")

	sConexao := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", usr_env, pass_env, database_env)

	db, err := sql.Open("mysql", sConexao)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("conectado")

	rows, err := db.Query("SELECT * FROM usuarios")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	fmt.Println(rows)
}
