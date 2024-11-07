package apitests

import (
	"bytes"
	"context"
	"fmt"
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
		ExpectedBody: utils.ImageToJSON(image),
		SetupMocks: func(db *mocks.Database) {
			db.EXPECT().AddImage(
				mock.Anything, testdata.GetImagesForAdding()[0],
			).RunAndReturn(func(ctx context.Context, newImage database.Image) (*database.Image, error) {
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

func TestImagesAPITestSuite(t *testing.T) {
	suite.Run(t, new(ImagesAPITestSuite))
}