package models

import (
	"entry-notificator/pkg/database"
)

type Place struct {
	database.Model
	Name string `json:"name"`
}
