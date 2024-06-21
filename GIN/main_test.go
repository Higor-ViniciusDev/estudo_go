package main

import (
	"bytes"
	"encoding/json"
	"estudo_go/GIN/controller"
	"estudo_go/GIN/database"
	"estudo_go/GIN/models"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func CriaAlunoMock() {
	var Aluno1 = models.Aluno{
		Nome: "CECILIA MINHA PRINCESA",
		Cpf:  "012303433",
		Rg:   "00000000000",
	}

	database.ConectaBanco()
	database.DB.Create(&Aluno1)

	ID = int(Aluno1.ID)
}

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

func TestDeletaAlunoHandler(t *testing.T) {
	database.ConectaBanco()
	// CriaAlunoMock()

	r := SetupDasRotasDeTest()
	r.DELETE("/alunos/:id", controller.DeletaAluno)

	req, _ := http.NewRequest("DELETE", "/alunos/5", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code, "Aluna não foi excluida, favor validar code, foi esperado 200")
}

func TestEditaUmAlunoHandle(t *testing.T) {
	database.ConectaBanco()

	CriaAlunoMock()

	r := SetupDasRotasDeTest()
	r.PATCH("/alunos/:id", controller.EditaAluno)

	pathParaEditar := "/alunos/" + strconv.Itoa(ID)

	novoAluno := models.Aluno{
		Nome: "TROQUE MINHA LINDA",
		Cpf:  "123456789",
		Rg:   "98745612332",
	}

	alunoConvertido, _ := json.Marshal(novoAluno)

	req, _ := http.NewRequest("PATCH", pathParaEditar, bytes.NewBuffer(alunoConvertido))
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	var alunoMockTest models.Aluno
	json.Unmarshal(response.Body.Bytes(), &alunoMockTest)

	assert.Equal(t, "123456789", string(alunoMockTest.Cpf), "Ids dos alunos diferente")
	assert.Equal(t, "98745612332", string(alunoMockTest.Rg), "Ids dos alunos diferente")
}
