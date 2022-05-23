package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func conectaComBancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=postgres password=postgres host=localhost sslmode=disabled"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	db := conectaComBancoDeDados()
	defer db.Close()
	http.HandleFunc("/", index)
	http.ListenAndServe(":8100", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	Produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Bela", Preco: 49.5, Quantidade: 1000},
		{Nome: "Notebook", Descricao: "Digno", Preco: 4900.99, Quantidade: 100},
		{Nome: "Tablet", Descricao: "Leve", Preco: 1999.37, Quantidade: 200},
		{Nome: "SmarTV", Descricao: "Gigante", Preco: 15200.15, Quantidade: 50}}

	temp.ExecuteTemplate(w, "index", Produtos)
}
