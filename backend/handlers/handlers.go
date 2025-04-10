package handlers

import (
	"fmt"

	"github.com/Slowers-team/Slowers-App/databases/mongo"
	"github.com/Slowers-team/Slowers-App/databases/sql"
	"github.com/cloudinary/cloudinary-go"
	"github.com/gofiber/fiber/v2"
)

var MongoDb mongo.Database
var sqlDb sql.Database

var cld *cloudinary.Cloudinary

func SetDatabases(newMongoDb mongo.Database, newSqlDb sql.Database) {
	MongoDb = newMongoDb
	sqlDb = newSqlDb
}

func SetCloudinary(newCld *cloudinary.Cloudinary) {
	cld = newCld
}

func GetCurrentUser(c *fiber.Ctx) (string, error) {
	id, ok := c.Locals("userID").(string)
	if !ok {
		return "none", fmt.Errorf("userID not set in local storage")
	}
	return id, nil
}

// func GetCurrentUser(c *fiber.Ctx) (MongoDb.ObjectID, error) {
// 	id, ok := c.Locals("userID").(string)
// 	if !ok {
// 		return database.NilObjectID, fmt.Errorf("userID not set in local storage")
// 	}

// 	userID, err := database.ParseID(id)
// 	if err != nil {
// 		return database.NilObjectID, err
// 	}
// 	return userID, nil
// }

func GetCurrentBusiness(c *fiber.Ctx) (string, error) {
	id, ok := c.Locals("businessID").(string)
	if !ok {
		return "none", fmt.Errorf("businessID not set in local storage")
	}
	return id, nil
}

func ResetDatabase(c *fiber.Ctx) error {
	if err := MongoDb.Clear(); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	if err := sqlDb.Clear(); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.SendString("Database reset successful")
}

func HealthCheck(c *fiber.Ctx) error {
	return c.SendString("OK")
}
