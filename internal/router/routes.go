package router

import (
	"github.com/gin-gonic/gin"
	"github.com/xXHachimanXx/rinha_de_backend/internal/handler"
)

func initializePersonRoutes(router *gin.Engine) {
	handler.InitializeHandler()

	api := router.Group("")
	{
		api.POST("pessoas", handler.CreatePersonHandler)
		api.GET("pessoas/:id", handler.FindPersonById)
		api.GET("pessoas", handler.FindPersonByTerm)
		api.GET("contagem-pessoas", handler.GetPersonCounter)
	}
}
