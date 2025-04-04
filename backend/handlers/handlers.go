package handlers

import (
	"fmt"

	"github.com/Slowers-team/Slowers-App/database"
	"github.com/cloudinary/cloudinary-go"
	"github.com/gofiber/fiber/v2"
)

var db database.Database

var cld *cloudinary.Cloudinary

func SetDatabase(newDb database.Database) {
	db = newDb
}

func SetCloudinary(newCld *cloudinary.Cloudinary) {
	cld = newCld
}

func GetCurrentUser(c *fiber.Ctx) (database.ObjectID, error) {
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

func ResetDatabase(c *fiber.Ctx) error {
	if err := db.Clear(); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.SendString("Database reset successful")
}

func HealthCheck(c *fiber.Ctx) error {
	return c.SendString("OK")
}
