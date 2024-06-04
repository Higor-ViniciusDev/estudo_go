package controller

import (
	"ESTUDO_GO/crud/models"
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

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")

	models.DeletaProduto(idProduto)

	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	dadosProd := models.BuscarProdutosId(idProduto)

	temp.ExecuteTemplate(w, "Edit", dadosProd)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		quantidade := r.FormValue("quantidade")
		preco := r.FormValue("preco")
		IdProduto := r.FormValue("btnSalvar")

		if IdProduto == "" {
			panic("NÂO FOI POSSIVEL IDENTIFICAR O ID DO PRODUTO")
		}

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

		models.Update(nome, quantiConvertido, precoConvertido, descricao, IdProduto)
	}

	http.Redirect(w, r, "/", 301)
}
