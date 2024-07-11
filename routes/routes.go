package routes

import (
	"gin-api-go/controllers"

	"github.com/gin-gonic/gin"

	docs "gin-api-go/docs"
   swaggerfiles "github.com/swaggo/files"
   ginSwagger "github.com/swaggo/gin-swagger"
)

func HandleRequests() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/alunos", controllers.BuscarAlunos)
	r.GET("/alunos/:id", controllers.BuscarAlunoById)
	r.GET("/alunos/cpf/:cpf", controllers.BuscarAlunoByCpf)
	r.GET("/:nome", controllers.Saudacao)
	r.POST("/alunos", controllers.CriarAluno)
	r.PATCH("/alunos/:id", controllers.EditarAluno)
	r.DELETE("/alunos/:id", controllers.DeletarAluno)

	r.Run()
}
