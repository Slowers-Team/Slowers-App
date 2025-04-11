package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Slowers-team/Slowers-App/databases/sql"
	"github.com/Slowers-team/Slowers-App/handlers"
)

func TestValidateBusinessWithCorrectInput(t *testing.T) {
	correctBusiness := sql.Business{
		ID:             1,
		BusinessName:   "Test Business",
		BusinessIdCode: "1234567-8",
		Type:           "grower",
		PhoneNumber:    "0101234567",
		Email:          "tester@test.fi",
		Address:        "Imaginary road 1",
		PostalCode:     "98765",
		City:           "Flowertown",
		AdditionalInfo: "No notes",
	}
	err := handlers.ValidateBusiness(correctBusiness)
	assert.NoError(t, err, "ValidateBusiness() should not return an error")
}

func TestValidateBusinessWithInCorrectBusinessEmail(t *testing.T) {
	incorrectBusiness := sql.Business{
		ID:             1,
		BusinessName:   "Test Business",
		BusinessIdCode: "1234567-8",
		Type:           "grower",
		PhoneNumber:    "0101234567",
		Email:          "testertest.fi",
		Address:        "Imaginary road 1",
		PostalCode:     "98765",
		City:           "Flowertown",
		AdditionalInfo: "No notes",
	}
	err := handlers.ValidateBusiness(incorrectBusiness)
	assert.ErrorContains(t, err, "invalid business email")
}

func TestValidateBusinessWithInCorrectBusinessIdCode(t *testing.T) {
	incorrectBusiness := sql.Business{
		ID:             1,
		BusinessName:   "Test Business",
		BusinessIdCode: "12345678",
		Type:           "grower",
		PhoneNumber:    "0101234567",
		Email:          "tester@test.fi",
		Address:        "Imaginary road 1",
		PostalCode:     "98765",
		City:           "Flowertown",
		AdditionalInfo: "No notes",
	}
	err := handlers.ValidateBusiness(incorrectBusiness)
	assert.ErrorContains(t, err, "invalid business id code")
}

func TestValidateBusinessWithInCorrectPostalCode(t *testing.T) {
	incorrectBusiness := sql.Business{
		ID:             1,
		BusinessName:   "Test Business",
		BusinessIdCode: "1234567-8",
		Type:           "grower",
		PhoneNumber:    "0101234567",
		Email:          "tester@test.fi",
		Address:        "Imaginary road 1",
		PostalCode:     "ABCDE",
		City:           "Flowertown",
		AdditionalInfo: "No notes",
	}
	err := handlers.ValidateBusiness(incorrectBusiness)
	assert.ErrorContains(t, err, "invalid postal code")
}

func TestValidateBusinessWithInCorrectPhoneNumber(t *testing.T) {
	incorrectBusiness := sql.Business{
		ID:             1,
		BusinessName:   "Test Business",
		BusinessIdCode: "1234567-8",
		Type:           "grower",
		PhoneNumber:    "000",
		Email:          "tester@test.fi",
		Address:        "Imaginary road 1",
		PostalCode:     "98765",
		City:           "Flowertown",
		AdditionalInfo: "No notes",
	}
	err := handlers.ValidateBusiness(incorrectBusiness)
	assert.ErrorContains(t, err, "invalid phone number")
}

func TestValidateBusinessWithEmptyField(t *testing.T) {
	incorrectBusiness := sql.Business{
		ID:             1,
		BusinessName:   "Test Business",
		BusinessIdCode: "1234567-8",
		Type:           "",
		PhoneNumber:    "0101234567",
		Email:          "tester@test.fi",
		Address:        "Imaginary road 1",
		PostalCode:     "98765",
		City:           "Flowertown",
		AdditionalInfo: "No notes",
	}
	err := handlers.ValidateBusiness(incorrectBusiness)
	assert.ErrorContains(t, err, "all fields are required")
}

func TestValidateUserEmailWithCorrectInput(t *testing.T) {
	correctEmail := handlers.UserEmail{
		UserEmail: "testuser@test.fi",
	}
	err := handlers.ValidateUserEmail(correctEmail)
	assert.NoError(t, err, "ValidateUserMail() should not return an error")
}

func TestValidateUserEmailWithIncorrectEmail(t *testing.T) {
	correctEmail := handlers.UserEmail{
		UserEmail: "testuser@testfi",
	}
	err := handlers.ValidateUserEmail(correctEmail)
	assert.ErrorContains(t, err, "invalid user email")
}

func TestValidateUserEmailWithEmptyEmail(t *testing.T) {
	correctEmail := handlers.UserEmail{
		UserEmail: "",
	}
	err := handlers.ValidateUserEmail(correctEmail)
	assert.ErrorContains(t, err, "all fields are required")
}
