package tests

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Slowers-team/Slowers-App/database"
	"github.com/Slowers-team/Slowers-App/testdata"
	"github.com/Slowers-team/Slowers-App/testutils"
)

type DbSiteTestSuite struct {
	suite.Suite
	Db   database.Database
	Site database.Site
	User database.User
}

func (s *DbSiteTestSuite) SetupSuite() {
	s.Db = testutils.ConnectDB()
	s.Db.Clear()
	s.Site = testdata.GetRootSites()[0]
	s.User = testdata.GetUsers()[0]
}

func (s *DbSiteTestSuite) TestAddSite() {
	site := s.Site
	siteToAdd := testdata.PrepareSiteForAdding(site)
	addedSite, err := s.Db.AddSite(context.Background(), siteToAdd)

	s.Require().NoError(
		err,
		"AddSite() should not return an error",
	)
	s.NotZero(
		addedSite.ID,
		"ID for the created site should be non-zero",
	)
	s.Equal(
		site.Name,
		addedSite.Name,
		"wrong name for the site returned from AddSite()",
	)
	s.Equal(
		site.AddedTime,
		addedSite.AddedTime,
		"wrong AddedTime for the site returned from AddSite()",
	)
	s.Equal(
		site.Note,
		addedSite.Note,
		"wrong note for the site returned from AddSite()",
	)
	s.Equal(
		site.Parent,
		addedSite.Parent,
		"wrong parent for the site returned from AddSite()",
	)
	s.Equal(
		site.Flowers,
		addedSite.Flowers,
		"wrong flowers for the site returned from AddSite()",
	)
	s.Equal(
		site.Owner,
		addedSite.Owner,
		"wrong owner for the site returned from AddSite()",
	)
}

func (s *DbSiteTestSuite) TestAddAndGetRootSites() {
	site := s.Site
	siteToAdd := testdata.PrepareSiteForAdding(site)
	addedSite, _ := s.Db.AddSite(context.Background(), siteToAdd)
	fetchedSites, err := s.Db.GetRootSites(context.Background(), s.User.ID)

	s.Require().NoError(
		err,
		"GetRootSites() should not return an error",
	)
	s.Require().Len(
		fetchedSites,
		1,
		"GetRootSites() should return a slice of length 1",
	)
	s.Equal(
		addedSite.ID,
		fetchedSites[0].ID,
		"ID for the fetched site should be non-zero",
	)
	s.Equal(
		site.Name,
		fetchedSites[0].Name,
		"wrong name for the site returned from GetRootSites()",
	)
	s.Equal(
		site.AddedTime,
		fetchedSites[0].AddedTime,
		"wrong AddedTime for the site returned from GetRootSites()",
	)
	s.Equal(
		site.Note,
		fetchedSites[0].Note,
		"wrong note for the site returned from GetRootSites()",
	)
	s.Equal(
		site.Parent,
		fetchedSites[0].Parent,
		"wrong parent for the site returned from GetRootSites()",
	)
	s.Equal(
		site.Flowers,
		fetchedSites[0].Flowers,
		"wrong flowers for the site returned from GetRootSites()",
	)
	s.Equal(
		site.Owner,
		fetchedSites[0].Owner,
		"wrong owner for the site returned from GetRootSites()",
	)
}

func (s *DbSiteTestSuite) TestAddAndGetSite() {
	siteData := testdata.GetSite()

	site := siteData["site"].(database.Site)
	siteToAdd := testdata.PrepareSiteForAdding(site)
	addedSite, _ := s.Db.AddSite(context.Background(), siteToAdd)

	subSiteBson := siteData["subsites"].([]bson.M)[0]
	subSite := database.Site{
		Name:      subSiteBson["name"].(string),
		AddedTime: time.Date(2024, 9, 19, 12, 11, 4, 0, time.UTC),
		Note:      subSiteBson["note"].(string),
		Parent:    &addedSite.ID,
		Flowers:   []*database.ObjectID{},
		Owner:     site.Owner,
	}
	addedSubSite, _ := s.Db.AddSite(context.Background(), subSite)

	site.ID = addedSite.ID
	siteData["site"] = site
	subSiteBson["_id"] = addedSubSite.ID.Hex()
	siteData["subsites"] = []bson.M{subSiteBson}

	fetchedSiteData, err := s.Db.GetSite(
		context.Background(), addedSite.ID, *site.Owner,
	)

	s.Require().NoError(
		err,
		"GetSite() should not return an error",
	)

	doc, err := bson.Marshal(fetchedSiteData["site"].(bson.M))
	if err != nil {
		log.Fatal(err)
	}

	var fetchedSite database.Site
	err = bson.Unmarshal(doc, &fetchedSite)
	if err != nil {
		log.Fatal(err)
	}

	fetchedSubSites := fetchedSiteData["subsites"].([]bson.M)

	s.Equal(
		addedSite.ID,
		fetchedSite.ID,
		"wrong ID for the site returned from GetSite()",
	)
	s.Equal(
		site.Name,
		fetchedSite.Name,
		"wrong name for the site returned from GetSite()",
	)
	s.Equal(
		site.AddedTime,
		fetchedSite.AddedTime,
		"wrong AddedTime for the site returned from GetSite()",
	)
	s.Equal(
		site.Note,
		fetchedSite.Note,
		"wrong note for the site returned from GetSite()",
	)
	s.Equal(
		site.Parent,
		fetchedSite.Parent,
		"wrong parent for the site returned from GetSite()",
	)
	s.Equal(
		site.Flowers,
		fetchedSite.Flowers,
		"wrong flowers for the site returned from GetSite()",
	)
	s.Equal(
		site.Owner,
		fetchedSite.Owner,
		"wrong owner for the site returned from GetSite()",
	)

	s.Require().Len(
		fetchedSubSites,
		1,
		"GetSite() should return exactly one subsite",
	)

	s.Equal(
		addedSubSite.ID,
		fetchedSubSites[0]["_id"],
		"wrong ID for the subsite returned from GetSite()",
	)
	s.Equal(
		subSite.Name,
		fetchedSubSites[0]["name"],
		"wrong name for the subsite returned from GetSite()",
	)
	s.Equal(
		subSite.Note,
		fetchedSubSites[0]["note"],
		"wrong note for the subsite returned from GetSite()",
	)
}

func (s *DbSiteTestSuite) TestAddAndGetSiteByID() {
	site := s.Site
	siteToAdd := testdata.PrepareSiteForAdding(site)
	addedSite, _ := s.Db.AddSite(context.Background(), siteToAdd)

	fetchedSite, err := s.Db.GetSiteByID(context.Background(), addedSite.ID)

	s.Require().NoError(
		err,
		"GetSiteByID() should not return an error",
	)

	s.Equal(
		addedSite.ID,
		fetchedSite.ID,
		"wrong ID for the site returned from GetSiteByID()",
	)
	s.Equal(
		site.Name,
		fetchedSite.Name,
		"wrong name for the site returned from GetSiteByID()",
	)
	s.Equal(
		site.AddedTime,
		fetchedSite.AddedTime,
		"wrong AddedTime for the site returned from GetSiteByID()",
	)
	s.Equal(
		site.Note,
		fetchedSite.Note,
		"wrong note for the site returned from GetSiteByID()",
	)
	s.Equal(
		site.Parent,
		fetchedSite.Parent,
		"wrong parent for the site returned from GetSiteByID()",
	)
	s.Equal(
		site.Flowers,
		fetchedSite.Flowers,
		"wrong flowers for the site returned from GetSiteByID()",
	)
	s.Equal(
		site.Owner,
		fetchedSite.Owner,
		"wrong owner for the site returned from GetSiteByID()",
	)
}

func (s *DbSiteTestSuite) TestAddFlowerToSite() {
	site := s.Site
	site.Flowers = []*database.ObjectID{}
	siteToAdd := testdata.PrepareSiteForAdding(site)
	addedSite, _ := s.Db.AddSite(context.Background(), siteToAdd)

	flower := testdata.GetFlowers()[0]
	flower.Site = &addedSite.ID
	flowerToAdd := testdata.PrepareFlowerForAdding(flower)
	addedFlower, _ := s.Db.AddFlower(context.Background(), flowerToAdd)
	err := s.Db.AddFlowerToSite(
		context.Background(), addedSite.ID, addedFlower.ID,
	)

	s.Require().NoError(
		err,
		"AddFlowerToSite() should not return an error",
	)

	fetchedSite, _ := s.Db.GetSiteByID(context.Background(), addedSite.ID)

	s.Require().Len(
		fetchedSite.Flowers,
		1,
		"Site should contain exactly one added flower",
	)

	s.Equal(
		addedFlower.ID,
		*fetchedSite.Flowers[0],
		"Added flower has wrong ID",
	)
}

func (s *DbSiteTestSuite) TestAddAndDeleteSite() {
	siteData := testdata.GetSite()

	site := siteData["site"].(database.Site)
	site.Flowers = []*database.ObjectID{}
	siteToAdd := testdata.PrepareSiteForAdding(site)
	addedSite, _ := s.Db.AddSite(context.Background(), siteToAdd)

	subSiteBson := siteData["subsites"].([]bson.M)[0]
	subSite := database.Site{
		Name:      subSiteBson["name"].(string),
		AddedTime: time.Date(2024, 9, 19, 12, 11, 4, 0, time.UTC),
		Note:      subSiteBson["note"].(string),
		Parent:    &addedSite.ID,
		Flowers:   []*database.ObjectID{},
		Owner:     site.Owner,
	}
	addedSubSite, _ := s.Db.AddSite(context.Background(), subSite)

	site.ID = addedSite.ID
	siteData["site"] = site
	subSiteBson["_id"] = addedSubSite.ID.Hex()
	siteData["subsites"] = []bson.M{subSiteBson}

	flower := testdata.GetFlowers()[0]
	flower.Site = &addedSite.ID
	flowerToAdd := testdata.PrepareFlowerForAdding(flower)
	addedFlower, _ := s.Db.AddFlower(context.Background(), flowerToAdd)
	s.Db.AddFlowerToSite(context.Background(), addedSite.ID, addedFlower.ID)

	site2 := testdata.GetRootSitesForUser2()[0]
	site2.Flowers = []*database.ObjectID{}
	site2.Owner = site.Owner
	siteToAdd2 := testdata.PrepareSiteForAdding(site2)
	addedSite2, _ := s.Db.AddSite(context.Background(), siteToAdd2)

	flower2 := testdata.GetFlowerForUser2()
	flower2.Site = &addedSite2.ID
	flowerToAdd2 := testdata.PrepareFlowerForAdding(flower2)
	addedFlower2, _ := s.Db.AddFlower(context.Background(), flowerToAdd2)
	s.Db.AddFlowerToSite(context.Background(), addedSite2.ID, addedFlower2.ID)

	deleteResult, err := s.Db.DeleteSite(
		context.Background(), addedSite.ID, *site.Owner,
	)

	s.Require().NoError(
		err,
		"DeleteSite() should not return an error",
	)
	s.EqualValues(
		2,
		deleteResult.DeletedCount,
		"DeleteSite() should delete exactly two sites",
	)

	fetchedFlowers, _ := s.Db.GetFlowers(context.Background())

	oneFlowerLeft := s.Len(
		fetchedFlowers,
		1,
		"DeleteSite() should leave only the one flower not in the site",
	)
	if oneFlowerLeft {
		s.Equal(
			addedFlower2.ID,
			fetchedFlowers[0].ID,
			"wrong flower left after calling DeleteSite()",
		)
	}

	fetchedSites, _ := s.Db.GetRootSites(context.Background(), *site.Owner)

	oneRootSiteLeft := s.Len(
		fetchedSites,
		1,
		"DeleteSite() should only delete the selected site and its descendants",
	)
	if oneRootSiteLeft {
		s.Equal(
			addedSite2.ID,
			fetchedSites[0].ID,
			"wrong root site left after calling DeleteSite()",
		)
	}

	_, err = s.Db.GetSiteByID(context.Background(), addedSubSite.ID)

	s.Equal(
		mongo.ErrNoDocuments,
		err,
		"DeleteSite() should delete all subsites",
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
