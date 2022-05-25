package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectaComBancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=postgres password=postgres host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
