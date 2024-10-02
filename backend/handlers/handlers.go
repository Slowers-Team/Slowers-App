package handlers

import (
	"github.com/Slowers-team/Slowers-App/database"
)

var db database.Database

func SetDatabase(newDb database.Database) {
	db = newDb
}
