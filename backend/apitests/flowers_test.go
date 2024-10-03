package apitests

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/Slowers-team/Slowers-App/database"
	"github.com/Slowers-team/Slowers-App/mocks"
	"github.com/Slowers-team/Slowers-App/testdata"
	"github.com/Slowers-team/Slowers-App/testutils"
	"github.com/Slowers-team/Slowers-App/utils"
)

type FlowersAPITestSuite struct {
	suite.Suite
	TestFlowers []database.Flower
	TestFlowersConcise []database.Flower
}

func (s *FlowersAPITestSuite) SetupSuite() {
	s.TestFlowers = testdata.GetTestFlowers()
	s.TestFlowersConcise = testdata.GetTestFlowersConcise()
}

func (s *FlowersAPITestSuite) TestListingFlowersWithoutError(t *testing.T) {
	testutils.RunTest(t, testutils.TestCase{
		Description:   "\"GET /api/flowers\" without error",
		Route:         "/api/flowers",
		Method:        "GET",
		Body:          "",
		ExpectedError: false,
		ExpectedCode:  200,
		ExpectedBody:  utils.FlowersToJSON(s.TestFlowers),
		SetupMocks:    func(db *mocks.Database) {
			db.On(
				"GetFlowers", mock.Anything,
			).Return(
				s.TestFlowers, nil,
			).Once()
		},
	})
}

func (s *FlowersAPITestSuite) TestListingFlowersWithError(t *testing.T) {
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

func (s *FlowersAPITestSuite) TestDeletingFlower(t *testing.T) {
	testutils.RunTest(t, testutils.TestCase{
		Description:   "DELETE /api/flowers/<id>",
		Route:         "/api/flowers/" + s.TestFlowers[0].ID.String(),
		Method:        "DELETE",
		Body:          "",
		ExpectedError: false,
		ExpectedCode:  204,
		ExpectedBody:  "",
		SetupMocks:    func(db *mocks.Database) {
			db.On(
				"DeleteFlower", mock.Anything, s.TestFlowers[0].ID.String(),
			).Return(
				true, nil,
			).Once()
		},
	})
}
