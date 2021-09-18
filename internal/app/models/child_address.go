package models

import (
	"entry-notificator/pkg/database"
	"github.com/google/uuid"
)

type ChildAddress struct {
	database.Model
	ChildID        uuid.UUID
	PostalCode     string `json:"postal_code" validate:"required"`
	PrefectureCode int    `json:"prefecture_code" validate:"required,min=1,max=47"`
	City           string `json:"city" validate:"required"`
	Street         string `json:"street" validate:"required"`
	Room           string `json:"room"`
}
