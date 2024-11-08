package testutils

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Slowers-team/Slowers-App/application"
	"github.com/Slowers-team/Slowers-App/handlers"
	"github.com/Slowers-team/Slowers-App/mocks"
)

type TestCase struct {
	Description      string
	Route            string
	Method           string
	ContentType      string
	Body             []byte
	ExpectedCode     int
	ExpectedBody     []byte
	ExpectedBodyFunc func(body []byte)
	SetupMocks       func(db *mocks.Database)
}

func RunTest(t *testing.T, test TestCase) {
	app := application.SetupAndSetAuthTo(false)
	db := mocks.NewDatabase(t)
	handlers.SetDatabase(db)

	test.SetupMocks(db)

	req, _ := http.NewRequest(
		test.Method,
		test.Route,
		bytes.NewReader(test.Body),
	)
	req.Header.Add("Content-Type", test.ContentType)
	res, err := app.Test(req, -1)

	db.AssertExpectations(t)

	assert.Equalf(t, test.ExpectedCode, res.StatusCode, test.Description)

	body, err := io.ReadAll(res.Body)
	assert.Nilf(t, err, test.Description)
	if test.ExpectedBodyFunc == nil {
		assert.Equalf(t, test.ExpectedBody, body, test.Description)
	} else {
		test.ExpectedBodyFunc(body)
	}
}
