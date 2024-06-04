package main

import (
	"ESTUDO_GO/crud/routes"
	"net/http"
)

func main() {
	routes.CarregarRoutas()
	http.ListenAndServe(":8000", nil)
}
