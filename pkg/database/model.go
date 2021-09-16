package database

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Model struct {
	gorm.Model
	ID uuid.UUID `gorm:"primaryKey;default:gen_random_uuid()"`
}
