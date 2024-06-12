package main

import (
	"estudo_go/GIN/database"
	"estudo_go/GIN/routes"
)

func main() {
	database.ConectaBanco()

	routes.HandleRequest()
}
