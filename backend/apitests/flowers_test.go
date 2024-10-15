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
}

func (s *FlowersAPITestSuite) SetupSuite() {
	s.TestFlowers = testdata.GetTestFlowers()
}

func (s *FlowersAPITestSuite) TestListingFlowersWithoutError() {
	testutils.RunTest(s.T(), testutils.TestCase{
		Description:  "\"GET /api/flowers\" without error",
		Route:        "/api/flowers",
		Method:       "GET",
		Body:         "",
		ExpectedCode: 200,
		ExpectedBody: utils.FlowersToJSON(s.TestFlowers),
		SetupMocks: func(db *mocks.Database) {
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
		Description:  "\"GET /api/flowers\" with error",
		Route:        "/api/flowers",
		Method:       "GET",
		Body:         "",
		ExpectedCode: 500,
		ExpectedBody: "Database error",
		SetupMocks: func(db *mocks.Database) {
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
		Description: "POST /api/flowers",
		Route:       "/api/flowers",
		Method:      "POST",
		Body: utils.FlowerToJSON(database.Flower{
			Name:      s.TestFlowers[0].Name,
			LatinName: s.TestFlowers[0].LatinName,
			Grower:    s.TestFlowers[0].Grower,
			Site:      s.TestFlowers[0].Site,
		}),
		ExpectedCode: 201,
		ExpectedBodyFunc: func(body string) {
			flower := database.Flower{}
			err := json.Unmarshal([]byte(body), &flower)
			s.NoError(err, "response body should include flower data: \""+body+"\"")
			s.Equal(flower.ID, s.TestFlowers[0].ID, "wrong ID in the added flower")
			s.Equal(flower.Name, s.TestFlowers[0].Name, "wrong Name in the added flower")
			s.Equal(flower.LatinName, s.TestFlowers[0].LatinName, "wrong LatinName in the added flower")
			s.Less(time.Since(flower.AddedTime).Seconds(), 10.0, "invalid AddedTime in the added flower")
			s.Equal(flower.Grower, s.TestFlowers[0].Grower, "wrong Grower in the added flower")
			s.Equal(flower.Site, s.TestFlowers[0].Site, "wrong Site in the added flower")
		},
		SetupMocks: func(db *mocks.Database) {
			db.EXPECT().AddFlower(
				mock.Anything, mock.Anything,
			).RunAndReturn(func(ctx context.Context, newFlower database.Flower) (*database.Flower, error) {
				return &database.Flower{
					ID:        s.TestFlowers[0].ID,
					Name:      newFlower.Name,
					LatinName: newFlower.LatinName,
					AddedTime: newFlower.AddedTime,
					Grower:    newFlower.Grower,
					Site:      newFlower.Site,
				}, nil
			}).Once()
			db.EXPECT().AddFlowerToSite(
				mock.Anything, mock.Anything, mock.Anything,
			).Return(
				nil,
			).Once()
		},
	})
}

func (s *FlowersAPITestSuite) TestDeletingFlower() {
	testutils.RunTest(s.T(), testutils.TestCase{
		Description:  "DELETE /api/flowers/<id>",
		Route:        "/api/flowers/" + s.TestFlowers[0].ID.Hex(),
		Method:       "DELETE",
		Body:         "",
		ExpectedCode: 204,
		ExpectedBody: "",
		SetupMocks: func(db *mocks.Database) {
			db.EXPECT().DeleteFlower(
				mock.Anything, s.TestFlowers[0].ID,
			).Return(
				true, nil,
			).Once()
		},
	})
}

func TestFlowersAPITestSuite(t *testing.T) {
	suite.Run(t, new(FlowersAPITestSuite))
}
