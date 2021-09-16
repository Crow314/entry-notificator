package models

import (
	"entry-notificator/pkg/database"
	"github.com/google/uuid"
)

type Attendance struct {
	database.Model
	ChildID uuid.UUID
	PlaceID uuid.UUID
	IsEntry bool
}
