package tests

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/Slowers-team/Slowers-App/databases/mongo"
	"github.com/Slowers-team/Slowers-App/testdata"
	"github.com/Slowers-team/Slowers-App/testutils"
)

type DbFlowerTestSuite struct {
	suite.Suite
	MongoDb mongo.Database
	Flowers []mongo.Flower
	Images  []mongo.Image
}

func (s *DbFlowerTestSuite) SetupSuite() {
	s.MongoDb = testutils.ConnectMongoDB()
	s.MongoDb.Clear()
	s.Flowers = testdata.GetFlowers()
	s.Images = testdata.GetImages()
}

func (s *DbFlowerTestSuite) TestAddFlower() {
	flowerToAdd := testdata.PrepareFlowerForAdding(s.Flowers[0])
	addedFlower, err := s.MongoDb.AddFlower(context.Background(), flowerToAdd)

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
	s.MongoDb.AddFlower(context.Background(), flowerToAdd)
	fetchedFlowers, err := s.MongoDb.GetFlowers(context.Background())

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
	addedFlower, _ := s.MongoDb.AddFlower(context.Background(), flowerToAdd)
	anyDeleted, err := s.MongoDb.DeleteFlower(context.Background(), addedFlower.ID)

	s.True(
		anyDeleted,
		"DeleteFlowers() should return true",
	)
	s.NoError(
		err,
		"DeleteFlowers() should not return an error",
	)

	fetchedFlowers, _ := s.MongoDb.GetFlowers(context.Background())

	s.Empty(
		fetchedFlowers,
		"deleted flower should not be returned by GetFlowers()",
	)
}

func (s *DbFlowerTestSuite) TestAddAndGetFlowersByUser() {
	flower := s.Flowers[0]
	flowerToAdd := testdata.PrepareFlowerForAdding(flower)
	addedFlower, _ := s.MongoDb.AddFlower(context.Background(), flowerToAdd)

	flower2 := testdata.GetFlowerForUser2()
	flowerToAdd2 := testdata.PrepareFlowerForAdding(flower2)
	s.MongoDb.AddFlower(context.Background(), flowerToAdd2)

	fetchedFlowers, err := s.MongoDb.GetUserFlowers(
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
	site.Flowers = []*mongo.ObjectID{}
	siteToAdd := testdata.PrepareSiteForAdding(site)
	addedSite, _ := s.MongoDb.AddSite(context.Background(), siteToAdd)

	site2 := testdata.GetRootSitesForUser2()[0]
	site2.Flowers = []*mongo.ObjectID{}
	siteToAdd2 := testdata.PrepareSiteForAdding(site2)
	addedSite2, _ := s.MongoDb.AddSite(context.Background(), siteToAdd2)

	flower := s.Flowers[0]
	flower.Site = &addedSite.ID
	flowerToAdd := testdata.PrepareFlowerForAdding(flower)
	addedFlower, _ := s.MongoDb.AddFlower(context.Background(), flowerToAdd)
	s.MongoDb.AddFlowerToSite(context.Background(), addedSite.ID, addedFlower.ID)

	flower2 := testdata.GetFlowerForUser2()
	flower2.Site = &addedSite2.ID
	flowerToAdd2 := testdata.PrepareFlowerForAdding(flower2)
	addedFlower2, _ := s.MongoDb.AddFlower(context.Background(), flowerToAdd2)
	s.MongoDb.AddFlowerToSite(context.Background(), addedSite2.ID, addedFlower2.ID)

	fetchedFlowers, err := s.MongoDb.GetAllFlowersRelatedToSite(
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
	addedFlower, _ := s.MongoDb.AddFlower(context.Background(), flowerToAdd)

	modifiedFields := mongo.Flower{
		Name:      "modified name",
		LatinName: "modified latin name",
		Quantity:  flower.Quantity + 1,
	}
	s.MongoDb.ModifyFlower(context.Background(), addedFlower.ID, modifiedFields)
	fetchedFlowers, err := s.MongoDb.GetFlowers(context.Background())

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
	addedFlowers := []mongo.Flower{}

	for _, flower := range flowers {
		flowerToAdd := testdata.PrepareFlowerForAdding(flower)
		addedFlower, _ := s.MongoDb.AddFlower(context.Background(), flowerToAdd)
		addedFlowers = append(addedFlowers, *addedFlower)
	}

	err := s.MongoDb.DeleteMultipleFlowers(
		context.Background(),
		[]mongo.ObjectID{addedFlowers[0].ID, addedFlowers[1].ID},
	)

	s.Require().NoError(
		err,
		"DeleteMultipleFlowers() should not return an error",
	)

	fetchedFlowers, _ := s.MongoDb.GetFlowers(context.Background())

	s.Equal(
		addedFlowers[2:],
		fetchedFlowers,
		"DeleteMultipleFlowers() should delete the correct flowers",
	)
}

func (s *DbFlowerTestSuite) TestSetVisibilityByTimeToFalse() {
	users := testdata.GetUsers()

	testFlower := mongo.Flower{
		Name:        s.Flowers[0].Name,
		LatinName:   s.Flowers[0].LatinName,
		AddedTime:   time.Now(),
		Grower:      s.Flowers[0].Grower,
		GrowerEmail: users[0].Email,
		Site:        s.Flowers[0].Site,
		SiteName:    testdata.GetRootSites()[0].Name,
		Quantity:    s.Flowers[0].Quantity,
		Visible:     true,
	}
	addedFlower, _ := s.MongoDb.AddFlower(context.Background(), testFlower)
	_ = addedFlower
	modified, err := s.MongoDb.UpdateVisibilityByTime(context.Background(), time.Now())

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

func (s *DbFlowerTestSuite) TestTimerResetsWhenTogglingToVisible() {
	users := testdata.GetUsers()
	firstTime := time.Now()

	testFlower := mongo.Flower{
		Name:        s.Flowers[0].Name,
		LatinName:   s.Flowers[0].LatinName,
		AddedTime:   firstTime,
		Grower:      s.Flowers[0].Grower,
		GrowerEmail: users[0].Email,
		Site:        s.Flowers[0].Site,
		SiteName:    testdata.GetRootSites()[0].Name,
		Quantity:    s.Flowers[0].Quantity,
		Visible:     false,
	}

	addedFlower, _ := s.MongoDb.AddFlower(context.Background(), testFlower)

	testImage := mongo.Image{
		ID:         s.Images[1].ID,
		FileFormat: s.Images[1].FileFormat,
		Note:       s.Images[1].Note,
		Entity:     &addedFlower.ID,
		Owner:      *s.Flowers[0].Grower,
	}

	fetchedFlowers, _ := s.MongoDb.GetUserFlowers(context.Background(), *s.Flowers[0].Grower)
	addedImage, _ := s.MongoDb.AddImage(context.Background(), testImage)

	err := s.MongoDb.SetFavoriteImage(
		context.Background(),
		*testFlower.Grower,
		fetchedFlowers[0].ID,
		addedImage.ID,
		"flowers",
	)
	s.Require().NoError(
		err,
		"SetFavoriteImage() should not return an error",
	)

	modified, err := s.MongoDb.ToggleFlowerVisibility(context.Background(), *testFlower.Grower, fetchedFlowers[0].ID)

	s.Require().NoError(
		err,
		"ToggleFlowerVisibility() should not return an error",
	)

	s.True(
		*modified,
		"modified value should be true",
	)

	fetchedFlowers, err = s.MongoDb.GetFlowers(context.Background())

	s.Require().NoError(
		err,
		"ToggleFlowerVisibility() should not return an error",
	)

	s.NotEqual(
		fetchedFlowers[0].AddedTime,
		firstTime,
		"AddedTime should not equal original timestamp",
	)

}

func (s *DbFlowerTestSuite) TearDownTest() {
	s.MongoDb.Clear()
}

func (s *DbFlowerTestSuite) TearDownSuite() {
	testutils.DisconnectMongoDB(s.MongoDb)
}

func TestDbFlowerTestSuite(t *testing.T) {
	suite.Run(t, new(DbFlowerTestSuite))
}
