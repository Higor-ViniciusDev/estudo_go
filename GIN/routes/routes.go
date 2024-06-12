package routes

import (
	"estudo_go/GIN/controller"

	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	r.POST("/alunos", controller.CriaNovoAluno)
	r.GET("/alunos", controller.ExibeTodosAlunos)
	r.GET("/:nome", controller.Saudacao)
	r.GET("/alunos/:id", controller.ExibeAlunoEspecifico)
	r.GET("/alunos/cpf/:cpf", controller.BuscaAlunoPorCpf)
	r.DELETE("/alunos/:id", controller.DeletaAluno)
	r.PATCH("/alunos/:id", controller.EditaAluno)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
