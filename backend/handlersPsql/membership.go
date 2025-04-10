package handlersPsql

import (
	"fmt"
	"strconv"
	"strings"

	database "github.com/Slowers-team/Slowers-App/database/psql"

	"github.com/gofiber/fiber/v2"
)

func AddMembership(c *fiber.Ctx, membership *database.Membership) error {
	_, err := db.AddMembership(c.Context(), *membership)
	if err != nil {
		fmt.Println("Failed adding membership")
		return c.Status(500).SendString(err.Error())
	}

	return c.SendStatus(204)
}

func AddMembershipHelper(c *fiber.Ctx) error {
	membership := new(database.Membership)

	if err := c.BodyParser(membership); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	_, err := db.AddMembership(c.Context(), *membership)
	if err != nil {
		fmt.Println("Jäsenyyden lisääminen epäonnistui")
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
	fmt.Println(result)
	fmt.Println(c.JSON(result))
	return c.JSON(result)
}

func GetAllMembersInBusiness(c *fiber.Ctx) error {
	businessID, err := strconv.Atoi(c.Params("businessID"))
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	fmt.Println("HELLO")
	fmt.Println("BUSINESS ID;", businessID)
	result, err := db.GetAllMembersInBusiness(c.Context(), businessID)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	var csvData []string
	for _, membership := range result {
		membershipCSV := fmt.Sprintf("%d,%s,%s,%s,%d,%s",
			membership.ID,
			membership.CreatedAt,
			membership.LastModified,
			membership.UserEmail,
			membership.BusinessID,
			membership.Designation,
		)
		csvData = append(csvData, membershipCSV)
	}
	csvResponse := strings.Join(csvData, "\n")

	fmt.Println(csvResponse)

	return c.SendString(csvResponse)
}

func DeleteMembership(c *fiber.Ctx) error {
	membership := new(database.Membership)
	if err := c.BodyParser(membership); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	err := db.DeleteMembership(c.Context(), membership.UserEmail, membership.BusinessID)
	if err != nil {
		fmt.Println("Failed deleting membership. Membership might not exist.")
		return c.Status(500).SendString(err.Error())
	}
	return c.SendStatus(204)
}
