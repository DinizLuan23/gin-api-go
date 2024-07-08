package routes

import (
	"gin-api-go/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.BuscarAlunos)
	r.GET("/alunos", controllers.BuscarAlunos)
	r.Run()
}
