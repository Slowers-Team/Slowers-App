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

func (s *DbImageTestSuite) TearDownTest() {
	s.Db.Clear()
}

func (s *DbImageTestSuite) TearDownSuite() {
	testutils.DisconnectDB(s.Db)
}

func TestDbImageTestSuite(t *testing.T) {
	suite.Run(t, new(DbImageTestSuite))
}
