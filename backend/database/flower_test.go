package database

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
)

type DbFlowerTestSuite struct {
	suite.Suite
	Db Database
}

func (suite *DbFlowerTestSuite) SetupTest() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Println("No .env file found")
	}

	databaseURI := os.Getenv("MONGODB_URI")
	if databaseURI == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable.")
	}

	suite.Db = NewMongoDatabase(databaseURI)
	if err := suite.Db.Connect("SlowersTest"); err != nil {
		log.Fatal(err)
	}
}

func (suite *DbFlowerTestSuite) TestAddFlower() {
	db := suite.Db

	flower := Flower{
		Name: "sunflower",
		LatinName: "Helianthus annuus",
		AddedTime: time.Date(2024, 9, 30, 21, 11, 54, 0, time.UTC),
	}

	var createdFlower *Flower
	var err error
	createdFlower, err = db.AddFlower(context.Background(), flower)

	suite.NoError(err, "AddFlower() should not return an error")
	suite.Equal(createdFlower.Name, flower.Name, "wrong name for the flower returned from AddFlower()")
	suite.Equal(createdFlower.LatinName, flower.LatinName, "wrong latin name for the flower returned from AddFlower()")
	suite.Equal(createdFlower.AddedTime, flower.AddedTime, "wrong AddedTime for the flower returned from AddFlower()")
	suite.NotEmpty(createdFlower.ID, "ID for the created flower should not be empty")
}

func (suite *DbFlowerTestSuite) TearDownTest() {
	suite.Db.Clear()
	err := suite.Db.Disconnect()
	if err != nil {
		log.Fatal(err)
	}
}

func TestDbFlowerTestSuite(t *testing.T) {
	suite.Run(t, new(DbFlowerTestSuite))
}
