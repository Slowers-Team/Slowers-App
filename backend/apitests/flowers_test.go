package apitests

import (
	"context"
	"encoding/json"
	"errors"
	"testing"
	"time"

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

func (s *FlowersAPITestSuite) TestListingFlowersWithoutError() {
	testutils.RunTest(s.T(), testutils.TestCase{
		Description:   "\"GET /api/flowers\" without error",
		Route:         "/api/flowers",
		Method:        "GET",
		Body:          "",
		ExpectedError: false,
		ExpectedCode:  200,
		ExpectedBody:  utils.FlowersToJSON(s.TestFlowers),
		SetupMocks:    func(db *mocks.Database) {
			db.EXPECT().GetFlowers(
				mock.Anything,
			).Return(
				s.TestFlowers, nil,
			).Once()
		},
	})
}

func (s *FlowersAPITestSuite) TestListingFlowersWithError() {
	testutils.RunTest(s.T(), testutils.TestCase{
		Description:   "\"GET /api/flowers\" with error",
		Route:         "/api/flowers",
		Method:        "GET",
		Body:          "",
		ExpectedError: false,
		ExpectedCode:  500,
		ExpectedBody:  "Database error",
		SetupMocks:    func(db *mocks.Database) {
			db.EXPECT().GetFlowers(
				mock.Anything,
			).Return(
				[]database.Flower{}, errors.New("Database error"),
			).Once()
		},
	})
}

func (s *FlowersAPITestSuite) TestAddingFlower() {
	testutils.RunTest(s.T(), testutils.TestCase{
		Description:      "POST /api/flowers",
		Route:            "/api/flowers",
		Method:           "POST",
		Body:             utils.FlowerToJSON(s.TestFlowersConcise[0]),
		ExpectedError:    false,
		ExpectedCode:     201,
		ExpectedBodyFunc: func(body string) bool {
			flower := database.Flower{}
			json.Unmarshal([]byte(body), &flower)
			return flower.ID == s.TestFlowers[0].ID &&
				flower.Name == s.TestFlowers[0].Name &&
				flower.LatinName == s.TestFlowers[0].LatinName &&
				time.Since(flower.AddedTime).Seconds() < 10.0
		},
		SetupMocks:       func(db *mocks.Database) {
			db.EXPECT().AddFlower(
				mock.Anything, mock.Anything,
			).RunAndReturn(func(ctx context.Context, newFlower database.Flower) (*database.Flower, error) {
				return &database.Flower{
					ID: s.TestFlowers[0].ID,
					Name: newFlower.Name,
					LatinName: newFlower.LatinName,
					AddedTime: newFlower.AddedTime,
				}, nil
			}).Once()
		},
	})
}

func (s *FlowersAPITestSuite) TestDeletingFlower() {
	testutils.RunTest(s.T(), testutils.TestCase{
		Description:   "DELETE /api/flowers/<id>",
		Route:         "/api/flowers/" + s.TestFlowers[0].ID.Hex(),
		Method:        "DELETE",
		Body:          "",
		ExpectedError: false,
		ExpectedCode:  204,
		ExpectedBody:  "",
		SetupMocks:    func(db *mocks.Database) {
			db.EXPECT().DeleteFlower(
				mock.Anything, s.TestFlowers[0].ID.Hex(),
			).Return(
				true, nil,
			).Once()
		},
	})
}

func TestFlowersAPITestSuite(t *testing.T) {
	suite.Run(t, new(FlowersAPITestSuite))
}
