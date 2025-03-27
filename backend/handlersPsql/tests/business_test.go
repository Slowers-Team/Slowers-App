package tests

import (
	"testing"

	"github.com/stretchr/testify/suite"

	database "github.com/Slowers-team/Slowers-App/database/psql"
	"github.com/Slowers-team/Slowers-App/handlersPsql"
)

type BusinessHandlerTestSuite struct {
	suite.Suite
}

func (s *BusinessHandlerTestSuite) TestValidateBusinessWithCorrectInput() {
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
	s.NoError(
		err,
		"ValidateBusiness() should not return an error",
	)
}

func (s *BusinessHandlerTestSuite) TestValidateBusinessWithInCorrectBusinessEmail() {
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
	s.ErrorContains(
		err,
		"invalid business email",
	)
}

func (s *BusinessHandlerTestSuite) TestValidateBusinessWithInCorrectBusinessIdCode() {
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
	s.ErrorContains(
		err,
		"invalid business id code",
	)
}

func (s *BusinessHandlerTestSuite) TestValidateBusinessWithEmptyField() {
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
	s.ErrorContains(
		err,
		"all fields are required",
	)
}

func (s *BusinessHandlerTestSuite) TestValidateUserEmailWithCorrectInput() {
	correctEmail := handlersPsql.UserEmail{
		UserEmail: "testuser@test.fi",
	}
	err := handlersPsql.ValidateUserEmail(correctEmail)
	s.NoError(
		err,
		"ValidateUserMail() should not return an error",
	)
}

func (s *BusinessHandlerTestSuite) TestValidateUserEmailWithIncorrectEmail() {
	correctEmail := handlersPsql.UserEmail{
		UserEmail: "testuser@testfi",
	}
	err := handlersPsql.ValidateUserEmail(correctEmail)
	s.ErrorContains(
		err,
		"invalid user email",
	)
}

func (s *BusinessHandlerTestSuite) TestValidateUserEmailWithEmptyEmail() {
	correctEmail := handlersPsql.UserEmail{
		UserEmail: "",
	}
	err := handlersPsql.ValidateUserEmail(correctEmail)
	s.ErrorContains(
		err,
		"all fields are required",
	)
}

func TestDbBusinessTestSuite(t *testing.T) {
	suite.Run(t, new(BusinessHandlerTestSuite))
}
