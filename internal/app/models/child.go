package models

import (
	"entry-notificator/pkg/database"
	"github.com/jackc/pgtype"
)

type Child struct {
	database.Model
	FirstName   string       `json:"first_name" validate:"required"`
	FamilyName  string       `json:"family_name" validate:"required"`
	BirthDate   pgtype.Date  `json:"birth_date" validate:"required"`
	Address     ChildAddress `json:"address"`
	Contact     ChildContact `json:"contact"`
	Cards       []Card       `json:"cards"`
	Attendances []Attendance `json:"attendances"`
}
