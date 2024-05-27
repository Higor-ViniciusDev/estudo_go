package models

import (
	"CRUD/list_produtos/db"
	"database/sql"
	"fmt"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscarProdutos() []Produto {
	db := db.ConectaComBancoDeDados()

	selectDeTodosOsProdutos, err := db.Query("select ID,NOME,DESCRICAO,PRECO,QUANTIDADE from produto")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	defer db.Close()

	return produtos
}

func CriaNovoProduto(nome string, quantidade int, preco float64, descricao string) {
	dbConxao := db.ConectaComBancoDeDados()

	stmt, err := dbConxao.Prepare("insert into produto(NOME,DESCRICAO,QUANTIDADE,PRECO) values (@p1, @p2, @p3, @p4)")

	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()
	_, erro := stmt.Exec(sql.Named("p1", nome), sql.Named("p2", descricao), sql.Named("p3", quantidade), sql.Named("p4", preco))

	if erro != nil {
		panic(erro.Error())
	}

	fmt.Println("SUCESSO AO INSERIIR, LINHAS INSERIDA ")
	defer dbConxao.Close()
}
