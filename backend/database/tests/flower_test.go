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
	Db database.Database
}

func (s *DbFlowerTestSuite) SetupSuite() {
	s.Db = testutils.ConnectDB()
	s.Db.Clear()
}

func (s *DbFlowerTestSuite) TestAddFlower() {
	flower := testdata.GetTestFlowers()[0]
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
	s.NotEmpty(
		createdFlower.ID,
		"ID for the created flower should not be empty",
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
