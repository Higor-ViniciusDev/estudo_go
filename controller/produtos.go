package controller

import (
	"CRUD/list_produtos/models"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOspRodutos := models.BuscarProdutos()

	temp.ExecuteTemplate(w, "Index", todosOspRodutos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		quantidade := r.FormValue("quantidade")
		preco := r.FormValue("preco")

		precoConvertido, err := strconv.ParseFloat(preco, 64)

		if err != nil {
			fmt.Println("ERRO AO CONVERTER VALOR ERRO: ", err)
			os.Exit(1)
		}

		quantiConvertido, err := strconv.Atoi(quantidade)

		if err != nil {
			fmt.Println("ERRO AO CONVERTER QUANTIDADE ERRO: ", err)
			os.Exit(1)
		}

		models.CriaNovoProduto(nome, quantiConvertido, precoConvertido, descricao)
	}

	http.Redirect(w, r, "/", 301)
}
