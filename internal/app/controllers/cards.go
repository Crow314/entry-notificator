package controllers

import (
	"entry-notificator/internal/app/models"
	"entry-notificator/pkg/database"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func CreateCard(ctx *gin.Context) {
	db := database.DB()

	card := models.Card{}
	err := ctx.Bind(&card)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid format"})
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