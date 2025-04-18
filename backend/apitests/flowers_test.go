//go:build api
// +build api

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
	Flowers []database.Flower
}

func (s *FlowersAPITestSuite) SetupSuite() {
	s.Flowers = testdata.GetFlowers()
}

func (s *FlowersAPITestSuite) TestListingFlowersWithoutError() {
	testutils.RunTest(s.T(), testutils.TestCase{
		Description:  "\"GET /api/flowers\" without error",
		Route:        "/api/flowers",
		Method:       "GET",
		ContentType:  "application/json",
		Body:         []byte{},
		ExpectedCode: 200,
		ExpectedBody: utils.ToJSON(s.Flowers),
		SetupMocks: func(db *mocks.Database) {
			db.EXPECT().GetFlowers(
				mock.Anything,
			).Return(
				s.Flowers, nil,
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
		Description:  "POST /api/flowers",
		Route:        "/api/flowers",
		Method:       "POST",
		ContentType:  "application/json",
		Body:         utils.ToJSON(testdata.PrepareFlowerForAdding(s.Flowers[0])),
		ExpectedCode: 201,
		ExpectedBodyFunc: func(body []byte) {
			flower := database.Flower{}
			err := json.Unmarshal(body, &flower)

			s.Require().NoError(
				err,
				"response body should include flower data: \""+string(body)+"\"",
			)
			s.Equal(
				s.Flowers[0].ID,
				flower.ID,
				"wrong ID in the added flower",
			)
			s.Equal(
				s.Flowers[0].Name,
				flower.Name,
				"wrong Name in the added flower",
			)
			s.Equal(
				s.Flowers[0].LatinName,
				flower.LatinName,
				"wrong LatinName in the added flower",
			)
			s.Less(
				time.Since(flower.AddedTime).Seconds(),
				10.0,
				"invalid AddedTime in the added flower",
			)
			s.Equal(
				s.Flowers[0].Grower,
				flower.Grower,
				"wrong Grower in the added flower",
			)
			s.Equal(
				s.Flowers[0].Site,
				flower.Site,
				"wrong Site in the added flower",
			)
			s.Equal(
				s.Flowers[0].Quantity,
				flower.Quantity,
				"wrong Quantity in the added flower",
			)
		},
		SetupMocks: func(db *mocks.Database) {
			user := testdata.GetUsers()[0]
			db.EXPECT().GetUserByID(
				mock.Anything, *s.Flowers[0].Grower,
			).Return(
				&user, nil,
			).Once()
			sites := testdata.GetRootSites()
			db.EXPECT().GetSiteByID(
				mock.Anything, sites[0].ID,
			).Return(
				&sites[0], nil,
			).Once()

			db.EXPECT().AddFlower(
				mock.Anything, mock.Anything,
			).RunAndReturn(func(ctx context.Context, newFlower database.Flower) (*database.Flower, error) {
				return &database.Flower{
					ID:        s.Flowers[0].ID,
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
		Route:        "/api/flowers/" + s.Flowers[0].ID.Hex(),
		Method:       "DELETE",
		ContentType:  "application/json",
		Body:         []byte{},
		ExpectedCode: 204,
		ExpectedBody: []byte{},
		SetupMocks: func(db *mocks.Database) {
			db.EXPECT().DeleteFlower(
				mock.Anything, s.Flowers[0].ID,
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
		ExpectedBody: utils.ToJSON(s.Flowers),
		SetupMocks: func(db *mocks.Database) {
			db.EXPECT().GetUserFlowers(
				mock.Anything, testdata.GetUsers()[0].ID,
			).Return(
				s.Flowers, nil,
			).Once()
		},
	})
}

func (s *FlowersAPITestSuite) TestListingFlowersOfSite() {
	site := testdata.GetRootSites()[0]
	user := testdata.GetUsers()[0]
	flowers := []database.Flower{s.Flowers[0]}

	testutils.RunTest(s.T(), testutils.TestCase{
		Description:  "GET /api/sites/<id>/flowers",
		Route:        "/api/sites/" + site.ID.Hex() + "/flowers",
		Method:       "GET",
		ContentType:  "application/json",
		Body:         []byte{},
		ExpectedCode: 200,
		ExpectedBody: utils.ToJSON(flowers),
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
	flower := s.Flowers[0]
	modifiedValues := database.Flower{
		Name:      "modified name",
		LatinName: "modified latin name",
		Quantity:  flower.Quantity + 1,
	}

	modifiedFlower := flower
	modifiedFlower.Name = modifiedValues.Name
	modifiedFlower.LatinName = modifiedValues.LatinName
	modifiedFlower.Quantity = modifiedValues.Quantity

	testutils.RunTest(s.T(), testutils.TestCase{
		Description:  "PUT /api/flowers/<id>",
		Route:        "/api/flowers/" + flower.ID.Hex(),
		Method:       "PUT",
		ContentType:  "application/json",
		Body:         utils.ToJSON(modifiedValues),
		ExpectedCode: 200,
		ExpectedBody: utils.ToJSON(modifiedFlower),
		SetupMocks: func(db *mocks.Database) {
			db.EXPECT().ModifyFlower(
				mock.Anything, flower.ID, modifiedValues,
			).Return(
				&modifiedFlower, nil,
			).Once()
		},
	})
}

func (s *FlowersAPITestSuite) TestDeletingMultipleFlowers() {
	var flowerIDs []string
	var ids []database.ObjectID
	for _, flower := range s.Flowers {
		flowerIDs = append(flowerIDs, flower.ID.Hex())
		ids = append(ids, flower.ID)
	}

	testutils.RunTest(s.T(), testutils.TestCase{
		Description:  "POST /api/flowers/delete-multiple",
		Route:        "/api/flowers/delete-multiple",
		Method:       "POST",
		ContentType:  "application/json",
		Body:         utils.ToJSON(flowerIDs),
		ExpectedCode: 204,
		ExpectedBody: []byte{},
		SetupMocks: func(db *mocks.Database) {
			db.EXPECT().DeleteMultipleFlowers(
				mock.Anything, ids,
			).Return(
				nil,
			).Once()
		},
	})
}

func TestFlowersAPITestSuite(t *testing.T) {
	suite.Run(t, new(FlowersAPITestSuite))
}
