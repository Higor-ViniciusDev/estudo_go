package controllers

import (
	"Higor-ViniciusDev/estudo_go/API/database"
	"Higor-ViniciusDev/estudo_go/API/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "PAGINA PRINCIPAL")
}

func TodasPersonalidade(w http.ResponseWriter, r *http.Request) {
	var p []models.Personalidade

	database.DB.Find(&p)

	json.NewEncoder(w).Encode(p)
}

func UmaPersonalidadeEspecifica(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idPesquisa := vars["id"]

	var persona models.Personalidade

	database.DB.First(&persona, idPesquisa)
	json.NewEncoder(w).Encode(persona)
}

func CriaNovaPersonalidade(w http.ResponseWriter, r *http.Request) {
	var NovaPersonalidade models.Personalidade

	json.NewDecoder(r.Body).Decode(&NovaPersonalidade)
	database.DB.Create(&NovaPersonalidade)

	json.NewEncoder(w).Encode(NovaPersonalidade)
}

func DeletarPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idPesquisa := vars["id"]
	var personalidade models.Personalidade

	database.DB.Delete(&personalidade, idPesquisa)

	json.NewEncoder(w).Encode(personalidade)
}

func EditaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idPesquisa := vars["id"]
	var personalidade models.Personalidade

	database.DB.First(&personalidade, idPesquisa)
	json.NewDecoder(r.Body).Decode(&personalidade)
	database.DB.Save(personalidade)

	json.NewEncoder(w).Encode(personalidade)
}
