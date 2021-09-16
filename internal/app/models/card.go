package models

import (
	"entry-notificator/pkg/database"
	"github.com/google/uuid"
)

type Card struct {
	database.Model
	ChildID   uuid.UUID `json:"child_id"`
	Token     string    `json:"token"`
	TokenType string    `json:"token_type"`
}
