package controllers

import "github.com/gin-gonic/gin"

func BuscarAlunos(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":   1,
		"nome": "Luan Diniz",
	})
}

func Saudacao(c *gin.Context) {
	
}	