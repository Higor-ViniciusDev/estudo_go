package controllers

import (
	"Higor-ViniciusDev/estudo_go/API/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "PAGINA PRINCIPAL")
}

func TodasPersonalidade(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalidades)
}

func UmaPersonalidadeEspecifica(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idPesquisa := vars["id"]

	for _, persona := range models.Personalidades {
		if strconv.Itoa(persona.Id) == idPesquisa {
			json.NewEncoder(w).Encode(persona)
		}
	}
}
