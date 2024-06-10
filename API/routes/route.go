package routesApi

import (
	"Higor-ViniciusDev/estudo_go/API/controllers"
	"Higor-ViniciusDev/estudo_go/API/database"
	"Higor-ViniciusDev/estudo_go/API/middleware"
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequisicao() {
	database.ConectaBancoDeDados()

	fmt.Println("INICIANDO CONEXÂO E SERVIDOR REST")

	r := mux.NewRouter()
	r.Use(middleware.ContetType)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidade", controllers.TodasPersonalidade).Methods("Get")
	r.HandleFunc("/api/personalidade/{id}", controllers.UmaPersonalidadeEspecifica).Methods("Get")
	r.HandleFunc("/api/personalidade", controllers.CriaNovaPersonalidade).Methods("Post")
	r.HandleFunc("/api/personalidade/{id}", controllers.DeletarPersonalidade).Methods("Delete")
	r.HandleFunc("/api/personalidade/{id}", controllers.EditaPersonalidade).Methods("Put")
	panic(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
