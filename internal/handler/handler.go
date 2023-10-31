package handler

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
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

func sendSuccess(ctx *gin.Context, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, data)
}

func CreatePersonHandler(ctx *gin.Context) {
	var createPersonRequest CreatePersonRequest

	if err := ctx.ShouldBindJSON(&createPersonRequest); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := createPersonRequest.Validate(); err != nil {
		logger.Errorf("Create person handler error: %v", err)
		return
	}

	// Formato da string de data
	dateFormat := "2006-01-02"
	birthdateParsed, err := time.Parse(dateFormat, createPersonRequest.Birthdate)
	if err != nil {
		logger.Errorf("Create person handler error: %v", err)
		return
	}

	stringStack := strings.Join(createPersonRequest.Stack, ", ")

	person := schemas.BuildPerson(
		createPersonRequest.Nickname,
		createPersonRequest.Name,
		birthdateParsed,
		createPersonRequest.Stack,
	)

	if err := db.Create(&person).Error; err != nil {
		logger.Errorf("Create person in database error: %v", err)
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, person)
}

func FindPersonById(ctx *gin.Context) {
	var person schemas.Person

	personID := ctx.Param("id")

	if result := db.Where("id = ?", personID).First(&person); result.Error != nil {
		logger.Errorf("Find person in database error: %v", result.Error)
		sendError(ctx, http.StatusInternalServerError, result.Error.Error())
		return
	}
	sendSuccess(ctx, &person)
}

func FindPersonByTerm(ctx *gin.Context) {
	var person schemas.Person

	term := ctx.Query("t")

	if result := db.Table("person").Where("searchable ILIKE ?", term).Find(&person); result.Error != nil {
		logger.Errorf("Find person by term in database error: %v", result.Error)
		sendError(ctx, http.StatusInternalServerError, result.Error.Error())
		return
	}

	sendSuccess(ctx, &person)
}

func GetPersonCounter(ctx *gin.Context) {
	var count int64

	db.Table("person").Count(&count)

	sendSuccess(ctx, &count)
}
