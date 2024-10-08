package apitests

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Slowers-team/Slowers-App/database"
	"github.com/Slowers-team/Slowers-App/mocks"
	"github.com/Slowers-team/Slowers-App/testdata"
	"github.com/Slowers-team/Slowers-App/testutils"
	"github.com/Slowers-team/Slowers-App/utils"
)

type SitesAPITestSuite struct {
	suite.Suite
	TestFlowers []database.Flower
	RootSites   []database.Site
	TestUser    database.User
}

func (s *SitesAPITestSuite) SetupSuite() {
	s.TestFlowers = testdata.GetTestFlowers()
	s.RootSites = testdata.GetRootSites()
	s.TestUser = testdata.GetUser()
}

func (s *SitesAPITestSuite) TestListingRootSites() {
	testutils.RunTest(s.T(), testutils.TestCase{
		Description:  "GET /api/sites",
		Route:        "/api/sites",
		Method:       "GET",
		Body:         "",
		ExpectedCode: 200,
		ExpectedBody: utils.SitesToJSON(s.RootSites),
		SetupMocks: func(db *mocks.Database) {
			db.EXPECT().GetRootSites(
				mock.Anything, s.TestUser.ID,
			).Return(
				s.RootSites, nil,
			).Once()
		},
	})
}

func (s *SitesAPITestSuite) TestFetchingSite() {
	testutils.RunTest(s.T(), testutils.TestCase{
		Description:  "GET /api/sites/<id>",
		Route:        "/api/sites/" + s.RootSites[0].ID.Hex(),
		Method:       "GET",
		Body:         "",
		ExpectedCode: 200,
		ExpectedBody: utils.SiteDataToJSON(testdata.GetSite()),
		SetupMocks: func(db *mocks.Database) {
			db.EXPECT().GetSite(
				mock.Anything, s.RootSites[0].ID.Hex(), s.TestUser.ID,
			).Return(
				testdata.GetSite(), nil,
			).Once()
		},
	})
}

func (s *SitesAPITestSuite) TestAddingSite() {
	testutils.RunTest(s.T(), testutils.TestCase{
		Description: "POST /api/sites",
		Route:       "/api/sites",
		Method:      "POST",
		Body: utils.SiteToJSON(database.Site{
			Flowers: s.RootSites[0].Flowers,
			Name:    s.RootSites[0].Name,
			Note:    s.RootSites[0].Note,
			Owner:   s.RootSites[0].Owner,
			Parent:  s.RootSites[0].Parent,
		}),
		ExpectedCode: 201,
		ExpectedBodyFunc: func(body string) {
			site := database.Site{}
			json.Unmarshal([]byte(body), &site)
			s.Equal(
				site.ID.Hex(),
				s.RootSites[0].ID.Hex(),
				"added site has wrong ID",
			)
			s.Less(
				time.Since(site.AddedTime).Seconds(),
				10.0,
				"added site has invalid AddedTime",
			)
			s.True(
				utils.AreIDPtrSlicesEql(site.Flowers, s.RootSites[0].Flowers),
				"added site has wrong flowers",
			)
			s.Equal(
				site.Name,
				s.RootSites[0].Name,
				"added site has wrong name",
			)
			s.Equal(
				site.Note,
				s.RootSites[0].Note,
				"added site has wrong note",
			)
			s.Equal(
				site.Owner,
				s.RootSites[0].Owner,
				"added site has wrong owner",
			)
			s.Equal(
				site.Parent,
				s.RootSites[0].Parent,
				"added site has wrong parent",
			)
		},
		SetupMocks: func(db *mocks.Database) {
			db.EXPECT().AddSite(
				mock.Anything, mock.Anything,
			).RunAndReturn(func(ctx context.Context, newSite database.Site) (*database.Site, error) {
				return &database.Site{
					ID:        s.RootSites[0].ID,
					AddedTime: newSite.AddedTime,
					Flowers:   newSite.Flowers,
					Name:      newSite.Name,
					Note:      newSite.Note,
					Owner:     newSite.Owner,
					Parent:    newSite.Parent,
				}, nil
			}).Once()
		},
	})
}

func (s *SitesAPITestSuite) TestDeletingSite() {
	testutils.RunTest(s.T(), testutils.TestCase{
		Description:  "DELETE /api/sites/<id>",
		Route:        "/api/sites/" + s.RootSites[0].ID.Hex(),
		Method:       "DELETE",
		Body:         "",
		ExpectedCode: 200,
		ExpectedBody: "{\"DeletedCount\":1}",
		SetupMocks: func(db *mocks.Database) {
			db.EXPECT().DeleteSite(
				mock.Anything, s.RootSites[0].ID.Hex(), s.TestUser.ID,
			).Return(
				&mongo.DeleteResult{DeletedCount: 1}, nil,
			).Once()
		},
	})
}

func TestSitesAPITestSuite(t *testing.T) {
	suite.Run(t, new(SitesAPITestSuite))
}
