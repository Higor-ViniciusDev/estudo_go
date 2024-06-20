package main

import (
	"encoding/json"
	"estudo_go/GIN/controller"
	"estudo_go/GIN/database"
	"estudo_go/GIN/models"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupDasRotasDeTest() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()

	return rotas
}

func TestVerificadorEndPoinDaSaudacao(t *testing.T) {
	r := SetupDasRotasDeTest()

	r.GET("/:nome", controller.Saudacao)

	req, _ := http.NewRequest("GET", "/higor", nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code, "DEVERIAM SER IGUAIS")

	mockResposta := `{"API DIZ: ":"OLA MEU AMIGO higor TUDO BELEZA"}`
	respostaBody, _ := io.ReadAll(response.Body)

	assert.Equal(t, mockResposta, string(respostaBody), "OS DOIS NÂO SÂO IGUAIS")
}

func TestVerificaTodosAlunosPorId(t *testing.T) {
	database.ConectaBanco()
	r := SetupDasRotasDeTest()
	r.GET("/alunos/:id", controller.ExibeAlunoEspecifico)

	req, _ := http.NewRequest("GET", "/alunos/4", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	var alunoMock models.Aluno
	json.Unmarshal(response.Body.Bytes(), &alunoMock)
	assert.Equal(t, 4, int(alunoMock.ID), "Ids dos alunos diferente")
}
