package handlers

import (
	"fmt"

	"github.com/Slowers-team/Slowers-App/database"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var db database.Database

func SetDatabase(newDb database.Database) {
	db = newDb
}

func GetCurrentUser(c *fiber.Ctx) (*primitive.ObjectID, error) {
	id, ok := c.Locals("userID").(string)
	if !ok {
		return nil, fmt.Errorf("userID not set in local storage")
	}
	userID, err := database.ParseID(id)
	if err != nil {
		return nil, err
	}
	return userID, nil
}
