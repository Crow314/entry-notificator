package controllers

import (
	"entry-notificator/internal/app/models"
	"entry-notificator/pkg/database"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

func CreateCard(ctx *gin.Context) {
	db := database.DB()

	var card models.Card
	err := ctx.Bind(&card)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid format"})
		return
	}

	err = validator.New().Struct(card)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Validation failed"})
		return
	}

	result := db.Create(&card)
	err = result.Error
	if err != nil {
		log.Warn(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Server error"})
		return
	}

	ctx.JSON(http.StatusOK, card)
}
