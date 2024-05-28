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

		p.Id = id
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

func DeletaProduto(idProduto string) {
	db := db.ConectaComBancoDeDados()

	stm, err := db.Prepare("DELETE FROM PRODUTO WHERE ID = @p1")

	if err != nil {
		panic(err.Error())
	}

	defer stm.Close()

	_, erro := stm.Exec(sql.Named("p1", idProduto))

	if erro != nil {
		panic(erro.Error())
	}

	fmt.Println("SUCESSO AO EXCLUIR PRODUTO ID = ", idProduto)
	defer db.Close()
}

func BuscarProdutosId(idProduto string) Produto {
	db := db.ConectaComBancoDeDados()

	stm, err := db.Query("SELECT * FROM PRODUTO WHERE ID = @p1", idProduto)

	if err != nil {
		panic(err.Error())
	}

	ProdutoBusca := Produto{}

	for stm.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = stm.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		ProdutoBusca.Id = id
		ProdutoBusca.Nome = nome
		ProdutoBusca.Descricao = descricao
		ProdutoBusca.Preco = preco
		ProdutoBusca.Quantidade = quantidade
	}

	defer db.Close()

	return ProdutoBusca
}

func Update(nome string, quantidade int, preco float64, descricao string, idProduto string) {
	dbConxao := db.ConectaComBancoDeDados()

	stmt, err := dbConxao.Prepare("UPDATE PRODUTO SET NOME = @p1, DESCRICAO = @p2, PRECO =  @p3,QUANTIDADE = @p4 WHERE ID = @p5")

	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()
	_, erro := stmt.Exec(sql.Named("p1", nome), sql.Named("p2", descricao), sql.Named("p3", quantidade), sql.Named("p4", preco), sql.Named("p5", idProduto))

	if erro != nil {
		panic(erro.Error())
	}

	fmt.Println("DADOS ALTERADO COM SUCESSO")
	defer dbConxao.Close()
}
