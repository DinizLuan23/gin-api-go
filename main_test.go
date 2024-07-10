package main

import (
	"bytes"
	"encoding/json"
	"gin-api-go/controllers"
	"gin-api-go/database"
	"gin-api-go/models"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var id int

func SetupRotas() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func CriarAlunoMock() {
	aluno := models.Aluno{Nome: "Aluno Teste", Cpf: "11790368910", Rg: "139901266"}
	database.DB.Create(&aluno)
	id = int(aluno.ID)
}

func DeletarAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, id)
}

func TestStatusCode(t *testing.T) {
	r := SetupRotas()
	r.GET("/:nome", controllers.Saudacao)
	req, _ := http.NewRequest("GET", "/gui", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code, "Status diferente do esperado")
	mockResposta := `{"API diz:":"E ai gui, tudo beleza?"}`
	bodyResposta, _ := io.ReadAll(resposta.Body)
	assert.Equal(t, mockResposta, string(bodyResposta))
}

func TestBuscarAlunos(t *testing.T) {
	database.ConectaBanco()
	CriarAlunoMock()
	defer DeletarAlunoMock()
	r := SetupRotas()
	r.GET("/alunos", controllers.BuscarAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestBuscarAlunoCpf(t *testing.T) {
	database.ConectaBanco()
	CriarAlunoMock()
	defer DeletarAlunoMock()
	r := SetupRotas()
	r.GET("/alunos/cpf/:cpf", controllers.BuscarAlunoByCpf)
	req, _ := http.NewRequest("GET", "/alunos/cpf/11790368910", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestBuscarAlunoId(t *testing.T) {
	database.ConectaBanco()
	CriarAlunoMock()
	defer DeletarAlunoMock()
	r := SetupRotas()
	r.GET("/alunos/:id", controllers.BuscarAlunoById)
	path := "/alunos/" + strconv.Itoa(id)
	req, _ := http.NewRequest("GET", path, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	var alunoMock models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMock)

	assert.Equal(t, "Aluno Teste", alunoMock.Nome)
	assert.Equal(t, "11790368910", alunoMock.Cpf)
	assert.Equal(t, "139901266", alunoMock.Rg)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestDeletarAluno(t *testing.T) {
	database.ConectaBanco()
	CriarAlunoMock()
	r := SetupRotas()
	r.DELETE("/alunos/:id", controllers.DeletarAluno)
	pathBusca := "/alunos/" + strconv.Itoa(id)
	req, _ := http.NewRequest("DELETE", pathBusca, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestEditarAluno(t *testing.T){
	database.ConectaBanco()
	CriarAlunoMock()
	defer DeletarAlunoMock()
	r := SetupRotas()
	r.PATCH("/alunos/:id", controllers.EditarAluno)
	aluno := models.Aluno{Nome: "Aluno Teste", Cpf: "40790368910", Rg: "409901266"}
	valorJson, _ := json.Marshal(aluno)
	pathEditar := "/alunos/" + strconv.Itoa(id)
	req, _ := http.NewRequest("PATCH", pathEditar, bytes.NewBuffer(valorJson))
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	var alunoMock models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMock)
	assert.Equal(t, "40790368910", alunoMock.Cpf)
	assert.Equal(t, "409901266", alunoMock.Rg)
	assert.Equal(t, "Aluno Teste", alunoMock.Nome)
}