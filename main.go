package main

import (
	"html/template"
	"net/http"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
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
