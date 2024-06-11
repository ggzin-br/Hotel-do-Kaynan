package database

import (
	"database/sql"
	//_ "github.com/lib/pq" 
	_ "github.com/go-sql-driver/mysql" // <- Estarei utilizando isto
)

func ConectaComBancoDeDados() *sql.DB {
	//conexao := "user=kaynanfullstackpython dbname=python password=python>java host=localhost sslmode=disable"
	db, err := sql.Open("mysql", "kaynanfullstackpython:python>java@tcp(127.0.0.1:3306)/python")
	if err != nil {
		panic(err.Error())
	}
	return db
}