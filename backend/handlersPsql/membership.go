package handlersPsql

import (
	//"fmt"

	database "github.com/Slowers-team/Slowers-App/database/psql"
	//"github.com/Slowers-team/Slowers-App/utils"

	"github.com/gofiber/fiber/v2"
)

func AddMembership(c *fiber.Ctx) error {
	membership := new(database.Membership)

	if err := c.BodyParser(membership); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	var user_email string

	if err := c.BodyParser(user_email); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	newMembership := database.Membership{
		CreatedAt:    membership.CreatedAt,
		LastModified: membership.LastModified,
		UserEmail:    membership.UserEmail,
		BusinessID:   membership.BusinessID,
		Designation:  membership.Designation,
	}

	_ = newMembership

	return c.SendStatus(204)
}
