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
	anyDeleted, err := s.Db.DeleteFlower(context.Background(), flower.ID.Hex())

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

func (s *DbFlowerTestSuite) TearDownTest() {
	s.Db.Clear()
}

func (s *DbFlowerTestSuite) TearDownSuite() {
	testutils.DisconnectDB(s.Db)
}

func TestDbFlowerTestSuite(t *testing.T) {
	suite.Run(t, new(DbFlowerTestSuite))
}
