package controllers

import (
	"entry-notificator/internal/app/models"
	"entry-notificator/pkg/database"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

func CreateChild(ctx *gin.Context) {
	db := database.DB()

	var child models.Child
	err := ctx.Bind(&child)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid format"})
		return
	}

	err = validator.New().Struct(child)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Validation failed"})
		return
	}

	result := db.Create(&child)
	err = result.Error
	if err != nil {
		log.Warn(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Server error"})
		return
	}

	ctx.JSON(http.StatusOK, child)
}
