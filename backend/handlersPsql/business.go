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

	var user_email string

	if err := c.BodyParser(user_email); err != nil {
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
		Address:        business.Address,
		PostalCode:     business.PostalCode,
		City:           business.City,
		AdditionalInfo: business.AdditionalInfo,
	}

	createdBusiness, err := db.CreateBusiness(c.Context(), newBusiness)

	if err != nil {
		fmt.Println("Yrityksen luominen ei onnistunut")
		return c.Status(500).SendString(err.Error())
	}

	var member database.Membership

	newMember, err := db.AddMembership(c.Context(), member) // tänne handler-kutsu

	_ = newMember

	if err != nil {
		fmt.Println("Yrityksen omistajan lisäys epäonnistui")
		return c.Status(500).SendString(err.Error())
	}

	fmt.Println("Creating business successful:", createdBusiness.BusinessName)

	return c.SendStatus(204)
}
