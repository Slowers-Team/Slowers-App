package handlers

import (
	"fmt"

	"github.com/Slowers-team/Slowers-App/databases/sql"
	"github.com/Slowers-team/Slowers-App/enums"
	"github.com/Slowers-team/Slowers-App/utils"

	"github.com/gofiber/fiber/v2"
)

type UserEmail struct {
	UserEmail string
}

func ValidateBusiness(business sql.Business) error {
	if business.BusinessName == "" ||
		business.Type == "" ||
		business.PhoneNumber == "" ||
		business.Email == "" ||
		business.PostalCode == "" ||
		business.City == "" { // tälle joku järkevämpi ratkasu
		return fmt.Errorf("all fields are required")
	}

	if !utils.IsEmailValid(business.Email) {
		return fmt.Errorf("invalid business email")
	}

	if !utils.IsBusinessIdCodeValid(business.BusinessIdCode) {
		return fmt.Errorf("invalid business id code")
	}

	if !utils.IsPostalCodeValid(business.PostalCode) {
		return fmt.Errorf("invalid postal code")
	}

	if !utils.IsPhoneNumberValid(business.PhoneNumber) {
		return fmt.Errorf("invalid phone number")
	}

	return nil
}

func ValidateUserEmail(userEmail UserEmail) error {
	if userEmail.UserEmail == "" {
		return fmt.Errorf("all fields are required")
	}

	if !utils.IsEmailValid(userEmail.UserEmail) {
		return fmt.Errorf("invalid user email")
	}

	return nil
}

func CreateBusiness(c *fiber.Ctx) error {
	business := new(sql.Business)
	userEmail := new(UserEmail)

	if err := c.BodyParser(business); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	if err := c.BodyParser(userEmail); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	if err := ValidateBusiness(*business); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	if err := ValidateUserEmail(*userEmail); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	if business.Type == "retailer" && business.Delivery == "yes" {
		return c.Status(400).SendString("cannot have retailer business with delivery")
	}

	newBusiness := sql.Business{
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
		Delivery:       business.Delivery,
	}

	createdBusiness, err := sqlDb.CreateBusiness(c.Context(), newBusiness)

	if err != nil {
		fmt.Println("Yrityksen luominen ei onnistunut")
		return c.Status(500).SendString(err.Error())
	}

	designation, err := enums.DesignationFromString("owner")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	newMember := &sql.Membership{
		UserEmail:   userEmail.UserEmail,
		BusinessID:  createdBusiness.ID,
		Designation: designation.String(),
	}

	if _, err := sqlDb.AddMembership(c.Context(), *newMember); err != nil {
		fmt.Println("Yrityksen omistajan lisäys epäonnistui")
		return c.Status(500).SendString(err.Error())
	}

	fmt.Println("Creating business successful:", createdBusiness.BusinessName)

	return c.SendStatus(204)
}

func GetBusiness(c *fiber.Ctx) error {
	userID, err := GetCurrentUser(c)

	if err != nil {
		return c.Status(400).SendString("Invalid business ID")
	}
	result, err := sqlDb.GetBusinessByUserID(c.Context(), userID)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(result)
}
