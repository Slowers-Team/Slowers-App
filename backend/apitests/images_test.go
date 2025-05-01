//go:build api
// +build api

package apitests

import (
	"log"
	"testing"

	"github.com/Slowers-team/Slowers-App/database"
	"github.com/Slowers-team/Slowers-App/mocks"
	"github.com/Slowers-team/Slowers-App/testdata"
	"github.com/Slowers-team/Slowers-App/testutils"
	"github.com/Slowers-team/Slowers-App/utils"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type ImagesAPITestSuite struct {
	suite.Suite
	Images []database.Image
}

func (s *ImagesAPITestSuite) SetupSuite() {
	s.Images = testdata.GetImages()
}

func (s *ImagesAPITestSuite) TestClearFavoriteImageOfFlower() {
	flower := testdata.GetFlowers()[0]
	entityType := "flower"
	collection := "flowers"

	b := make(map[string]string)
	b["EntityID"] = flower.ID.Hex()
	b["EntityType"] = entityType

	body := utils.ToJSON(b)
	log.Println("body ", string(body))

	testutils.RunTest(s.T(), testutils.TestCase{
		Description:  "POST /api/images/clearfavorite",
		Route:        "/api/images/clearfavorite",
		Method:       "POST",
		ContentType:  "application/json",
		Body:         body,
		ExpectedCode: 200,
		ExpectedBody: []byte{},
		SetupMocks: func(db *mocks.Database) {
			db.EXPECT().ClearFavoriteImage(
				mock.Anything, *flower.Grower, flower.ID, collection,
			).Return(
				nil,
			).Once()
		},
	})
}

func (s *ImagesAPITestSuite) TestClearFavoriteImageOfSite() {
	site := testdata.GetRootSites()[0]
	entityType := "site"
	collection := "sites"

	b := make(map[string]string)
	b["EntityID"] = site.ID.Hex()
	b["EntityType"] = entityType

	body := utils.ToJSON(b)
	log.Println("body ", string(body))

	testutils.RunTest(s.T(), testutils.TestCase{
		Description:  "POST /api/images/clearfavorite",
		Route:        "/api/images/clearfavorite",
		Method:       "POST",
		ContentType:  "application/json",
		Body:         body,
		ExpectedCode: 200,
		ExpectedBody: []byte{},
		SetupMocks: func(db *mocks.Database) {
			db.EXPECT().ClearFavoriteImage(
				mock.Anything, *site.Owner, site.ID, collection,
			).Return(
				nil,
			).Once()
		},
	})
}

func (s *ImagesAPITestSuite) TestClearFavoriteImageFailsWithIncorrectEntity() {
	flower := testdata.GetFlowers()[0]
	entityType := "incorrect"

	b := make(map[string]string)
	b["EntityID"] = flower.ID.Hex()
	b["EntityType"] = entityType

	body := utils.ToJSON(b)
	log.Println("body ", string(body))

	testutils.RunTest(s.T(), testutils.TestCase{
		Description:  "POST /api/images/clearfavorite",
		Route:        "/api/images/clearfavorite",
		Method:       "POST",
		ContentType:  "application/json",
		Body:         body,
		ExpectedCode: 400,
		ExpectedBody: []byte("Invalid EntityType: incorrect"),
		SetupMocks: func(db *mocks.Database) {
		},
	})
}

func TestImagesAPITestSuite(t *testing.T) {
	suite.Run(t, new(ImagesAPITestSuite))
}
