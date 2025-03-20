package handlersPsql

import (
	"fmt"
	"time"

	database "github.com/Slowers-team/Slowers-App/database/psql"
	"github.com/Slowers-team/Slowers-App/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateBusiness(c *fiber.Ctx) error {
	business := new(database.Business)

	if err := c.BodyParser(business); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	if business.Name == "" ||
		business.Type == "" ||
		business.PhoneNumber == "" ||
		business.Email == "" ||
		business.PostalCode == 00000 || // tälle joku järkevämpi ratkasu
		business.City == "" {
		return c.Status(400).SendString("All fields are required")
	}

	if !utils.IsEmailValid(business.Email) {
		return c.Status(400).SendString("invalid email")
	}

	timestamp := time.Now().Format(time.RFC3339)
	fmt.Println(timestamp)

	newBusiness := database.Business{
		CreatedAt:    business.CreatedAt,
		LastModified: business.LastModified,
		Name:         business.Name,
		Type:         business.Type,
		PhoneNumber:  business.PhoneNumber,
		Email:        business.Email,
		PostAddress:  business.PostAddress,
		PostalCode:   business.PostalCode,
		City:         business.City,
		Notes:        business.Notes,
	}
	fmt.Println("uusi yritys:", newBusiness)

	createdBusiness, err := db.CreateBusiness(c.Context(), newBusiness)

	if err != nil {
		fmt.Println("Yrityksen luominen ei onnistunut")
		return c.Status(500).SendString(err.Error())
	}

	fmt.Println("Yrityksen luominen onnistui:", createdBusiness.Name)

	return nil
}
