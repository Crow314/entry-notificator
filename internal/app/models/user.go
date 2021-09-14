package models

import (
	"entry-notificator/pkg/database"
)

type User struct {
	database.Model
	Name      string
	Cards     []Card
	Receivers []Receiver
}
