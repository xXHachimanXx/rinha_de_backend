package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePersonHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "POST pessoas",
	})
}

func FindPersonById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pessoasbyid",
	})
}

func FindPersonByTerm(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pessoas by term",
	})
}

func GetPersonCounter(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "contagem-pessoas",
	})
}
