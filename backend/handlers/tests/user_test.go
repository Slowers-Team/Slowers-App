package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Slowers-team/Slowers-App/databases/sql"
	"github.com/Slowers-team/Slowers-App/handlers"
)

func TestRequireAllFieldsWithCorrectInput(t *testing.T) {
	correctUser := sql.User{
		Username: "Test user",
		Password: "testpassword",
		Email:    "tester@test.fi",
	}
	err := handlers.RequireAllFields(correctUser)
	assert.NoError(t, err, "ValidateUser() should not return an error")
}

func TestRequireAllFieldsWithIncorrectInput(t *testing.T) {
	incorrectUser := sql.User{
		Username: "Test user",
		Password: "testpassword",
		Email:    "",
	}
	err := handlers.RequireAllFields(incorrectUser)
	assert.ErrorContains(t, err, "all fields are required")
}
