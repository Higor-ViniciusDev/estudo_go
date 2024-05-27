package routes

import (
	"CRUD/list_produtos/controller"
	"net/http"
)

func CarregarRoutas() {
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/new", controller.New)
	http.HandleFunc("/insert", controller.Insert)
}
