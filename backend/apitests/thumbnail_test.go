package apitests

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"

	"github.com/stretchr/testify/suite"

	"github.com/Slowers-team/Slowers-App/database"
	"github.com/Slowers-team/Slowers-App/mocks"
	"github.com/Slowers-team/Slowers-App/testutils"
	"github.com/Slowers-team/Slowers-App/utils"

	"github.com/Slowers-team/Slowers-App/testdata"
)

type ThumbnailAPITestSuite struct {
	suite.Suite
	TestImages []database.Image
}

func (s *ThumbnailAPITestSuite) SetupSuite() {
	s.TestImages = testdata.GetImages()
}

// siirrä omaan testikansioon, apitesti ei hyvä
func (s *ThumbnailAPITestSuite) TestResizeImage() {
	// var image renamed to photo to suit image package
	photo := s.TestImages[0]
	filename := photo.ID.Hex() + "." + photo.FileFormat

	var err error

	filedata, err := os.ReadFile("../testdata/images/" + filename)
	if err != nil {
		log.Fatal(err)
	}

	fileReader := bytes.NewReader(filedata)

	savepath := "./thumbnails/" + photo.ID.Hex() + "." + photo.FileFormat
	createdThumbnail, err := os.Create(savepath)
	if err != nil {
		log.Fatal(err)
	}

	err = utils.ResizeImage(fileReader, createdThumbnail, photo.FileFormat, 100, 100)
	if err != nil {
		s.T().Fatalf("ResizeImage failed: %v", err)
	}

	// Decode the output image to verify the resize operation
	var decoded image.Image

	switch photo.FileFormat {
	case "png":
		decoded, err = png.Decode(createdThumbnail)
	case "jpg", "jpeg":
		decoded, err = jpeg.Decode(createdThumbnail)
	}
	if err != nil {
		s.T().Fatalf("failed to decode output image: %v", err)
	}

	if decoded.Bounds().Dx() != 100 || decoded.Bounds().Dy() != 100 {
		s.T().Fatalf("expected resized image to be 50x50, got %dx%d", decoded.Bounds().Dx(), decoded.Bounds().Dy())
	}
}

func (s *ThumbnailAPITestSuite) TestThumbnailDownload() {
	image := s.TestImages[0]
	filename := image.ID.Hex() + "." + image.FileFormat

	os.Mkdir("./thumbnails", 0775)

	filedata, err := os.ReadFile("../testdata/images/" + filename)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile("./thumbnails/"+filename, filedata, 0664); err != nil {
		log.Fatal(err)
	}

	testutils.RunTest(s.T(), testutils.TestCase{
		Description:  "GET /api/thumbnails/<filename>",
		Route:        "/api/thumbnails/" + filename,
		Method:       "GET",
		Body:         []byte{},
		ExpectedCode: 200,
		ExpectedBody: filedata,
		SetupMocks:   func(db *mocks.Database) {},
	})

	if err := os.RemoveAll("./thumbnails"); err != nil {
		log.Fatal(err)
	}
}
