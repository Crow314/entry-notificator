package models

import (
	"entry-notificator/pkg/database"
	"github.com/jackc/pgtype"
)

type Supporter struct {
	database.Model
	FirstName  string           `json:"first_name"`
	FamilyName string           `json:"family_name"`
	BirthDate  pgtype.Date      `json:"birth_date"`
	Email      string           `json:"email"`
	Address    SupporterAddress `json:"address"`
}
