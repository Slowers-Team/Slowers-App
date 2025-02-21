package apitests

import (
	"bytes"
	"context"
	"fmt"

	// "image"
	// "image/jpeg"
	// "image/png"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/textproto"
	"os"
	"strings"
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

func (s *ImagesAPITestSuite) TestImageDownload() {
	image := s.Images[0]
	filename := image.ID.Hex() + "." + image.FileFormat

	os.Mkdir("./images", 0775)

	filedata, err := ioutil.ReadFile("../testdata/images/" + filename)
	if err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile("./images/"+filename, filedata, 0664); err != nil {
		log.Fatal(err)
	}

	testutils.RunTest(s.T(), testutils.TestCase{
		Description:  "GET /api/images/<filename>",
		Route:        "/api/images/" + filename,
		Method:       "GET",
		Body:         []byte{},
		ExpectedCode: 200,
		ExpectedBody: filedata,
		SetupMocks:   func(db *mocks.Database) {},
	})

	if err := os.RemoveAll("./images"); err != nil {
		log.Fatal(err)
	}
}

func (s *ImagesAPITestSuite) TestImageUpload() {
	image := s.Images[0]
	filename := image.ID.Hex() + "." + image.FileFormat

	var fieldWriter io.Writer
	var formData bytes.Buffer
	formWriter := multipart.NewWriter(&formData)
	var err error

	if fieldWriter, err = formWriter.CreateFormField("note"); err != nil {
		log.Fatal(err)
	}
	if _, err := io.Copy(fieldWriter, strings.NewReader(image.Note)); err != nil {
		log.Fatal(err)
	}

	if fieldWriter, err = formWriter.CreateFormField("entity"); err != nil {
		log.Fatal(err)
	}
	if _, err := io.Copy(fieldWriter, strings.NewReader(image.Entity.Hex())); err != nil {
		log.Fatal(err)
	}

	header := make(textproto.MIMEHeader)
	header.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, "image", filename))
	if image.FileFormat == "jpg" {
		header.Set("Content-Type", "image/jpeg")
	} else if image.FileFormat == "png" {
		header.Set("Content-Type", "image/png")
	} else {
		log.Fatal("apitests/image_test.go: TestImageUpload(): Unknown image format \"" + image.FileFormat + "\"")
	}
	if fieldWriter, err = formWriter.CreatePart(header); err != nil {
		log.Fatal(err)
	}
	file, err := os.Open("../testdata/images/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := io.Copy(fieldWriter, file); err != nil {
		log.Fatal(err)
	}

	formWriter.Close()

	os.Mkdir("./images", 0775)

	testutils.RunTest(s.T(), testutils.TestCase{
		Description:  "POST /api/images",
		Route:        "/api/images",
		Method:       "POST",
		ContentType:  formWriter.FormDataContentType(),
		Body:         formData.Bytes(),
		ExpectedCode: 201,
		ExpectedBody: utils.ToJSON(image),
		SetupMocks: func(db *mocks.Database) {
			db.EXPECT().AddImage(
				mock.Anything, testdata.GetImagesForAdding()[0], "images",
			).RunAndReturn(func(ctx context.Context, newImage database.Image, target string) (*database.Image, error) {
				return &database.Image{
					ID:         image.ID,
					FileFormat: newImage.FileFormat,
					Note:       newImage.Note,
					Entity:     newImage.Entity,
					Owner:      newImage.Owner,
				}, nil
			}).Once()

			db.EXPECT().AddImage(
				mock.Anything, testdata.GetImagesForAdding()[0], "thumbnails",
			).RunAndReturn(func(ctx context.Context, newImage database.Image, target string) (*database.Image, error) {
				return &database.Image{
					ID:         image.ID,
					FileFormat: newImage.FileFormat,
					Note:       newImage.Note,
					Entity:     newImage.Entity,
					Owner:      newImage.Owner,
				}, nil
			}).Once()
		},
	})

	fileDataReceived, err := ioutil.ReadFile("../testdata/images/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	fileDataSaved, err := ioutil.ReadFile("./images/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	s.Equal(
		fileDataReceived,
		fileDataSaved,
		"Saved file is different from the file received",
	)

	if err := os.RemoveAll("./images"); err != nil {
		log.Fatal(err)
	}
}

// func (s *ImagesAPITestSuite) TestResizeImage() {
// 	// var image renamed to photo to suit image package
// 	photo := s.Images[0]
// 	filename := photo.ID.Hex() + "." + photo.FileFormat

// 	var fieldWriter io.Writer
// 	var formData bytes.Buffer
// 	formWriter := multipart.NewWriter(&formData)
// 	var err error

// 	if fieldWriter, err = formWriter.CreateFormField("entity"); err != nil {
// 		log.Fatal(err)
// 	}
// 	if _, err := io.Copy(fieldWriter, strings.NewReader(photo.Entity.Hex())); err != nil {
// 		log.Fatal(err)
// 	}

// 	header := make(textproto.MIMEHeader)
// 	header.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, "image", filename))
// 	if photo.FileFormat == "jpg" {
// 		header.Set("Content-Type", "image/jpeg")
// 	} else if photo.FileFormat == "png" {
// 		header.Set("Content-Type", "image/png")
// 	} else {
// 		log.Fatal("apitests/image_test.go: TestImageUpload(): Unknown image format \"" + photo.FileFormat + "\"")
// 	}
// 	if fieldWriter, err = formWriter.CreatePart(header); err != nil {
// 		log.Fatal(err)
// 	}
// 	file, err := os.Open("../testdata/images/" + filename)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if _, err := io.Copy(fieldWriter, file); err != nil {
// 		log.Fatal(err)
// 	}

// 	formWriter.Close()

// 	// Read the first few bytes to check the SOI marker
// 	buffer := make([]byte, 2)
// 	if _, err := file.Read(buffer); err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("First two bytes: %x\n", buffer)

// 	var output bytes.Buffer
// 	err = utils.ResizeImage(&formData, &output, photo.FileFormat, 100, 100)
// 	if err != nil {
// 		s.T().Fatalf("ResizeImage failed: %v", err)
// 	}

// 	// Decode the output image to verify the resize operation
// 	var decoded image.Image

// 	switch photo.FileFormat {
// 	case "png":
// 		decoded, err = png.Decode(&output)
// 	case "jpg", "jpeg":
// 		decoded, err = jpeg.Decode(&output)
// 	}
// 	if err != nil {
// 		s.T().Fatalf("failed to decode output image: %v", err)
// 	}

// 	if decoded.Bounds().Dx() != 100 || decoded.Bounds().Dy() != 100 {
// 		s.T().Fatalf("expected resized image to be 50x50, got %dx%d", decoded.Bounds().Dx(), decoded.Bounds().Dy())
// 	}
// }

func (s *ImagesAPITestSuite) TestFetchingImagesByEntity() {
	testutils.RunTest(s.T(), testutils.TestCase{
		Description:  "GET /api/images/entity/:entityID",
		Route:        "/api/images/entity/" + s.Images[0].Entity.Hex(),
		Method:       "GET",
		Body:         []byte{},
		ExpectedCode: 200,
		ExpectedBody: utils.ToJSON([]database.Image{s.Images[0]}),
		SetupMocks: func(db *mocks.Database) {
			db.EXPECT().GetImagesByEntity(
				mock.Anything, s.Images[0].Entity.Hex(),
			).Return(
				[]database.Image{s.Images[0]}, nil,
			).Once()
		},
	})
}

func (s *ImagesAPITestSuite) TestDeletingImage() {
	image := s.Images[0]
	filename := image.ID.Hex() + "." + image.FileFormat

	os.Mkdir("./images", 0775)

	filedata, err := ioutil.ReadFile("../testdata/images/" + filename)
	if err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile("./images/"+filename, filedata, 0664); err != nil {
		log.Fatal(err)
	}

	testutils.RunTest(s.T(), testutils.TestCase{
		Description:  "DELETE /api/images/:id",
		Route:        "/api/images/" + image.ID.Hex(),
		Method:       "DELETE",
		Body:         []byte{},
		ExpectedCode: 204,
		ExpectedBody: []byte{},
		SetupMocks: func(db *mocks.Database) {
			db.EXPECT().DeleteImage(
				mock.Anything, image.ID, "images",
			).Return(
				true, nil,
			).Once()
		},
	})

	s.NoFileExists("./images/"+filename, "File should have been deleted")

	if err := os.RemoveAll("./images"); err != nil {
		log.Fatal(err)
	}
}

func (s *ImagesAPITestSuite) TestGetImageByID() {
	image := s.Images[0]
	filename := image.ID.Hex() + "." + image.FileFormat

	os.Mkdir("./images", 0775)

	filedata, err := ioutil.ReadFile("../testdata/images/" + filename)
	if err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile("./images/"+filename, filedata, 0664); err != nil {
		log.Fatal(err)
	}

	testutils.RunTest(s.T(), testutils.TestCase{
		Description:  "GET /api/images/id/<id>",
		Route:        "/api/images/id/" + image.ID.Hex(),
		Method:       "GET",
		ContentType:  "application/json",
		Body:         []byte{},
		ExpectedCode: 200,
		ExpectedBody: filedata,
		SetupMocks: func(db *mocks.Database) {
			db.EXPECT().GetImageByID(
				mock.Anything, image.ID, "images",
			).Return(
				&image, nil,
			).Once()
		},
	})

	if err := os.RemoveAll("./images"); err != nil {
		log.Fatal(err)
	}
}

func (s *ImagesAPITestSuite) TestClearFavoriteImageOfFlower() {
	flower := testdata.GetTestFlowers()[0]
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
	flower := testdata.GetTestFlowers()[0]
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
