package models

import (
	"entry-notificator/pkg/database"
)

type User struct {
	database.Model
	Cards     []Card
	Receivers []Receiver
}
