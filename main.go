package main

import (
	"CRUD/list_produtos/routes"
	"net/http"
)

func main() {
	routes.CarregarRoutas()
	http.ListenAndServe(":8000", nil)
}
