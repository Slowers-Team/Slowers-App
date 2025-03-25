package handlersPsql

import (
	"fmt"

	database "github.com/Slowers-team/Slowers-App/database/psql"
	"github.com/Slowers-team/Slowers-App/enums"
	"github.com/Slowers-team/Slowers-App/utils"

	"github.com/gofiber/fiber/v2"
)

type UserEmail struct {
	UserEmail string
}

func CreateBusiness(c *fiber.Ctx) error {
	business := new(database.Business)
	userEmail := new(UserEmail)

	if err := c.BodyParser(business); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	if err := c.BodyParser(userEmail); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	if business.BusinessName == "" ||
		business.Type == "" ||
		business.PhoneNumber == "" ||
		business.Email == "" ||
		business.PostalCode == "" ||
		business.City == "" || // tälle joku järkevämpi ratkasu
		userEmail.UserEmail == "" {
		return c.Status(400).SendString("All fields are required")
	}

	if !utils.IsEmailValid(business.Email) {
		return c.Status(400).SendString("invalid business email")
	}

	if !utils.IsEmailValid(userEmail.UserEmail) {
		return c.Status(400).SendString("invalid user email")
	}

	if !utils.IsBusinessIdCodeValid(business.BusinessIdCode) {
		return c.Status(400).SendString("invalid business id code")
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

	//var member database.Membership

	//newMember, err := db.AddMembership(c.Context(), member) // tänne handler-kutsu

	//_ = newMember

	// TÄSTÄ RIVISTÄ

	designation, err := enums.DesignationFromString("owner")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	newMember := &database.Membership{
		UserEmail:    business.Email,
		BusinessID:   createdBusiness.ID,
		Designation:  designation.String(),
		BusinessName: business.BusinessName,
	}

	if err := AddMembership(c, newMember); err != nil {
		fmt.Println("Yrityksen omistajan lisäys epäonnistui")
		return c.Status(500).SendString(err.Error())
	}
	// TÄNNE SAAKKA EPÄVARMAA

	fmt.Println("Creating business successful:", createdBusiness.BusinessName)

	return c.SendStatus(204)
}
