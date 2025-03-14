package tests

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/Slowers-team/Slowers-App/database"
	"github.com/Slowers-team/Slowers-App/testdata"
	"github.com/Slowers-team/Slowers-App/testutils"
)

type DbFlowerTestSuite struct {
	suite.Suite
	Db      database.Database
	Flowers []database.Flower
}

func (s *DbFlowerTestSuite) SetupSuite() {
	s.Db = testutils.ConnectDB()
	s.Db.Clear()
	s.Flowers = testdata.GetFlowers()
}

func (s *DbFlowerTestSuite) TestAddFlower() {
	flowerToAdd := testdata.PrepareFlowerForAdding(s.Flowers[0])
	addedFlower, err := s.Db.AddFlower(context.Background(), flowerToAdd)

	s.Require().NoError(
		err,
		"AddFlower() should not return an error",
	)
	s.NotZero(
		addedFlower.ID,
		"ID for the created flower should be non-zero",
	)
	s.Equal(
		flowerToAdd.Name,
		addedFlower.Name,
		"wrong name for the flower returned from AddFlower()",
	)
	s.Equal(
		flowerToAdd.LatinName,
		addedFlower.LatinName,
		"wrong latin name for the flower returned from AddFlower()",
	)
	s.Equal(
		flowerToAdd.AddedTime,
		addedFlower.AddedTime,
		"wrong AddedTime for the flower returned from AddFlower()",
	)
	s.Equal(
		flowerToAdd.Quantity,
		addedFlower.Quantity,
		"wrong Quantity for the flower returned from AddFlower()",
	)
}

func (s *DbFlowerTestSuite) TestAddAndGetFlower() {
	flowerToAdd := testdata.PrepareFlowerForAdding(s.Flowers[0])
	s.Db.AddFlower(context.Background(), flowerToAdd)
	fetchedFlowers, err := s.Db.GetFlowers(context.Background())

	s.Require().NoError(
		err,
		"GetFlowers() should not return an error",
	)
	s.Require().Len(
		fetchedFlowers,
		1,
		"GetFlowers() should return a slice of length 1",
	)
	s.NotZero(
		fetchedFlowers[0].ID,
		"fetched flower should have non-zero ID",
	)
	s.Equal(
		flowerToAdd.Name,
		fetchedFlowers[0].Name,
		"wrong Name for the flower returned from GetFlowers()",
	)
	s.Equal(
		flowerToAdd.LatinName,
		fetchedFlowers[0].LatinName,
		"wrong LatinName for the flower returned from GetFlowers()",
	)
	s.Equal(
		flowerToAdd.AddedTime,
		fetchedFlowers[0].AddedTime,
		"wrong AddedTime for the flower returned from GetFlowers()",
	)
	s.Equal(
		flowerToAdd.Quantity,
		fetchedFlowers[0].Quantity,
		"wrong Quantity for the flower returned from GetFlowers()",
	)
	s.Equal(
		flowerToAdd.Visible,
		fetchedFlowers[0].Visible,
		"wrong Visible for the flower returned from GetFlowers()",
	)
}

func (s *DbFlowerTestSuite) TestAddAndDeleteFlower() {
	flowerToAdd := testdata.PrepareFlowerForAdding(s.Flowers[0])
	addedFlower, _ := s.Db.AddFlower(context.Background(), flowerToAdd)
	anyDeleted, err := s.Db.DeleteFlower(context.Background(), addedFlower.ID)

	s.True(
		anyDeleted,
		"DeleteFlowers() should return true",
	)
	s.NoError(
		err,
		"DeleteFlowers() should not return an error",
	)

	fetchedFlowers, _ := s.Db.GetFlowers(context.Background())

	s.Empty(
		fetchedFlowers,
		"deleted flower should not be returned by GetFlowers()",
	)
}

func (s *DbFlowerTestSuite) TestAddAndGetFlowersByUser() {
	flower := s.Flowers[0]
	flowerToAdd := testdata.PrepareFlowerForAdding(flower)
	addedFlower, _ := s.Db.AddFlower(context.Background(), flowerToAdd)

	flower2 := testdata.GetFlowerForUser2()
	flowerToAdd2 := testdata.PrepareFlowerForAdding(flower2)
	s.Db.AddFlower(context.Background(), flowerToAdd2)

	fetchedFlowers, err := s.Db.GetUserFlowers(
		context.Background(), *flower.Grower,
	)

	s.Require().NoError(
		err,
		"GetUserFlowers() should not return an error",
	)

	s.Require().Len(
		fetchedFlowers,
		1,
		"GetUserFlowers() should return a slice of length 1",
	)

	s.Equal(
		addedFlower.ID,
		fetchedFlowers[0].ID,
		"wrong ID for the flower returned from GetUserFlowers()",
	)
	s.Equal(
		flower.Name,
		fetchedFlowers[0].Name,
		"wrong Name for the flower returned from GetUserFlowers()",
	)
	s.Equal(
		flower.LatinName,
		fetchedFlowers[0].LatinName,
		"wrong LatinName for the flower returned from GetUserFlowers()",
	)
	s.Equal(
		flower.AddedTime,
		fetchedFlowers[0].AddedTime,
		"wrong AddedTime for the flower returned from GetUserFlowers()",
	)
	s.Equal(
		*flower.Grower,
		*fetchedFlowers[0].Grower,
		"wrong Grower for the flower returned from GetUserFlowers()",
	)
	s.Equal(
		flower.GrowerEmail,
		fetchedFlowers[0].GrowerEmail,
		"wrong GrowerEmail for the flower returned from GetUserFlowers()",
	)
	s.Equal(
		*flower.Site,
		*fetchedFlowers[0].Site,
		"wrong Site for the flower returned from GetUserFlowers()",
	)
	s.Equal(
		flower.SiteName,
		fetchedFlowers[0].SiteName,
		"wrong SiteName for the flower returned from GetUserFlowers()",
	)
	s.Equal(
		flower.Quantity,
		fetchedFlowers[0].Quantity,
		"wrong Quantity for the flower returned from GetUserFlowers()",
	)
	s.Equal(
		flower.Visible,
		fetchedFlowers[0].Visible,
		"wrong Visible for the flower returned from GetUserFlowers()",
	)
}

func (s *DbFlowerTestSuite) TestAddAndGetFlowersRelatedToSite() {
	site := testdata.GetRootSites()[0]
	site.Flowers = []*database.ObjectID{}
	siteToAdd := testdata.PrepareSiteForAdding(site)
	addedSite, _ := s.Db.AddSite(context.Background(), siteToAdd)

	site2 := testdata.GetRootSitesForUser2()[0]
	site2.Flowers = []*database.ObjectID{}
	siteToAdd2 := testdata.PrepareSiteForAdding(site2)
	addedSite2, _ := s.Db.AddSite(context.Background(), siteToAdd2)

	flower := s.Flowers[0]
	flower.Site = &addedSite.ID
	flowerToAdd := testdata.PrepareFlowerForAdding(flower)
	addedFlower, _ := s.Db.AddFlower(context.Background(), flowerToAdd)
	s.Db.AddFlowerToSite(context.Background(), addedSite.ID, addedFlower.ID)

	flower2 := testdata.GetFlowerForUser2()
	flower2.Site = &addedSite2.ID
	flowerToAdd2 := testdata.PrepareFlowerForAdding(flower2)
	addedFlower2, _ := s.Db.AddFlower(context.Background(), flowerToAdd2)
	s.Db.AddFlowerToSite(context.Background(), addedSite2.ID, addedFlower2.ID)

	fetchedFlowers, err := s.Db.GetAllFlowersRelatedToSite(
		context.Background(), addedSite.ID, *site.Owner,
	)

	s.Require().NoError(
		err,
		"GetAllFlowersRelatedToSite() should not return an error",
	)

	s.Require().Len(
		fetchedFlowers,
		1,
		"GetAllFlowersRelatedToSite() should return a slice of length 1",
	)

	s.Equal(
		addedFlower.ID,
		fetchedFlowers[0].ID,
		"wrong ID for the flower returned from GetAllFlowersRelatedToSite()",
	)
	s.Equal(
		flower.Name,
		fetchedFlowers[0].Name,
		"wrong Name for the flower returned from GetAllFlowersRelatedToSite()",
	)
	s.Equal(
		flower.LatinName,
		fetchedFlowers[0].LatinName,
		"wrong LatinName for the flower returned from GetAllFlowersRelatedToSite()",
	)
	s.Equal(
		flower.AddedTime,
		fetchedFlowers[0].AddedTime,
		"wrong AddedTime for the flower returned from GetAllFlowersRelatedToSite()",
	)
	s.Equal(
		*flower.Grower,
		*fetchedFlowers[0].Grower,
		"wrong Grower for the flower returned from GetAllFlowersRelatedToSite()",
	)
	s.Equal(
		flower.GrowerEmail,
		fetchedFlowers[0].GrowerEmail,
		"wrong GrowerEmail for the flower returned from GetAllFlowersRelatedToSite()",
	)
	s.Equal(
		*flower.Site,
		*fetchedFlowers[0].Site,
		"wrong Site for the flower returned from GetAllFlowersRelatedToSite()",
	)
	s.Equal(
		flower.SiteName,
		fetchedFlowers[0].SiteName,
		"wrong SiteName for the flower returned from GetAllFlowersRelatedToSite()",
	)
	s.Equal(
		flower.Quantity,
		fetchedFlowers[0].Quantity,
		"wrong Quantity for the flower returned from GetAllFlowersRelatedToSite()",
	)
}

func (s *DbFlowerTestSuite) TestModifyAndGetFlower() {
	flower := s.Flowers[0]
	flowerToAdd := testdata.PrepareFlowerForAdding(flower)
	addedFlower, _ := s.Db.AddFlower(context.Background(), flowerToAdd)

	modifiedFields := database.Flower{
		Name:      "modified name",
		LatinName: "modified latin name",
		Quantity:  flower.Quantity + 1,
	}
	s.Db.ModifyFlower(context.Background(), addedFlower.ID, modifiedFields)
	fetchedFlowers, err := s.Db.GetFlowers(context.Background())

	s.Require().NoError(
		err,
		"ModifyFlower() should not return an error",
	)
	s.Equal(
		modifiedFields.Name,
		fetchedFlowers[0].Name,
		"wrong Name for the flower returned from GetFlowers()",
	)
	s.Equal(
		modifiedFields.LatinName,
		fetchedFlowers[0].LatinName,
		"wrong LatinName for the flower returned from GetFlowers()",
	)
	s.Equal(
		modifiedFields.Quantity,
		fetchedFlowers[0].Quantity,
		"wrong Quantity for the flower returned from GetFlowers()",
	)
	s.Equal(
		addedFlower.ID,
		fetchedFlowers[0].ID,
		"wrong ID for the flower returned from GetFlowers()",
	)
	s.Equal(
		flower.AddedTime,
		fetchedFlowers[0].AddedTime,
		"wrong AddedTime for the flower returned from GetFlowers()",
	)
	s.Equal(
		flower.Grower,
		fetchedFlowers[0].Grower,
		"wrong Grower for the flower returned from GetFlowers()",
	)
	s.Equal(
		flower.Site,
		fetchedFlowers[0].Site,
		"wrong Site for the flower returned from GetFlowers()",
	)
	s.Equal(
		flower.GrowerEmail,
		fetchedFlowers[0].GrowerEmail,
		"wrong GrowerEmail for the flower returned from GetFlowers()",
	)
	s.Equal(
		flower.Visible,
		fetchedFlowers[0].Visible,
		"wrong Visible for the flower returned from GetFlowers()",
	)
}

func (s *DbFlowerTestSuite) TestDeleteAndGetMultipleFlowers() {
	flowers := testdata.GetFlowers()
	addedFlowers := []database.Flower{}

	for _, flower := range flowers {
		flowerToAdd := testdata.PrepareFlowerForAdding(flower)
		addedFlower, _ := s.Db.AddFlower(context.Background(), flowerToAdd)
		addedFlowers = append(addedFlowers, *addedFlower)
	}

	err := s.Db.DeleteMultipleFlowers(
		context.Background(),
		[]database.ObjectID{addedFlowers[0].ID, addedFlowers[1].ID},
	)

	s.Require().NoError(
		err,
		"DeleteMultipleFlowers() should not return an error",
	)

	fetchedFlowers, _ := s.Db.GetFlowers(context.Background())

	s.Equal(
		addedFlowers[2:],
		fetchedFlowers,
		"DeleteMultipleFlowers() should delete the correct flowers",
	)
}

func (s *DbFlowerTestSuite) TestSetVisibilityByTimeToFalse() {
	users := testdata.GetUsers()

	testFlower := database.Flower{
		Name:        s.TestFlowers[0].Name,
		LatinName:   s.TestFlowers[0].LatinName,
		AddedTime:   time.Now(),
		Grower:      s.TestFlowers[0].Grower,
		GrowerEmail: users[0].Email,
		Site:        s.TestFlowers[0].Site,
		SiteName:    testdata.GetRootSites()[0].Name,
		Quantity:    s.TestFlowers[0].Quantity,
		Visible:     true,
	}
	addedFlower, _ := s.Db.AddFlower(context.Background(), testFlower)
	_ = addedFlower
	modified, err := s.Db.UpdateVisibilityByTime(context.Background(), time.Now())

	s.Require().NoError(
		err,
		"UpdateVisibilityByTime() should not return an error",
	)
	s.Equal(
		modified,
		int64(1),
		"UpdateVisibilityByTime() should set one flower invisible",
	)
}

func (s *DbFlowerTestSuite) TearDownTest() {
	s.Db.Clear()
}

func (s *DbFlowerTestSuite) TearDownSuite() {
	testutils.DisconnectDB(s.Db)
}

func TestDbFlowerTestSuite(t *testing.T) {
	suite.Run(t, new(DbFlowerTestSuite))
}
