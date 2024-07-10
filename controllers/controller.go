package controllers

import (
	"gin-api-go/database"
	"gin-api-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
	_"github.com/swaggo/swag/example/celler/httputil"
)

// BuscarAlunos godoc
// @Summary Mostrar os alunos
// @Description Rota para buscar todos os alunos
// @Tags alunos
// @Accept json
// @Produce json
// @Success 200 {object} models.Aluno
// @Failure 400 {object} httputil.HTTPError
// @Router /alunos [get]
func BuscarAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(200, alunos)
}

func BuscarAlunoById(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.Find(&aluno, id)
	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{ "Not Found": "Aluno não encontrado" })
		return
	}

	c.JSON(http.StatusOK, aluno)
}

func BuscarAlunoByCpf(c *gin.Context) {
	var aluno models.Aluno
	cpf := c.Param("cpf")
	database.DB.Where(&models.Aluno{Cpf: cpf}).First(&aluno)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{ "Not Found": "Aluno não encontrado" })
		return
	}

	c.JSON(http.StatusOK, aluno)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz:": "E ai " + nome + ", tudo beleza?",
	})
}

func CriarAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.ValidarDadosAluno(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

func EditarAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.ValidarDadosAluno(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&aluno).UpdateColumns(aluno)
	c.JSON(http.StatusOK, aluno)
}

func DeletarAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.Delete(&aluno, id)

	c.JSON(http.StatusOK, gin.H{"data": "Aluno deletado com sucesso"})
}
