package models

import (
	"entry-notificator/pkg/database"
	"github.com/google/uuid"
)

type Card struct {
	database.Model
	ChildID   uuid.UUID `json:"child_id" validate:"required"`
	Token     string    `json:"token" validate:"required"`
	TokenType string    `json:"token_type" validate:"required"`
}
