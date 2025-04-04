package handlersPsql

import (
	"fmt"
	"strconv"

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
		fmt.Println("Failed adding membership")
		return c.Status(500).SendString(err.Error())
	}

	return c.SendStatus(204)
}

func GetDesignation(c *fiber.Ctx) error {
	userIDStr, err := GetCurrentUser(c)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	userID, err := strconv.Atoi(userIDStr)

	if err != nil {
		return c.Status(400).SendString("Invalid user ID")
	}
	fmt.Println("USERID", userID)

	result, err := db.GetMembershipByUserId(c.Context(), userID)

	if err != nil {
		return c.JSON(err)
	}
	return c.JSON(result)
}

func DeleteMembership(c *fiber.Ctx, user_email string, business_id int) error {

	err := db.DeleteMembership(c.Context(), user_email, business_id)
	if err != nil {
		fmt.Println("Failed deleting membership. Membership might not exist.")
		return c.Status(500).SendString(err.Error())
	}
	return c.SendStatus(204)
}