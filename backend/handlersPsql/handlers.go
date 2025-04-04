package handlersPsql

import (
	"fmt"

	psql "github.com/Slowers-team/Slowers-App/database/psql"
	"github.com/gofiber/fiber/v2"
)

var db psql.SQLDatabase

func SetDatabase(newDb psql.SQLDatabase) {
	db = newDb
}

func GetCurrentUser(c *fiber.Ctx) (string, error) {
	id, ok := c.Locals("userID").(string)
	if !ok {
		return "none", fmt.Errorf("userID not set in local storage")
	}

	// userID, err := psql.ParseID(id)
	// if err != nil {
	// 	return "none", err
	// }
	// return userID, nil
	return id, nil
}

// func GetCurrentEmail(c *fiber.Ctx) (string, error) {
// 	email, ok := c.Locals("user_email").(string)
// 	if !ok {
// 		return "none", fmt.Errorf("user email not set in local storage")
// 	}

// 	return email, nil
// }

// func ResetDatabase(c *fiber.Ctx) error {
// 	if err := db.Clear(); err != nil {
// 		return c.Status(500).SendString(err.Error())
// 	}
// 	return c.SendString("Database reset successful")
// }

// func HealthCheck(c *fiber.Ctx) error {
// 	return c.SendString("OK")
// }

// tämä ehkä turha?
func GetCurrentBusiness(c *fiber.Ctx) (string, error) {
	id, ok := c.Locals("businessID").(string)
	if !ok {
		return "none", fmt.Errorf("businessID not set in local storage")
	}

	return id, nil
}
