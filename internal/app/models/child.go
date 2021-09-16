package models

import (
	"entry-notificator/pkg/database"
	"github.com/jackc/pgtype"
)

type Child struct {
	database.Model
	FirstName   string       `json:"first_name"`
	FamilyName  string       `json:"family_name"`
	BirthDate   pgtype.Date  `json:"birth_date"`
	Address     ChildAddress `json:"address"`
	Contact     ChildContact `json:"contact"`
	Cards       []Card       `json:"cards"`
	Attendances []Attendance `json:"attendances"`
}
