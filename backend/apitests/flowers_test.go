package apitests

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/Slowers-team/Slowers-App/database"
	"github.com/Slowers-team/Slowers-App/mocks"
	"github.com/Slowers-team/Slowers-App/testdata"
	"github.com/Slowers-team/Slowers-App/testutils"
	"github.com/Slowers-team/Slowers-App/utils"
)

func TestListingFlowersWithoutError(t *testing.T) {
	testutils.RunTest(t, testutils.TestCase{
		Description:   "\"GET /api/flowers\" without error",
		Route:         "/api/flowers",
		Method:        "GET",
		Body:          "",
		ExpectedError: false,
		ExpectedCode:  200,
		ExpectedBody:  utils.FlowersToJSON(testdata.GetTestFlowers()),
		SetupMocks:    func(db *mocks.Database) {
			db.On(
				"GetFlowers", mock.Anything,
			).Return(
				testdata.GetTestFlowers(), nil,
			).Once()
		},
	})
}

func TestListingFlowersWithError(t *testing.T) {
	testutils.RunTest(t, testutils.TestCase{
		Description:   "\"GET /api/flowers\" with error",
		Route:         "/api/flowers",
		Method:        "GET",
		Body:          "",
		ExpectedError: false,
		ExpectedCode:  500,
		ExpectedBody:  "Database error",
		SetupMocks:    func(db *mocks.Database) {
			db.On(
				"GetFlowers", mock.Anything,
			).Return(
				[]database.Flower{}, errors.New("Database error"),
			).Once()
		},
	})
}

func TestDeletingFlower(t *testing.T) {
	testutils.RunTest(t, testutils.TestCase{
		Description:   "DELETE /api/flowers/<id>",
		Route:         "/api/flowers/" + testdata.GetTestID(),
		Method:        "DELETE",
		Body:          utils.IDToJSON(testdata.GetTestID()),
		ExpectedError: false,
		ExpectedCode:  204,
		ExpectedBody:  "",
		SetupMocks:    func(db *mocks.Database) {
			db.On(
				"DeleteFlower", mock.Anything, testdata.GetTestID(),
			).Return(
				true, nil,
			).Once()
		},
	})
}
