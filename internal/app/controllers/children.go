package controllers

import (
	"entry-notificator/internal/app/models"
	"entry-notificator/pkg/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateChild(ctx *gin.Context) {
	db := database.DB()

	child := models.Child{}
	err := ctx.Bind(&child)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid format"})
		return
	}

	result := db.Create(&child)
	err = result.Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Server error"})
		return
	}

	ctx.JSON(http.StatusOK, child)
}
