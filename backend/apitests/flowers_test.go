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
		ContentType:  "application/json",
		Body:         []byte{},
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
		ContentType:  "application/json",
		Body:         []byte{},
		ExpectedCode: 500,
		ExpectedBody: []byte("Database error"),
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
		ContentType: "application/json",
		Body: utils.FlowerToJSON(database.Flower{
			Name:      s.TestFlowers[0].Name,
			LatinName: s.TestFlowers[0].LatinName,
			Grower:    s.TestFlowers[0].Grower,
			Site:      s.TestFlowers[0].Site,
			Quantity:  s.TestFlowers[0].Quantity,
		}),
		ExpectedCode: 201,
		ExpectedBodyFunc: func(body []byte) {
			flower := database.Flower{}
			err := json.Unmarshal(body, &flower)
			s.NoError(err, "response body should include flower data: \""+string(body)+"\"")
			s.Equal(flower.ID, s.TestFlowers[0].ID, "wrong ID in the added flower")
			s.Equal(flower.Name, s.TestFlowers[0].Name, "wrong Name in the added flower")
			s.Equal(flower.LatinName, s.TestFlowers[0].LatinName, "wrong LatinName in the added flower")
			s.Less(time.Since(flower.AddedTime).Seconds(), 10.0, "invalid AddedTime in the added flower")
			s.Equal(flower.Grower, s.TestFlowers[0].Grower, "wrong Grower in the added flower")
			s.Equal(flower.Site, s.TestFlowers[0].Site, "wrong Site in the added flower")
			s.Equal(flower.Quantity, s.TestFlowers[0].Quantity, "wrong Quantity in the added flower")
		},
		SetupMocks: func(db *mocks.Database) {
			user := testdata.GetUsers()[0]
			db.EXPECT().GetUserByID(mock.Anything, *s.TestFlowers[0].Grower).Return(&user, nil).Once()
			sites := testdata.GetRootSites()
			db.EXPECT().GetSiteByID(mock.Anything, sites[0].ID).Return(&sites[0], nil).Once()

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
					Quantity:  newFlower.Quantity,
					Visible:   newFlower.Visible,
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
		ContentType:  "application/json",
		Body:         []byte{},
		ExpectedCode: 204,
		ExpectedBody: []byte{},
		SetupMocks: func(db *mocks.Database) {
			db.EXPECT().DeleteFlower(
				mock.Anything, s.TestFlowers[0].ID,
			).Return(
				true, nil,
			).Once()
		},
	})
}

func (s *FlowersAPITestSuite) TestListingFlowersOfCurrentUser() {
	testutils.RunTest(s.T(), testutils.TestCase{
		Description:  "GET /api/flowers/user",
		Route:        "/api/flowers/user",
		Method:       "GET",
		ContentType:  "application/json",
		Body:         []byte{},
		ExpectedCode: 200,
		ExpectedBody: utils.FlowersToJSON(s.TestFlowers),
		SetupMocks: func(db *mocks.Database) {
			db.EXPECT().GetUserFlowers(
				mock.Anything, testdata.GetUsers()[0].ID,
			).Return(
				s.TestFlowers, nil,
			).Once()
		},
	})
}

func (s *FlowersAPITestSuite) TestListingFlowersOfSite() {
	site := testdata.GetRootSites()[0]
	user := testdata.GetUsers()[0]
	flowers := []database.Flower{s.TestFlowers[0]}

	testutils.RunTest(s.T(), testutils.TestCase{
		Description:  "GET /api/sites/<id>/flowers",
		Route:        "/api/sites/" + site.ID.Hex() + "/flowers",
		Method:       "GET",
		ContentType:  "application/json",
		Body:         []byte{},
		ExpectedCode: 200,
		ExpectedBody: utils.FlowersToJSON(flowers),
		SetupMocks: func(db *mocks.Database) {
			db.EXPECT().GetAllFlowersRelatedToSite(
				mock.Anything, site.ID, user.ID,
			).Return(
				flowers, nil,
			).Once()
		},
	})
}

func (s *FlowersAPITestSuite) TestModifyingFlower() {
	flower := s.TestFlowers[0]
	modifiedFlower := database.Flower{
		Name:      "modified name",
		LatinName: "modified latin name",
		Quantity:  10,
	}

	testutils.RunTest(s.T(), testutils.TestCase{
		Description:  "PUT /api/flowers/<id>",
		Route:        "/api/flowers/" + s.TestFlowers[0].ID.Hex(),
		Method:       "PUT",
		ContentType:  "application/json",
		Body:         []byte(utils.FlowerToJSON(modifiedFlower)),
		ExpectedCode: 200,
		ExpectedBody: utils.FlowerToJSON(flower),
		SetupMocks: func(db *mocks.Database) {
			db.EXPECT().ModifyFlower(
				mock.Anything, s.TestFlowers[0].ID, modifiedFlower,
			).Return(
				&flower, nil,
			).Once()
		},
	})
}

func TestFlowersAPITestSuite(t *testing.T) {
	suite.Run(t, new(FlowersAPITestSuite))
}
