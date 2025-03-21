package handlersPsql

import (
	"fmt"

	database "github.com/Slowers-team/Slowers-App/database/psql"
	"github.com/Slowers-team/Slowers-App/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateBusiness(c *fiber.Ctx) error {
	business := new(database.Business)

	if err := c.BodyParser(business); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	if business.BusinessName == "" ||
		business.Type == "" ||
		business.PhoneNumber == "" ||
		business.Email == "" ||
		business.PostalCode == "" || // tälle joku järkevämpi ratkasu
		business.City == "" {
		return c.Status(400).SendString("All fields are required")
	}

	if !utils.IsEmailValid(business.Email) {
		return c.Status(400).SendString("invalid email")
	}

	newBusiness := database.Business{
		CreatedAt:      business.CreatedAt,
		LastModified:   business.LastModified,
		BusinessName:   business.BusinessName,
		BusinessIdCode: business.BusinessIdCode,
		Type:           business.Type,
		PhoneNumber:    business.PhoneNumber,
		Email:          business.Email,
		PostAddress:    business.PostAddress,
		PostalCode:     business.PostalCode,
		City:           business.City,
		Notes:          business.Notes,
	}

	createdBusiness, err := db.CreateBusiness(c.Context(), newBusiness)

	if err != nil {
		fmt.Println("Creating business unsuccessful")
		return c.Status(500).SendString(err.Error())
	}

	fmt.Println("Creating business successful:", createdBusiness.BusinessName)

	return c.SendStatus(204)
}
