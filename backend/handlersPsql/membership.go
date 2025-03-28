package handlersPsql

import (
	"fmt"

	database "github.com/Slowers-team/Slowers-App/database/psql"

	"github.com/gofiber/fiber/v2"
)

func AddMembership(c *fiber.Ctx, membership *database.Membership) error {
	if membership == nil {
		membership = new(database.Membership)

		if err := c.BodyParser(membership); err != nil {
			return c.Status(400).SendString(err.Error())
		}
	}

	_, err := db.AddMembership(c.Context(), *membership)
	if err != nil {
		fmt.Println("Jäsenyyden lisääminen epäonnistui")
		return c.Status(500).SendString(err.Error())
	}

	return c.SendStatus(204)
}

// func GetDesignation(c *fiber.Ctx) error {
// 	email, err := GetCurrentEmail(c)

// 	if err != nil {
// 		return c.Status(500).SendString(err.Error())
// 	}
// 	// userID, err := strconv.Atoi(userIDStr)

// 	// if err != nil {
// 	// 	return c.Status(400).SendString("Invalid user ID")
// 	// }
// 	result, err := db.GetDesignationByEmail(c.Context(), email)

// 	if err != nil {
// 		return c.Status(500).SendString(err.Error())
// 	}
// 	return c.JSON(result)
// }
