package tests

import (
	"context"
	"testing"

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
	flower := s.Flowers[0]
	createdFlower, err := s.Db.AddFlower(context.Background(), flower)

	s.NoError(
		err,
		"AddFlower() should not return an error",
	)
	s.Equal(
		flower.Name,
		createdFlower.Name,
		"wrong name for the flower returned from AddFlower()",
	)
	s.Equal(
		flower.LatinName,
		createdFlower.LatinName,
		"wrong latin name for the flower returned from AddFlower()",
	)
	s.Equal(
		flower.AddedTime,
		createdFlower.AddedTime,
		"wrong AddedTime for the flower returned from AddFlower()",
	)
	s.Equal(
		flower.Quantity,
		createdFlower.Quantity,
		"wrong Quantity for the flower returned from AddFlower()",
	)
	s.NotZero(
		createdFlower.ID,
		"ID for the created flower should be non-zero",
	)
}

func (s *DbFlowerTestSuite) TestAddAndGetFlower() {
	flower := database.Flower{
		Name:      s.Flowers[0].Name,
		LatinName: s.Flowers[0].LatinName,
		Grower:    s.Flowers[0].Grower,
		Site:      s.Flowers[0].Site,
		Quantity:  s.Flowers[0].Quantity,
		Visible:   s.Flowers[0].Visible,
	}
	s.Db.AddFlower(context.Background(), flower)
	fetchedFlowers, err := s.Db.GetFlowers(context.Background())

	s.NoError(
		err,
		"GetFlowers() should not return an error",
	)
	s.Len(
		fetchedFlowers,
		1,
		"GetFlowers() should return a slice of length 1",
	)
	s.NotZero(
		fetchedFlowers[0].ID,
		"fetched flower should have non-zero ID",
	)
	s.Equal(
		flower.Name,
		fetchedFlowers[0].Name,
		"wrong Name for the flower returned from GetFlowers()",
	)
	s.Equal(
		flower.LatinName,
		fetchedFlowers[0].LatinName,
		"wrong LatinName for the flower returned from GetFlowers()",
	)
	s.Equal(
		flower.AddedTime,
		fetchedFlowers[0].AddedTime,
		"wrong AddedTime for the flower returned from GetFlowers()",
	)
	s.Equal(
		flower.Quantity,
		fetchedFlowers[0].Quantity,
		"wrong Quantity for the flower returned from GetFlowers()",
	)
	s.Equal(
		flower.Visible,
		fetchedFlowers[0].Visible,
		"wrong Visible for the flower returned from GetFlowers()",
	)
}

func (s *DbFlowerTestSuite) TestAddAndDeleteFlower() {
	flower := database.Flower{
		Name:      s.Flowers[0].Name,
		LatinName: s.Flowers[0].LatinName,
		Grower:    s.Flowers[0].Grower,
		Site:      s.Flowers[0].Site,
		Quantity:  s.Flowers[0].Quantity,
	}
	createdFlower, _ := s.Db.AddFlower(context.Background(), flower)
	anyDeleted, err := s.Db.DeleteFlower(
		context.Background(), createdFlower.ID,
	)

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
	users := testdata.GetUsers()

	flower := database.Flower{
		Name:        s.Flowers[0].Name,
		LatinName:   s.Flowers[0].LatinName,
		AddedTime:   s.Flowers[0].AddedTime,
		Grower:      s.Flowers[0].Grower,
		GrowerEmail: users[0].Email,
		Site:        s.Flowers[0].Site,
		SiteName:    testdata.GetRootSites()[0].Name,
		Quantity:    s.Flowers[0].Quantity,
		Visible:     s.Flowers[0].Visible,
	}
	addedFlower, _ := s.Db.AddFlower(context.Background(), flower)

	fullFlower2 := testdata.GetTestFlowerForUser2()
	flower2 := database.Flower{
		Name:        fullFlower2.Name,
		LatinName:   fullFlower2.LatinName,
		AddedTime:   fullFlower2.AddedTime,
		Grower:      fullFlower2.Grower,
		GrowerEmail: users[1].Email,
		Site:        fullFlower2.Site,
		SiteName:    testdata.GetRootSitesForUser2()[0].Name,
		Quantity:    fullFlower2.Quantity,
		Visible:     fullFlower2.Visible,
	}
	s.Db.AddFlower(context.Background(), flower2)

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
	users := testdata.GetUsers()

	site1 := testdata.GetRootSites()[0]
	site1.ID = database.NilObjectID
	site1.Flowers = []*database.ObjectID{}
	addedSite1, _ := s.Db.AddSite(context.Background(), site1)

	site2 := testdata.GetRootSitesForUser2()[0]
	site2.ID = database.NilObjectID
	site2.Flowers = []*database.ObjectID{}
	addedSite2, _ := s.Db.AddSite(context.Background(), site2)

	flower := database.Flower{
		Name:        s.Flowers[0].Name,
		LatinName:   s.Flowers[0].LatinName,
		AddedTime:   s.Flowers[0].AddedTime,
		Grower:      s.Flowers[0].Grower,
		GrowerEmail: users[0].Email,
		Site:        &addedSite1.ID,
		SiteName:    site1.Name,
		Quantity:    s.Flowers[0].Quantity,
	}
	addedFlower, _ := s.Db.AddFlower(context.Background(), flower)
	s.Db.AddFlowerToSite(context.Background(), addedSite1.ID, addedFlower.ID)

	fullFlower2 := testdata.GetTestFlowerForUser2()
	flower2 := database.Flower{
		Name:        fullFlower2.Name,
		LatinName:   fullFlower2.LatinName,
		AddedTime:   fullFlower2.AddedTime,
		Grower:      fullFlower2.Grower,
		GrowerEmail: users[1].Email,
		Site:        &addedSite2.ID,
		SiteName:    site2.Name,
		Quantity:    fullFlower2.Quantity,
	}
	addedFlower2, _ := s.Db.AddFlower(context.Background(), flower2)
	s.Db.AddFlowerToSite(context.Background(), addedSite2.ID, addedFlower2.ID)

	fetchedFlowers, err := s.Db.GetAllFlowersRelatedToSite(
		context.Background(), addedSite1.ID, *flower.Grower,
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
	s.Db.AddFlower(context.Background(), flower)

	modifiedFlower := database.Flower{
		Name:      "modified name",
		LatinName: "modified latin name",
		Quantity:  flower.Quantity + 1,
	}
	s.Db.ModifyFlower(context.Background(), flower.ID, modifiedFlower)
	fetchedFlowers, err := s.Db.GetFlowers(context.Background())

	s.Require().NoError(
		err,
		"ModifyFlower() should not return an error",
	)
	s.Equal(
		modifiedFlower.Name,
		fetchedFlowers[0].Name,
		"wrong Name for the flower returned from GetFlowers()",
	)
	s.Equal(
		modifiedFlower.LatinName,
		fetchedFlowers[0].LatinName,
		"wrong LatinName for the flower returned from GetFlowers()",
	)
	s.Equal(
		modifiedFlower.Quantity,
		fetchedFlowers[0].Quantity,
		"wrong Quantity for the flower returned from GetFlowers()",
	)
	s.Equal(
		flower.ID,
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

	for _, flower := range flowers {
		s.Db.AddFlower(context.Background(), flower)
	}

	err := s.Db.DeleteMultipleFlowers(context.Background(), []database.ObjectID{flowers[0].ID, flowers[1].ID})
	fetchedFlowers, _ := s.Db.GetFlowers(context.Background())

	s.Require().NoError(
		err,
		"DeleteMultipleFlowers() should not return an error",
	)
	s.Equal(
		flowers[2:],
		fetchedFlowers,
		"DeleteMultipleFlowers() should delete the correct flowers",
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
