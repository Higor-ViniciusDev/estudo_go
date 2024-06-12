package controller

import (
	"estudo_go/GIN/database"
	"estudo_go/GIN/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ExibeTodosAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)

	c.JSON(200, alunos)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")

	c.JSON(200, gin.H{
		"API DIZ: ": "OLA MEU AMIGO " + nome + " TUDO BELEZA",
	})
}

func CriaNovoAluno(c *gin.Context) {
	var aluno models.Aluno

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

func ExibeAlunoEspecifico(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")

	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusFound, gin.H{
			"not found": "Aluno não encotrado",
		})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

func DeletaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")

	database.DB.Delete(&aluno, id)

	c.JSON(http.StatusOK, gin.H{
		"data": "Aluno deletado com sucesso",
	})
}

func EditaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error na requisicao": err.Error()})
		return
	}

	database.DB.Model(&aluno).UpdateColumns(aluno)

	c.JSON(200, aluno)
}

func BuscaAlunoPorCpf(c *gin.Context) {
	var aluno models.Aluno
	cpf := c.Params.ByName("cpf")

	database.DB.Where(&models.Aluno{Cpf: cpf}).First(&aluno)

	if aluno.ID == 0 {
		c.JSON(http.StatusFound, gin.H{
			"not found": "Aluno com o cpf: " + cpf + " não encontrado",
		})
		return
	}
	//Teste pull
	c.JSON(http.StatusOK, aluno)
}
