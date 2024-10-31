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
	Db          database.Database
	TestFlowers []database.Flower
}

func (s *DbFlowerTestSuite) SetupSuite() {
	s.Db = testutils.ConnectDB()
	s.Db.Clear()
	s.TestFlowers = testdata.GetTestFlowers()
}

func (s *DbFlowerTestSuite) TestAddFlower() {
	flower := s.TestFlowers[0]
	createdFlower, err := s.Db.AddFlower(context.Background(), flower)

	s.NoError(
		err,
		"AddFlower() should not return an error",
	)
	s.Equal(
		createdFlower.Name,
		flower.Name,
		"wrong name for the flower returned from AddFlower()",
	)
	s.Equal(
		createdFlower.LatinName,
		flower.LatinName,
		"wrong latin name for the flower returned from AddFlower()",
	)
	s.Equal(
		createdFlower.AddedTime,
		flower.AddedTime,
		"wrong AddedTime for the flower returned from AddFlower()",
	)
	s.NotZero(
		createdFlower.ID,
		"ID for the created flower should be non-zero",
	)
}

func (s *DbFlowerTestSuite) TestAddAndGetFlower() {
	flower := database.Flower{
		Name:      s.TestFlowers[0].Name,
		LatinName: s.TestFlowers[0].LatinName,
		Grower:    s.TestFlowers[0].Grower,
		Site:      s.TestFlowers[0].Site,
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
		fetchedFlowers[0].Name,
		flower.Name,
		"wrong Name for the flower returned from GetFlowers()",
	)
	s.Equal(
		fetchedFlowers[0].LatinName,
		flower.LatinName,
		"wrong LatinName for the flower returned from GetFlowers()",
	)
	s.Equal(
		fetchedFlowers[0].AddedTime,
		flower.AddedTime,
		"wrong AddedTime for the flower returned from GetFlowers()",
	)
}

func (s *DbFlowerTestSuite) TestAddAndDeleteFlower() {
	testFlower := database.Flower{
		Name:      s.TestFlowers[0].Name,
		LatinName: s.TestFlowers[0].LatinName,
		Grower:    s.TestFlowers[0].Grower,
		Site:      s.TestFlowers[0].Site,
	}
	flower, _ := s.Db.AddFlower(context.Background(), testFlower)
	anyDeleted, err := s.Db.DeleteFlower(context.Background(), flower.ID)

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

	testFlower := database.Flower{
		Name:        s.TestFlowers[0].Name,
		LatinName:   s.TestFlowers[0].LatinName,
		Grower:      s.TestFlowers[0].Grower,
		GrowerEmail: users[0].Email,
		Site:        s.TestFlowers[0].Site,
		SiteName:    testdata.GetRootSites()[0].Name,
	}
	addedFlower, _ := s.Db.AddFlower(context.Background(), testFlower)

	fullFlower2 := testdata.GetTestFlowerForUser2()
	testFlower2 := database.Flower{
		Name:        fullFlower2.Name,
		LatinName:   fullFlower2.LatinName,
		Grower:      fullFlower2.Grower,
		GrowerEmail: users[1].Email,
		Site:        fullFlower2.Site,
		SiteName:    testdata.GetRootSitesForUser2()[0].Name,
	}
	s.Db.AddFlower(context.Background(), testFlower2)

	fetchedFlowers, err := s.Db.GetUserFlowers(context.Background(), *testFlower.Grower)

	s.NoError(
		err,
		"GetUserFlowers() should not return an error",
	)
	s.Len(
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
		testFlower.Name,
		fetchedFlowers[0].Name,
		"wrong Name for the flower returned from GetUserFlowers()",
	)
	s.Equal(
		testFlower.LatinName,
		fetchedFlowers[0].LatinName,
		"wrong LatinName for the flower returned from GetUserFlowers()",
	)
	s.Equal(
		testFlower.AddedTime,
		fetchedFlowers[0].AddedTime,
		"wrong AddedTime for the flower returned from GetUserFlowers()",
	)
	s.Equal(
		*testFlower.Grower,
		*fetchedFlowers[0].Grower,
		"wrong Grower for the flower returned from GetUserFlowers()",
	)
	s.Equal(
		testFlower.GrowerEmail,
		fetchedFlowers[0].GrowerEmail,
		"wrong GrowerEmail for the flower returned from GetUserFlowers()",
	)
	s.Equal(
		*testFlower.Site,
		*fetchedFlowers[0].Site,
		"wrong Site for the flower returned from GetUserFlowers()",
	)
	s.Equal(
		testFlower.SiteName,
		fetchedFlowers[0].SiteName,
		"wrong SiteName for the flower returned from GetUserFlowers()",
	)
}

func (s *DbFlowerTestSuite) TestAddAndGetFlowersRelatedToSite() {
	users := testdata.GetUsers()

	site1 := testdata.GetRootSites()[0]
	site1.ID = database.NilObjectID
	addedSite1, _ := s.Db.AddSite(context.Background(), site1)

	site2 := testdata.GetRootSitesForUser2()[0]
	site2.ID = database.NilObjectID
	addedSite2, _ := s.Db.AddSite(context.Background(), site2)

	testFlower := database.Flower{
		Name:        s.TestFlowers[0].Name,
		LatinName:   s.TestFlowers[0].LatinName,
		Grower:      s.TestFlowers[0].Grower,
		GrowerEmail: users[0].Email,
		Site:        &addedSite1.ID,
		SiteName:    site1.Name,
	}
	addedFlower, _ := s.Db.AddFlower(context.Background(), testFlower)

	fullFlower2 := testdata.GetTestFlowerForUser2()
	testFlower2 := database.Flower{
		Name:        fullFlower2.Name,
		LatinName:   fullFlower2.LatinName,
		Grower:      fullFlower2.Grower,
		GrowerEmail: users[1].Email,
		Site:        &addedSite2.ID,
		SiteName:    site2.Name,
	}
	s.Db.AddFlower(context.Background(), testFlower2)

	fetchedFlowers, err := s.Db.GetAllFlowersRelatedToSite(
		context.Background(), addedSite1.ID, *testFlower.Grower,
	)

	s.NoError(
		err,
		"GetAllFlowersRelatedToSite() should not return an error",
	)
	s.Len(
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
		testFlower.Name,
		fetchedFlowers[0].Name,
		"wrong Name for the flower returned from GetAllFlowersRelatedToSite()",
	)
	s.Equal(
		testFlower.LatinName,
		fetchedFlowers[0].LatinName,
		"wrong LatinName for the flower returned from GetAllFlowersRelatedToSite()",
	)
	s.Equal(
		testFlower.AddedTime,
		fetchedFlowers[0].AddedTime,
		"wrong AddedTime for the flower returned from GetAllFlowersRelatedToSite()",
	)
	s.Equal(
		*testFlower.Grower,
		*fetchedFlowers[0].Grower,
		"wrong Grower for the flower returned from GetAllFlowersRelatedToSite()",
	)
	s.Equal(
		testFlower.GrowerEmail,
		fetchedFlowers[0].GrowerEmail,
		"wrong GrowerEmail for the flower returned from GetAllFlowersRelatedToSite()",
	)
	s.Equal(
		*testFlower.Site,
		*fetchedFlowers[0].Site,
		"wrong Site for the flower returned from GetAllFlowersRelatedToSite()",
	)
	s.Equal(
		testFlower.SiteName,
		fetchedFlowers[0].SiteName,
		"wrong SiteName for the flower returned from GetAllFlowersRelatedToSite()",
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
