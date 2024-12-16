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
	"github.com/Slowers-team/Slowers-App/utils"
)

type DbSiteTestSuite struct {
	suite.Suite
	Db       database.Database
	TestSite database.Site
	TestUser database.User
}

func (s *DbSiteTestSuite) SetupSuite() {
	s.Db = testutils.ConnectDB()
	s.Db.Clear()
	s.TestSite = testdata.GetRootSites()[0]
	s.TestUser = testdata.GetUsers()[0]
}

func (s *DbSiteTestSuite) TestAddSite() {
	site := database.Site{
		Name:      s.TestSite.Name,
		AddedTime: s.TestSite.AddedTime,
		Note:      s.TestSite.Note,
		Parent:    s.TestSite.Parent,
		Flowers:   s.TestSite.Flowers,
		Owner:     s.TestSite.Owner,
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
		site.Name,
		createdSite.Name,
		"wrong name for the site returned from AddSite()",
	)
	s.Equal(
		site.AddedTime,
		createdSite.AddedTime,
		"wrong AddedTime for the site returned from AddSite()",
	)
	s.Equal(
		site.Note,
		createdSite.Note,
		"wrong note for the site returned from AddSite()",
	)
	s.Equal(
		site.Parent,
		createdSite.Parent,
		"wrong parent for the site returned from AddSite()",
	)
	s.True(
		utils.AreIDPtrSlicesEql(createdSite.Flowers, site.Flowers),
		"wrong flowers for the site returned from AddSite()",
	)
	s.Equal(
		site.Owner,
		createdSite.Owner,
		"wrong owner for the site returned from AddSite()",
	)
}

func (s *DbSiteTestSuite) TestAddAndGetRootSites() {
	site := database.Site{
		Name:      s.TestSite.Name,
		AddedTime: s.TestSite.AddedTime,
		Note:      s.TestSite.Note,
		Parent:    s.TestSite.Parent,
		Flowers:   s.TestSite.Flowers,
		Owner:     s.TestSite.Owner,
	}
	createdSite, _ := s.Db.AddSite(context.Background(), site)
	rootSites, err := s.Db.GetRootSites(context.Background(), s.TestUser.ID)

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
		createdSite.ID,
		rootSites[0].ID,
		"ID for the fetched site should be non-zero",
	)
	s.Equal(
		site.Name,
		rootSites[0].Name,
		"wrong name for the site returned from GetRootSites()",
	)
	s.Equal(
		site.AddedTime,
		rootSites[0].AddedTime,
		"wrong AddedTime for the site returned from GetRootSites()",
	)
	s.Equal(
		site.Note,
		rootSites[0].Note,
		"wrong note for the site returned from GetRootSites()",
	)
	s.Equal(
		site.Parent,
		rootSites[0].Parent,
		"wrong parent for the site returned from GetRootSites()",
	)
	s.True(
		utils.AreIDPtrSlicesEql(rootSites[0].Flowers, site.Flowers),
		"wrong flowers for the site returned from GetRootSites()",
	)
	s.Equal(
		site.Owner,
		rootSites[0].Owner,
		"wrong owner for the site returned from GetRootSites()",
	)
}

func (s *DbSiteTestSuite) TestAddAndGetSite() {
	siteData := testdata.GetSite()

	site := siteData["site"].(database.Site)
	site.ID = database.NilObjectID
	createdSite, _ := s.Db.AddSite(context.Background(), site)

	subSiteBson := siteData["subsites"].([]bson.M)[0]
	subSite := database.Site{
		Name:      subSiteBson["name"].(string),
		AddedTime: time.Date(2024, 9, 19, 12, 11, 4, 0, time.UTC),
		Note:      subSiteBson["note"].(string),
		Parent:    &createdSite.ID,
		Flowers:   []*database.ObjectID{},
		Owner:     site.Owner,
	}
	createdSubSite, _ := s.Db.AddSite(context.Background(), subSite)

	site.ID = createdSite.ID
	siteData["site"] = site
	subSiteBson["_id"] = createdSubSite.ID.Hex()
	siteData["subsites"] = []bson.M{subSiteBson}

	fetchedSiteData, err := s.Db.GetSite(context.Background(), createdSite.ID, *site.Owner)

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
		createdSite.ID,
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
		createdSubSite.ID,
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
	site := testdata.GetRootSites()[0]
	site.ID = database.NilObjectID
	createdSite, _ := s.Db.AddSite(context.Background(), site)

	fetchedSite, err := s.Db.GetSiteByID(context.Background(), createdSite.ID)

	s.Require().NoError(
		err,
		"GetSiteByID() should not return an error",
	)

	s.Equal(
		createdSite.ID,
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
	site := testdata.GetRootSites()[0]
	site.ID = database.NilObjectID
	site.Flowers = []*database.ObjectID{}
	createdSite, _ := s.Db.AddSite(context.Background(), site)

	fullFlower := testdata.GetFlowers()[0]
	flowerToAdd := database.Flower{
		Name:        fullFlower.Name,
		LatinName:   fullFlower.LatinName,
		AddedTime:   fullFlower.AddedTime,
		Grower:      fullFlower.Grower,
		GrowerEmail: testdata.GetUsers()[0].Email,
		Site:        &createdSite.ID,
		SiteName:    site.Name,
	}
	addedFlower, _ := s.Db.AddFlower(context.Background(), flowerToAdd)
	err := s.Db.AddFlowerToSite(context.Background(), createdSite.ID, addedFlower.ID)

	s.Require().NoError(
		err,
		"AddFlowerToSite() should not return an error",
	)

	fetchedSite, _ := s.Db.GetSiteByID(context.Background(), createdSite.ID)

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
	site.ID = database.NilObjectID
	site.Flowers = []*database.ObjectID{}
	createdSite, _ := s.Db.AddSite(context.Background(), site)

	subSiteBson := siteData["subsites"].([]bson.M)[0]
	subSite := database.Site{
		Name:      subSiteBson["name"].(string),
		AddedTime: time.Date(2024, 9, 19, 12, 11, 4, 0, time.UTC),
		Note:      subSiteBson["note"].(string),
		Parent:    &createdSite.ID,
		Flowers:   []*database.ObjectID{},
		Owner:     site.Owner,
	}
	createdSubSite, _ := s.Db.AddSite(context.Background(), subSite)

	site.ID = createdSite.ID
	siteData["site"] = site
	subSiteBson["_id"] = createdSubSite.ID.Hex()
	siteData["subsites"] = []bson.M{subSiteBson}

	fullFlower := testdata.GetFlowers()[0]
	flowerToAdd := database.Flower{
		Name:        fullFlower.Name,
		LatinName:   fullFlower.LatinName,
		AddedTime:   fullFlower.AddedTime,
		Grower:      fullFlower.Grower,
		GrowerEmail: testdata.GetUsers()[0].Email,
		Site:        &createdSite.ID,
		SiteName:    site.Name,
		Visible:     fullFlower.Visible,
	}
	addedFlower, _ := s.Db.AddFlower(context.Background(), flowerToAdd)
	s.Db.AddFlowerToSite(context.Background(), createdSite.ID, addedFlower.ID)

	site2 := testdata.GetRootSitesForUser2()[0]
	site2.ID = database.NilObjectID
	site2.Flowers = []*database.ObjectID{}
	site2.Owner = site.Owner
	createdSite2, _ := s.Db.AddSite(context.Background(), site2)

	fullFlower2 := testdata.GetTestFlowerForUser2()
	flowerToAdd2 := database.Flower{
		Name:        fullFlower2.Name,
		LatinName:   fullFlower2.LatinName,
		AddedTime:   fullFlower2.AddedTime,
		Grower:      fullFlower.Grower,
		GrowerEmail: testdata.GetUsers()[0].Email,
		Site:        &createdSite2.ID,
		SiteName:    site2.Name,
		Visible:     fullFlower2.Visible,
	}
	addedFlower2, _ := s.Db.AddFlower(context.Background(), flowerToAdd2)
	s.Db.AddFlowerToSite(context.Background(), createdSite2.ID, addedFlower2.ID)

	deleteResult, err := s.Db.DeleteSite(context.Background(), createdSite.ID, *site.Owner)

	s.NoError(
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

	rootSites, _ := s.Db.GetRootSites(context.Background(), *site.Owner)

	oneRootSiteLeft := s.Len(
		rootSites,
		1,
		"DeleteSite() should only delete the selected site and its descendants",
	)
	if oneRootSiteLeft {
		s.Equal(
			createdSite2.ID,
			rootSites[0].ID,
			"wrong root site left after calling DeleteSite()",
		)
	}

	_, err = s.Db.GetSiteByID(context.Background(), createdSubSite.ID)

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
