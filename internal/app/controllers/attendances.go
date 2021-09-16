package controllers

import (
	"entry-notificator/internal/app/models"
	"entry-notificator/pkg/database"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/now"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
)

func CreateAttendance(ctx *gin.Context) {
	db := database.DB()

	card := models.Card{}
	err := ctx.Bind(&card)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid format"})
		return
	}

	// Retrieve Place
	placeId, err := uuid.Parse(ctx.Param("place_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid url format"})
		return
	}
	place := models.Place{}
	result := db.First(&place, placeId)
	err = result.Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "The place is unknown"})
		} else {
			log.Warn(err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Server error"})
		}
		return
	}

	// Retrieve Card (retrieve child_id)
	result = db.Where("token = ? AND token_type = ?", card.Token, card.TokenType).First(&card)
	err = result.Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "The card is not registered"})
		} else {
			log.Warn(err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Server error"})
		}
		return
	}

	// Retrieve Child
	child := models.Child{}
	result = db.First(&child, card.ChildID)
	err = result.Error
	if err != nil {
		log.Warn(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Server error"})
		return
	}

	// Retrieve whether or not the child is stay at
	newestAttendance := models.Attendance{}
	isEntered := false
	result = db.Where("child_id = ? AND place_id = ? AND updated_at >= ?", child.ID, placeId, now.BeginningOfDay()).Order("updated_at desc").First(&newestAttendance)
	err = result.Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			isEntered = false
		} else {
			log.Warn(err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Server error"})
			return
		}
	} else {
		isEntered = newestAttendance.IsEntry
	}

	// Create Attendance
	attendance := models.Attendance{ChildID: child.ID, PlaceID: placeId, IsEntry: !isEntered}
	result = db.Create(&attendance)
	err = result.Error
	if err != nil {
		log.Warn(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Server error"})
		return
	}

	ctx.JSON(http.StatusOK, attendance)
}
