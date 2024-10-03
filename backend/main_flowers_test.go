package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/Slowers-team/Slowers-App/database"
	sltest "github.com/Slowers-team/Slowers-App/testing"
	"github.com/Slowers-team/Slowers-App/utils"
)

func TestFlowersRoute(t *testing.T) {
	testFlowers := sltest.GetTestFlowers()
	testID := sltest.GetTestID()

	tests := []sltest.TestCase{
		{
			Description:   "\"GET /api/flowers\" without error",
			Route:         "/api/flowers",
			Method:        "GET",
			Body:          "",
			ExpectedError: false,
			ExpectedCode:  200,
			ExpectedBody:  utils.FlowersToJSON(testFlowers),
			SetupMocks:    func(db *database.MockDatabase) {
				db.On(
					"GetFlowers", mock.Anything,
				).Return(
					testFlowers, nil,
				).Once()
			},
		},
		{
			Description:   "\"GET /api/flowers\" with error",
			Route:         "/api/flowers",
			Method:        "GET",
			Body:          "",
			ExpectedError: false,
			ExpectedCode:  500,
			ExpectedBody:  "Database error",
			SetupMocks:    func(db *database.MockDatabase) {
				db.On(
					"GetFlowers", mock.Anything,
				).Return(
					[]database.Flower{}, errors.New("Database error"),
				).Once()
			},
		},
		{
			Description:   "DELETE /api/flowers/<id>",
			Route:         "/api/flowers/" + testID,
			Method:        "DELETE",
			Body:          utils.IDToJSON(testID),
			ExpectedError: false,
			ExpectedCode:  204,
			ExpectedBody:  "",
			SetupMocks:    func(db *database.MockDatabase) {
				db.On(
					"DeleteFlower", mock.Anything, testID,
				).Return(
					true, nil,
				).Once()
			},
		},

	}

	sltest.RunTests(t, tests)
}
