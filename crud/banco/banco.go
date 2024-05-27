package banco

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // Driver de conexao do Mysql
	"github.com/joho/godotenv"
)

func Conectar() (*sql.DB, error) {
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
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
