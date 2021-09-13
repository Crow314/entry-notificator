package models

import (
	"entry-notificator/pkg/database"
	"github.com/google/uuid"
)

type Receiver struct {
	database.Model
	UserID  uuid.UUID `json:"user"`
	Address string
}
