package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializePersonRoutes(router *gin.Engine) {
	api := router.Group("")
	{
		api.POST("pessoas", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "POST pessoas",
			})
		})
		api.GET("pessoas/:id", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "pessoasbyid",
			})
		})
		api.GET("pessoas", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "pessoas by term",
			})
		})
		api.GET("contagem-pessoas", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "contagem-pessoas",
			})
		})
	}
}
