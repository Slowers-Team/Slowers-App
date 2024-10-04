package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/Slowers-team/Slowers-App/database"
	"github.com/Slowers-team/Slowers-App/testdata"
	"github.com/Slowers-team/Slowers-App/testutils"
	"github.com/Slowers-team/Slowers-App/utils"
)

type DbSiteTestSuite struct {
	suite.Suite
	Db database.Database
	TestSite database.Site
}

func (s *DbSiteTestSuite) SetupSuite() {
	s.Db = testutils.ConnectDB()
	s.Db.Clear()
	s.TestSite = testdata.GetRootSites()[0]
}

func (s *DbSiteTestSuite) TestAddSite() {
	site := database.Site{
		Name: s.TestSite.Name,
		AddedTime: s.TestSite.AddedTime,
		Note: s.TestSite.Note,
		Parent: s.TestSite.Parent,
		Flowers: s.TestSite.Flowers,
		Owner: s.TestSite.Owner,
	}
	createdSite, err := s.Db.AddSite(context.Background(), site)

	s.NoError(
		err,
		"AddSite() should not return an error",
	)
	s.NotZero(
		createdSite.ID,
		"ID for the created site should be non-zero",
	)
	s.Equal(
		createdSite.Name,
		site.Name,
		"wrong name for the site returned from AddSite()",
	)
	s.Equal(
		createdSite.AddedTime,
		site.AddedTime,
		"wrong AddedTime for the site returned from AddSite()",
	)
	s.Equal(
		createdSite.Note,
		site.Note,
		"wrong note for the site returned from AddSite()",
	)
	s.Equal(
		createdSite.Parent,
		site.Parent,
		"wrong parent for the site returned from AddSite()",
	)
	s.True(
		utils.AreIDPtrSlicesEql(createdSite.Flowers, site.Flowers),
		"wrong flowers for the site returned from AddSite()",
	)
	s.Equal(
		createdSite.Owner,
		site.Owner,
		"wrong owner for the site returned from AddSite()",
	)
}

func (s *DbSiteTestSuite) TestAddAndGetRootSites() {
	site := database.Site{
		Name: s.TestSite.Name,
		AddedTime: s.TestSite.AddedTime,
		Note: s.TestSite.Note,
		Parent: s.TestSite.Parent,
		Flowers: s.TestSite.Flowers,
		Owner: s.TestSite.Owner,
	}
	createdSite, _ := s.Db.AddSite(context.Background(), site)
	rootSites, err := s.Db.GetRootSites(context.Background())

	s.NoError(
		err,
		"GetRootSites() should not return an error",
	)
	s.Len(
		rootSites,
		1,
		"GetRootSites() should return a slice of length 1",
	)
	s.Equal(
		rootSites[0].ID,
		createdSite.ID,
		"ID for the fetched site should be non-zero",
	)
	s.Equal(
		rootSites[0].Name,
		site.Name,
		"wrong name for the site returned from GetRootSites()",
	)
	s.Equal(
		rootSites[0].AddedTime,
		site.AddedTime,
		"wrong AddedTime for the site returned from GetRootSites()",
	)
	s.Equal(
		rootSites[0].Note,
		site.Note,
		"wrong note for the site returned from GetRootSites()",
	)
	s.Equal(
		rootSites[0].Parent,
		site.Parent,
		"wrong parent for the site returned from GetRootSites()",
	)
	s.True(
		utils.AreIDPtrSlicesEql(rootSites[0].Flowers, site.Flowers),
		"wrong flowers for the site returned from GetRootSites()",
	)
	s.Equal(
		rootSites[0].Owner,
		site.Owner,
		"wrong owner for the site returned from GetRootSites()",
	)
}

func (s *DbSiteTestSuite) TearDownTest() {
	s.Db.Clear()
}

func (s *DbSiteTestSuite) TearDownSuite() {
	testutils.DisconnectDB(s.Db)
}

func TestDbSiteTestSuite(t *testing.T) {
	suite.Run(t, new(DbSiteTestSuite))
}
