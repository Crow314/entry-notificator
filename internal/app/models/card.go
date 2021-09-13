package models

import (
	"entry-notificator/pkg/database"
	"github.com/google/uuid"
)

type Card struct {
	database.Model
	UserID    uuid.UUID `json:"user"`
	Token     string    `json:"token"`
	TokenType string    `json:"token_type"`
}
