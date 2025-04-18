package handlers

import (
	"fmt"
	"strconv"

	"github.com/Slowers-team/Slowers-App/databases/sql"
	"github.com/Slowers-team/Slowers-App/utils"

	"github.com/gofiber/fiber/v2"
)

func AddMembership(c *fiber.Ctx) error {
	membership := new(sql.Membership)

	if err := c.BodyParser(membership); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	if _, err := sqlDb.AddMembership(c.Context(), *membership); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.SendStatus(204)
}

func GetDesignation(c *fiber.Ctx) error {
	userID, err := GetCurrentUser(c)

	if err != nil {
		return c.Status(400).SendString("Invalid user ID")
	}

	result, err := sqlDb.GetMembershipByUserId(c.Context(), userID)

	if err != nil {
		return c.JSON(err)
	}
	fmt.Println(result)
	fmt.Println(c.JSON(result))
	return c.JSON(result)
}

func GetAllMembersInBusiness(c *fiber.Ctx) error {
	businessID, err := strconv.Atoi(c.Params("businessID"))
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	result, err := sqlDb.GetAllMembersInBusiness(c.Context(), businessID)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	csvResponse := utils.MembersIntoCSV(result)

	return c.SendString(csvResponse)
}

func DeleteMembership(c *fiber.Ctx) error {
	email := c.Params("email")
	businessIDStr := c.Params("businessID")

	businessID, err := strconv.Atoi(businessIDStr)
	if err != nil {
		return c.Status(400).SendString("Invalid business ID")
	}

	fmt.Println("Säpo", email)

	err = sqlDb.DeleteMembership(c.Context(), email, businessID)
	if err != nil {
		fmt.Println("Failed deleting membership. Membership might not exist.")
		return c.Status(500).SendString(err.Error())
	}
	return c.SendStatus(204)
}

func EditMembership(c *fiber.Ctx) error {
	member := new(sql.Membership)
	if err := c.BodyParser(member); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	err := sqlDb.EditMembership(c.Context(), *member)

	if err != nil {
		fmt.Println("Failed editing membership")
		return c.Status(500).SendString(err.Error())
	}

	return c.SendStatus(204)
}
