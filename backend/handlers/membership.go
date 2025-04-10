package handlers

import (
	"fmt"
	"strconv"

	"github.com/Slowers-team/Slowers-App/databases/sql"

	"github.com/gofiber/fiber/v2"
)

func AddMembership(c *fiber.Ctx, membership *sql.Membership) error {
	_, err := sqlDb.AddMembership(c.Context(), *membership)
	if err != nil {
		fmt.Println("Failed adding membership")
		return c.Status(500).SendString(err.Error())
	}

	return c.SendStatus(204)
}

func AddMembershipHelper(c *fiber.Ctx) error {
	membership := new(sql.Membership)

	if err := c.BodyParser(membership); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	if err := AddMembership(c, membership); err != nil {
		fmt.Println("Jäsenen lisäys epäonnistui")
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

	result, err := sqlDb.GetMembershipByUserId(c.Context(), userID)

	if err != nil {
		return c.JSON(err)
	}
	return c.JSON(result)
}

func GetAllMembersInBusiness(c *fiber.Ctx) error {
	var businessID int

	if err := c.BodyParser(businessID); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	result, err := sqlDb.GetAllMembersInBusiness(c.Context(), businessID)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(result)
}

func DeleteMembership(c *fiber.Ctx) error {
	membership := new(sql.Membership)
	if err := c.BodyParser(membership); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	err := sqlDb.DeleteMembership(c.Context(), membership.UserEmail, membership.BusinessID)
	if err != nil {
		fmt.Println("Failed deleting membership. Membership might not exist.")
		return c.Status(500).SendString(err.Error())
	}
	return c.SendStatus(204)
}
