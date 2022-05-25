package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"appwebgolang.com/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "Index", todosOsProdutos)
}
func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", "")
}
func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco, err := strconv.ParseFloat(r.FormValue("preco"), 64)
		if err != nil {
			log.Println("Erro ao converter preço!")
		}
		quantidade, err := strconv.Atoi(r.FormValue("quantidade"))
		if err != nil {
			log.Println("Erro ao converter quantidade!")
		}
		models.CriarNovoProduto(nome, descricao, preco, quantidade)

	}
	temp.ExecuteTemplate(w, "New", "")
}
func Delete(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	models.DeletaProduto(idDoProduto)
	http.Redirect(w, r, "/", 301)
}
func Edit(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	produto := models.EditaProduto(idDoProduto)
	temp.ExecuteTemplate(w, "Edit", produto)
}
func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")

		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			log.Println("Erro ao converter id!")
		}
		preco, err := strconv.ParseFloat(r.FormValue("preco"), 64)
		if err != nil {
			log.Println("Erro ao converter preço!")
		}
		quantidade, err := strconv.Atoi(r.FormValue("quantidade"))
		if err != nil {
			log.Println("Erro ao converter quantidade!")
		}
		models.AtualizaProduto(nome, descricao, preco, id, quantidade)
	}
	http.Redirect(w, r, "/", 301)
}
