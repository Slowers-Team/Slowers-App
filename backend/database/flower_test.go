package database

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type DbFlowerTestSuite struct {
	suite.Suite
	DbClient *DatabaseClient
}

func (suite *DbFlowerTestSuite) SetupTest() {
	databaseURI := os.Getenv("MONGODB_URI")
	if databaseURI == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable.")
	}

	var err error
	suite.DbClient, err = Connect(databaseURI, "SlowersTest")
	if err != nil {
		log.Fatal(err)
	}
}

func (suite *DbFlowerTestSuite) TestAddFlower() {
	flower := Flower{
		Name: "sunflower",
		LatinName: "Helianthus annuus",
		AddedTime: time.Date(2024, 9, 30, 21, 11, 54, 0, time.UTC),
	}

	var createdFlower *Flower
	var err error
	createdFlower, err = AddFlower(context.Background(), flower)

	suite.NoError(err, "AddFlower() should not return an error")
	suite.Equal(createdFlower.Name, flower.Name, "wrong name for the flower returned from AddFlower()")
	suite.Equal(createdFlower.LatinName, flower.LatinName, "wrong latin name for the flower returned from AddFlower()")
	suite.Equal(createdFlower.AddedTime, flower.AddedTime, "wrong AddedTime for the flower returned from AddFlower()")
	suite.NotEmpty(createdFlower.ID, "ID for the created flower should not be empty")
}

func (suite *DbFlowerTestSuite) TearDownTest() {
	Clear()
	err := Disconnect(suite.DbClient)
	if err != nil {
		log.Fatal(err)
	}
}

func TestDbFlowerTestSuite(t *testing.T) {
	suite.Run(t, new(DbFlowerTestSuite))
}
