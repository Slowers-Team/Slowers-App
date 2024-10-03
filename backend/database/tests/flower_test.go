package tests

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"

	"github.com/Slowers-team/Slowers-App/database"
	"github.com/Slowers-team/Slowers-App/testdata"
)

type DbFlowerTestSuite struct {
	suite.Suite
	Db database.Database
}

func (suite *DbFlowerTestSuite) SetupTest() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Println("No .env file found")
	}

	databaseURI := os.Getenv("MONGODB_URI")
	if databaseURI == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable.")
	}

	suite.Db = database.NewMongoDatabase(databaseURI)
	if err := suite.Db.Connect("SlowersTest"); err != nil {
		log.Fatal(err)
	}
}

func (suite *DbFlowerTestSuite) TestAddFlower() {
	db := suite.Db

	flower := testdata.GetTestFlowers()[0]

	var createdFlower *database.Flower
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
