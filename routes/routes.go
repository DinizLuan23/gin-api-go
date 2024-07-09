package routes

import (
	"gin-api-go/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.BuscarAlunos)
	r.GET("/alunos/:id", controllers.BuscarAlunoById)
	r.GET("/alunos/cpf/:cpf", controllers.BuscarAlunoByCpf)
	r.GET("/:nome", controllers.Saudacao)
	r.POST("/alunos", controllers.CriarAluno)
	r.PATCH("/alunos/:id", controllers.EditarAluno)
	r.DELETE("/alunos/:id", controllers.DeletarAluno)

	r.Run()
}
