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
	for _, image := range testdata.GetImagesForAdding() {
		createdImage, err := s.Db.AddImage(context.Background(), image)

		s.NoError(
			err,
			"AddImage() should not return an error",
		)
		s.NotZero(
			createdImage.ID,
			"ID for the added image should be non-zero",
		)
		s.Equal(
			createdImage.FileFormat,
			image.FileFormat,
			"wrong file format for the image returned from AddImage()",
		)
		s.Equal(
			createdImage.Note,
			image.Note,
			"wrong note for the image returned from AddImage()",
		)
		s.Equal(
			*createdImage.Entity,
			*image.Entity,
			"wrong entity for the image returned from AddImage()",
		)
		s.Equal(
			createdImage.Owner,
			image.Owner,
			"wrong owner for the image returned from AddImage()",
		)
	}
}

func (s *DbImageTestSuite) TestAddAndDeleteImage() {
	for _, image := range testdata.GetImagesForAdding() {
		createdImage, _ := s.Db.AddImage(context.Background(), image)
		anyDeleted, err := s.Db.DeleteImage(context.Background(), createdImage.ID)

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
	imagesForAdding := testdata.GetImagesForAdding()

	createdImage, _ := s.Db.AddImage(context.Background(), imagesForAdding[0])
	s.Db.AddImage(context.Background(), imagesForAdding[1])

	fetchedImages, err := s.Db.GetImagesByEntity(
		context.Background(), imagesForAdding[0].Entity.Hex(),
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
		createdImage.ID,
		fetchedImages[0].ID,
		"wrong ID for the image returned from GetImagesByEntity()",
	)
	s.Equal(
		createdImage.FileFormat,
		fetchedImages[0].FileFormat,
		"wrong file format for the image returned from GetImagesByEntity()",
	)
	s.Equal(
		createdImage.Note,
		fetchedImages[0].Note,
		"wrong note for the image returned from GetImagesByEntity()",
	)
	s.Equal(
		*createdImage.Entity,
		*fetchedImages[0].Entity,
		"wrong entity for the image returned from GetImagesByEntity()",
	)
	s.Equal(
		createdImage.Owner,
		fetchedImages[0].Owner,
		"wrong owner for the image returned from GetImagesByEntity()",
	)

}

func (s *DbImageTestSuite) TestAddAndGetImageByID() {
	imagesForAdding := testdata.GetImagesForAdding()

	createdImage, _ := s.Db.AddImage(context.Background(), imagesForAdding[0])
	s.Db.AddImage(context.Background(), imagesForAdding[1])

	fetchedImage, err := s.Db.GetImageByID(
		context.Background(), createdImage.ID)

	s.Require().NoError(
		err,
		"GetImageByID() should not return an error",
	)
	s.Equal(
		createdImage.ID,
		fetchedImage.ID,
		"wrong ID for the image returned from GetImageByID()",
	)
	s.Equal(
		createdImage.FileFormat,
		fetchedImage.FileFormat,
		"wrong file format for the image returned from GetImageByID()",
	)
	s.Equal(
		createdImage.Note,
		fetchedImage.Note,
		"wrong note for the image returned from GetImageByID()",
	)
	s.Equal(
		*createdImage.Entity,
		*fetchedImage.Entity,
		"wrong entity for the image returned from GetImageByID()",
	)
	s.Equal(
		createdImage.Owner,
		fetchedImage.Owner,
		"wrong owner for the image returned from GetImageByID()",
	)
}

func (s *DbImageTestSuite) TestClearFavoriteImageForFlower() {
	flowerToAdd := testdata.GetTestFlowers()[0]
	flowerToAdd.ID = database.NilObjectID
	flowerToAdd.FavoriteImage = ""
	addedFlower, _ := s.Db.AddFlower(context.Background(), flowerToAdd)

	imagesForAdding := testdata.GetImagesForAdding()
	addedImage, _ := s.Db.AddImage(context.Background(), imagesForAdding[0])

	s.Db.SetFavoriteImage(
		context.Background(),
		*flowerToAdd.Grower,
		addedFlower.ID,
		addedImage.ID,
		"flowers",
	)

	err := s.Db.ClearFavoriteImage(
		context.Background(),
		*flowerToAdd.Grower,
		addedFlower.ID,
		"flowers",
	)

	s.Require().NoError(
		err,
		"ClearFavoriteImage() should not return an error",
	)

	fetchedFlowers, _ := s.Db.GetFlowers(context.Background())

	for _, flower := range fetchedFlowers {
		if flower.ID == addedFlower.ID {
			s.Zero(
				flower.FavoriteImage,
				"fetched flower should not have favorite image",
			)
		}
	}
}

func (s *DbImageTestSuite) TestClearFavoriteImageForSite() {
	siteToAdd := testdata.GetRootSites()[0]
	siteToAdd.ID = database.NilObjectID
	siteToAdd.FavoriteImage = ""
	addedSite, _ := s.Db.AddSite(context.Background(), siteToAdd)

	imagesForAdding := testdata.GetImagesForAdding()
	addedImage, _ := s.Db.AddImage(context.Background(), imagesForAdding[1])

	s.Db.SetFavoriteImage(
		context.Background(),
		*siteToAdd.Owner,
		addedSite.ID,
		addedImage.ID,
		"sites",
	)

	err := s.Db.ClearFavoriteImage(
		context.Background(),
		*siteToAdd.Owner,
		addedSite.ID,
		"sites",
	)

	s.Require().NoError(
		err,
		"ClearFavoriteImage() should not return an error",
	)

	fetchedSites, _ := s.Db.GetRootSites(
		context.Background(),
		*siteToAdd.Owner,
	)

	for _, site := range fetchedSites {
		if site.ID == addedSite.ID {
			s.Zero(
				site.FavoriteImage,
				"fetched site should not have favorite image",
			)
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
