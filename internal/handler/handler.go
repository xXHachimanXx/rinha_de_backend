package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xXHachimanXx/rinha_de_backend/config"
	"github.com/xXHachimanXx/rinha_de_backend/internal/schemas"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	db = config.GetPostgres()
	db.AutoMigrate(&schemas.Person{})
	db.Model(&schemas.Person{})
	logger = config.GetLogger("Handler")
}

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successfull", op),
		"data":    data,
	})
}

func CreatePersonHandler(ctx *gin.Context) {
	createPersonRequest := CreatePersonRequest{}

	ctx.BindJSON(&createPersonRequest)

	if err := createPersonRequest.Validate(); err != nil {
		logger.Errorf("Create person handler error: %v", err)
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	person := schemas.BuildPerson(
		createPersonRequest.Nickname,
		createPersonRequest.Name,
		createPersonRequest.Birthdate,
		fmt.Sprintf("%v", createPersonRequest.Stack),
	)

	if err := db.Create(&person).Error; err != nil {
		logger.Errorf("Create person in database error: %v", err)
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, "create-person", person)
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
