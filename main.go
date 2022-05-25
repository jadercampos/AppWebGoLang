package main

import (
	"html/template"
	"net/http"

	"appwebgolang.com/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8100", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "index", todosOsProdutos)
}
