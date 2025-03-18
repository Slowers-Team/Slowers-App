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

	var user_email string

	if err := c.BodyParser(user_email); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	if business.BusinessName == "" ||
		business.BusinessType == "" ||
		business.BusinessPhoneNumber == "" ||
		business.BusinessEmail == "" ||
		business.PostalCode == 00000 || // tälle joku järkevämpi ratkasu
		business.City == "" {
		return c.Status(400).SendString("All fields are required")
	}

	if !utils.IsEmailValid(business.BusinessEmail) {
		return c.Status(400).SendString("invalid email")
	}

	timestamp := time.Now().Format(time.RFC3339)
	fmt.Println(timestamp)

	newBusiness := database.Business{
		CreatedAt:           business.CreatedAt,
		LastModified:        business.LastModified,
		BusinessName:        business.BusinessName,
		BusinessType:        business.BusinessType,
		BusinessPhoneNumber: business.BusinessPhoneNumber,
		BusinessEmail:       business.BusinessEmail,
		BusinessAddress:     business.BusinessAddress,
		PostalCode:          business.PostalCode,
		City:                business.City,
		Notes:               business.Notes,
	}
	fmt.Println("uusi yritys:", newBusiness)

	createdBusiness, err := db.CreateBusiness(c.Context(), newBusiness)

	if err != nil {
		fmt.Println("Yrityksen luominen ei onnistunut")
		return c.Status(500).SendString(err.Error())
	}

	err = db.AddMembership(c.Context(), user_email, newBusiness.BusinessEmail, "owner")

	if err != nil {
		fmt.Println("Yrityksen omistajan lisäys epäonnistui")
		return c.Status(500).SendString(err.Error())
	}

	fmt.Println("Yrityksen luominen onnistui:", createdBusiness.BusinessName)

	return nil
}
