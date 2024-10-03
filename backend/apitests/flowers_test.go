package apitests

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/Slowers-team/Slowers-App/database"
	"github.com/Slowers-team/Slowers-App/testdata"
	sltest "github.com/Slowers-team/Slowers-App/testing"
	"github.com/Slowers-team/Slowers-App/utils"
)

func TestListingFlowersWithoutError(t *testing.T) {
	sltest.RunTest(t, sltest.TestCase{
		Description:   "\"GET /api/flowers\" without error",
		Route:         "/api/flowers",
		Method:        "GET",
		Body:          "",
		ExpectedError: false,
		ExpectedCode:  200,
		ExpectedBody:  utils.FlowersToJSON(testdata.GetTestFlowers()),
		SetupMocks:    func(db *database.MockDatabase) {
			db.On(
				"GetFlowers", mock.Anything,
			).Return(
				testdata.GetTestFlowers(), nil,
			).Once()
		},
	})
}

func TestListingFlowersWithError(t *testing.T) {
	sltest.RunTest(t, sltest.TestCase{
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
	})
}

func TestDeletingFlower(t *testing.T) {
	sltest.RunTest(t, sltest.TestCase{
		Description:   "DELETE /api/flowers/<id>",
		Route:         "/api/flowers/" + testdata.GetTestID(),
		Method:        "DELETE",
		Body:          utils.IDToJSON(testdata.GetTestID()),
		ExpectedError: false,
		ExpectedCode:  204,
		ExpectedBody:  "",
		SetupMocks:    func(db *database.MockDatabase) {
			db.On(
				"DeleteFlower", mock.Anything, testdata.GetTestID(),
			).Return(
				true, nil,
			).Once()
		},
	})
}
