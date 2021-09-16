package models

import (
	"entry-notificator/pkg/database"
	"github.com/google/uuid"
)

type SupporterAddress struct {
	database.Model
	SupporterID    uuid.UUID
	PostalCode     string `json:"postal_code"`
	PrefectureCode int    `json:"prefecture_code"`
	City           string `json:"city"`
	Street         string `json:"street"`
	Room           string `json:"room"`
}
