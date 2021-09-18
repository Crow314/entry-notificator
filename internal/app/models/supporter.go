package models

import (
	"entry-notificator/pkg/database"
	"github.com/jackc/pgtype"
)

type Supporter struct {
	database.Model
	FirstName  string           `json:"first_name" validate:"required"`
	FamilyName string           `json:"family_name" validate:"required"`
	BirthDate  pgtype.Date      `json:"birth_date" validate:"required"`
	Email      string           `json:"email" validate:"required,email"`
	Address    SupporterAddress `json:"address"`
}
