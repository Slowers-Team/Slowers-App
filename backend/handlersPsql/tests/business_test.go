package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	database "github.com/Slowers-team/Slowers-App/database/psql"
	"github.com/Slowers-team/Slowers-App/handlersPsql"
)

func TestValidateBusinessWithCorrectInput(t *testing.T) {
	correctBusiness := database.Business{
		ID:             1,
		BusinessName:   "Test Business",
		BusinessIdCode: "1234567-8",
		Type:           "grower",
		PhoneNumber:    "010234567",
		Email:          "tester@test.fi",
		Address:        "Imaginary road 1",
		PostalCode:     "98765",
		City:           "Flowertown",
		AdditionalInfo: "No notes",
	}
	err := handlersPsql.ValidateBusiness(correctBusiness)
	assert.NoError(t, err, "ValidateBusiness() should not return an error")
}

func TestValidateBusinessWithInCorrectBusinessEmail(t *testing.T) {
	incorrectBusiness := database.Business{
		ID:             1,
		BusinessName:   "Test Business",
		BusinessIdCode: "1234567-8",
		Type:           "grower",
		PhoneNumber:    "010234567",
		Email:          "testertest.fi",
		Address:        "Imaginary road 1",
		PostalCode:     "98765",
		City:           "Flowertown",
		AdditionalInfo: "No notes",
	}
	err := handlersPsql.ValidateBusiness(incorrectBusiness)
	assert.ErrorContains(t, err, "invalid business email")
}

func TestValidateBusinessWithInCorrectBusinessIdCode(t *testing.T) {
	incorrectBusiness := database.Business{
		ID:             1,
		BusinessName:   "Test Business",
		BusinessIdCode: "12345678",
		Type:           "grower",
		PhoneNumber:    "010234567",
		Email:          "tester@test.fi",
		Address:        "Imaginary road 1",
		PostalCode:     "98765",
		City:           "Flowertown",
		AdditionalInfo: "No notes",
	}
	err := handlersPsql.ValidateBusiness(incorrectBusiness)
	assert.ErrorContains(t, err, "invalid business id code")
}

func TestValidateBusinessWithInCorrectPostalCode(t *testing.T) {
	incorrectBusiness := database.Business{
		ID:             1,
		BusinessName:   "Test Business",
		BusinessIdCode: "1234567-8",
		Type:           "grower",
		PhoneNumber:    "010234567",
		Email:          "tester@test.fi",
		Address:        "Imaginary road 1",
		PostalCode:     "ABCDE",
		City:           "Flowertown",
		AdditionalInfo: "No notes",
	}
	err := handlersPsql.ValidateBusiness(incorrectBusiness)
	assert.ErrorContains(t, err, "invalid postal code")
}

func TestValidateBusinessWithEmptyField(t *testing.T) {
	incorrectBusiness := database.Business{
		ID:             1,
		BusinessName:   "Test Business",
		BusinessIdCode: "1234567-8",
		Type:           "",
		PhoneNumber:    "010234567",
		Email:          "tester@test.fi",
		Address:        "Imaginary road 1",
		PostalCode:     "98765",
		City:           "Flowertown",
		AdditionalInfo: "No notes",
	}
	err := handlersPsql.ValidateBusiness(incorrectBusiness)
	assert.ErrorContains(t, err, "all fields are required")
}

func TestValidateUserEmailWithCorrectInput(t *testing.T) {
	correctEmail := handlersPsql.UserEmail{
		UserEmail: "testuser@test.fi",
	}
	err := handlersPsql.ValidateUserEmail(correctEmail)
	assert.NoError(t, err, "ValidateUserMail() should not return an error")
}

func TestValidateUserEmailWithIncorrectEmail(t *testing.T) {
	correctEmail := handlersPsql.UserEmail{
		UserEmail: "testuser@testfi",
	}
	err := handlersPsql.ValidateUserEmail(correctEmail)
	assert.ErrorContains(t, err, "invalid user email")
}

func TestValidateUserEmailWithEmptyEmail(t *testing.T) {
	correctEmail := handlersPsql.UserEmail{
		UserEmail: "",
	}
	err := handlersPsql.ValidateUserEmail(correctEmail)
	assert.ErrorContains(t, err, "all fields are required")
}
