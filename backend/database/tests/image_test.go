package tests

import (
	"context"
	"testing"

	"github.com/Slowers-team/Slowers-App/database"
	"github.com/Slowers-team/Slowers-App/testdata"
	"github.com/Slowers-team/Slowers-App/testutils"
	"github.com/stretchr/testify/suite"
)

type DbImageTestSuite struct {
	suite.Suite
	Db database.Database
}

func (s *DbImageTestSuite) SetupSuite() {
	s.Db = testutils.ConnectDB()
	s.Db.Clear()
}

func (s *DbImageTestSuite) TestAddImage() {
	for _, image := range testdata.GetImages() {
		imageToAdd := testdata.PrepareImageForAdding(image)
		addedImage, err := s.Db.AddImage(context.Background(), imageToAdd)

		s.Require().NoError(
			err,
			"AddImage() should not return an error",
		)
		s.NotZero(
			addedImage.ID,
			"ID for the added image should be non-zero",
		)
		s.Equal(
			image.FileFormat,
			addedImage.FileFormat,
			"wrong file format for the image returned from AddImage()",
		)
		s.Equal(
			image.Note,
			addedImage.Note,
			"wrong note for the image returned from AddImage()",
		)
		s.Equal(
			*image.Entity,
			*addedImage.Entity,
			"wrong entity for the image returned from AddImage()",
		)
		s.Equal(
			image.Owner,
			addedImage.Owner,
			"wrong owner for the image returned from AddImage()",
		)
	}
}

func (s *DbImageTestSuite) TestAddAndDeleteImage() {
	for _, image := range testdata.GetImages() {
		imageToAdd := testdata.PrepareImageForAdding(image)
		addedImage, _ := s.Db.AddImage(context.Background(), imageToAdd)
		anyDeleted, err := s.Db.DeleteImage(
			context.Background(), addedImage.ID,
		)

		s.True(
			anyDeleted,
			"DeleteImage() should return true",
		)
		s.NoError(
			err,
			"DeleteImage() should not return an error",
		)
	}
}

func (s *DbImageTestSuite) TestAddAndGetImageByEntity() {
	images := testdata.GetImages()
	imagesToAdd := []database.Image{
		testdata.PrepareImageForAdding(images[0]),
		testdata.PrepareImageForAdding(images[1]),
	}

	addedImage, _ := s.Db.AddImage(context.Background(), imagesToAdd[0])
	s.Db.AddImage(context.Background(), imagesToAdd[1])

	fetchedImages, err := s.Db.GetImagesByEntity(
		context.Background(), images[0].Entity.Hex(),
	)

	s.Require().NoError(
		err,
		"GetImagesByEntity() should not return an error",
	)
	s.Require().Len(
		fetchedImages,
		1,
		"GetImagesByEntity() should return a slice containing exactly 1 image",
	)
	s.Equal(
		addedImage.ID,
		fetchedImages[0].ID,
		"wrong ID for the image returned from GetImagesByEntity()",
	)
	s.Equal(
		images[0].FileFormat,
		fetchedImages[0].FileFormat,
		"wrong file format for the image returned from GetImagesByEntity()",
	)
	s.Equal(
		images[0].Note,
		fetchedImages[0].Note,
		"wrong note for the image returned from GetImagesByEntity()",
	)
	s.Equal(
		*images[0].Entity,
		*fetchedImages[0].Entity,
		"wrong entity for the image returned from GetImagesByEntity()",
	)
	s.Equal(
		images[0].Owner,
		fetchedImages[0].Owner,
		"wrong owner for the image returned from GetImagesByEntity()",
	)

}

func (s *DbImageTestSuite) TestAddAndGetImageByID() {
	images := testdata.GetImages()
	imagesToAdd := []database.Image{
		testdata.PrepareImageForAdding(images[0]),
		testdata.PrepareImageForAdding(images[1]),
	}

	addedImage, _ := s.Db.AddImage(context.Background(), imagesToAdd[0])
	s.Db.AddImage(context.Background(), imagesToAdd[1])

	fetchedImage, err := s.Db.GetImageByID(context.Background(), addedImage.ID)

	s.Require().NoError(
		err,
		"GetImageByID() should not return an error",
	)
	s.Equal(
		addedImage.ID,
		fetchedImage.ID,
		"wrong ID for the image returned from GetImageByID()",
	)
	s.Equal(
		images[0].FileFormat,
		fetchedImage.FileFormat,
		"wrong file format for the image returned from GetImageByID()",
	)
	s.Equal(
		images[0].Note,
		fetchedImage.Note,
		"wrong note for the image returned from GetImageByID()",
	)
	s.Equal(
		*images[0].Entity,
		*fetchedImage.Entity,
		"wrong entity for the image returned from GetImageByID()",
	)
	s.Equal(
		images[0].Owner,
		fetchedImage.Owner,
		"wrong owner for the image returned from GetImageByID()",
	)
}

func (s *DbImageTestSuite) TestClearFavoriteImageForFlower() {
	flower := testdata.GetFlowers()[0]
	flowerToAdd := testdata.PrepareFlowerForAdding(flower)
	addedFlower, _ := s.Db.AddFlower(context.Background(), flowerToAdd)

	image := testdata.GetImages()[0]
	imageToAdd := testdata.PrepareImageForAdding(image)
	addedImage, _ := s.Db.AddImage(context.Background(), imageToAdd)

	s.Db.SetFavoriteImage(
		context.Background(),
		*flower.Grower,
		addedFlower.ID,
		addedImage.ID,
		"flowers",
	)

	err := s.Db.ClearFavoriteImage(
		context.Background(),
		*flower.Grower,
		addedFlower.ID,
		"flowers",
	)

	s.Require().NoError(
		err,
		"ClearFavoriteImage() should not return an error",
	)

	fetchedFlowers, _ := s.Db.GetFlowers(context.Background())

	for _, fetchedFlower := range fetchedFlowers {
		if fetchedFlower.ID == addedFlower.ID {
			s.Zero(
				fetchedFlower.FavoriteImage,
				"fetched flower should not have favorite image",
			)
			break
		}
	}
}

func (s *DbImageTestSuite) TestClearFavoriteImageForSite() {
	site := testdata.GetRootSites()[0]
	siteToAdd := testdata.PrepareSiteForAdding(site)
	addedSite, _ := s.Db.AddSite(context.Background(), siteToAdd)

	image := testdata.GetImages()[1]
	imageToAdd := testdata.PrepareImageForAdding(image)
	addedImage, _ := s.Db.AddImage(context.Background(), imageToAdd)

	s.Db.SetFavoriteImage(
		context.Background(),
		*site.Owner,
		addedSite.ID,
		addedImage.ID,
		"sites",
	)

	err := s.Db.ClearFavoriteImage(
		context.Background(),
		*site.Owner,
		addedSite.ID,
		"sites",
	)

	s.Require().NoError(
		err,
		"ClearFavoriteImage() should not return an error",
	)

	fetchedSites, _ := s.Db.GetRootSites(
		context.Background(),
		*site.Owner,
	)

	for _, fetchedSite := range fetchedSites {
		if fetchedSite.ID == addedSite.ID {
			s.Zero(
				fetchedSite.FavoriteImage,
				"fetched site should not have favorite image",
			)
			break
		}
	}
}

func (s *DbImageTestSuite) TearDownTest() {
	s.Db.Clear()
}

func (s *DbImageTestSuite) TearDownSuite() {
	testutils.DisconnectDB(s.Db)
}

func TestDbImageTestSuite(t *testing.T) {
	suite.Run(t, new(DbImageTestSuite))
}
