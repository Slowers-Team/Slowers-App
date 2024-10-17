package handlers

import (
	"fmt"

	"github.com/Slowers-team/Slowers-App/database"
	"github.com/gofiber/fiber/v2"
)

var db database.Database

func SetDatabase(newDb database.Database) {
	db = newDb
}

func GetCurrentUser(c *fiber.Ctx) (database.ObjectID, error) {
	var userID database.ObjectID

	id, ok := c.Locals("userID").(string)
	if !ok {
		return database.NilObjectID, fmt.Errorf("userID not set in local storage")
	}

	userID, err := database.ParseID(id)
	if err != nil {
		return database.NilObjectID, err
	}
	return userID, nil
}
