package models

import (
	"entry-notificator/pkg/database"
	"github.com/google/uuid"
)

type ChildContact struct {
	database.Model
	ChildID     uuid.UUID
	Email       string `json:"email" validate:"required,email"`
	PhoneNumber string `json:"phone_number" validate:"required"`
}
