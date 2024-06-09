package routesApi

import (
	"Higor-ViniciusDev/estudo_go/API/controllers"
	"Higor-ViniciusDev/estudo_go/API/database"
	"Higor-ViniciusDev/estudo_go/API/models"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequisicao() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "RUA HIGOR", Historia: "QUE RUA TRISTE DO CARAMBA EM"},
		{Id: 2, Nome: "RUA DA CAROL", Historia: "QUE RUA MAIS FELIZ DA PORRA"},
	}

	database.ConectaBancoDeDados()

	fmt.Println("INICIANDO CONEXÂO E SERVIDOR REST")

	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidade", controllers.TodasPersonalidade).Methods("Get")
	r.HandleFunc("/api/personalidade/{id}", controllers.UmaPersonalidadeEspecifica).Methods("Get")
	panic(http.ListenAndServe(":8000", r))
}
