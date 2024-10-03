package testing

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Slowers-team/Slowers-App/application"
	"github.com/Slowers-team/Slowers-App/database"
	"github.com/Slowers-team/Slowers-App/handlers"
)

type TestCase struct {
	Description   string
	Route         string
	Method        string
	Body          string
	ExpectedError bool
	ExpectedCode  int
	ExpectedBody  string
	SetupMocks    func(db *database.MockDatabase)
}

func RunTests(t *testing.T, tests []TestCase) {
	for i, test := range tests {
		app := application.SetupAndSetAuthTo(false)

		db := new(database.MockDatabase)
		handlers.SetDatabase(db)

		test.SetupMocks(db)

		req, _ := http.NewRequest(test.Method, test.Route, strings.NewReader(test.Body))

		res, err := app.Test(req, -1)

		db.AssertExpectations(t)

		assert.Equalf(t, test.ExpectedError, err != nil, test.Description)

		if test.ExpectedError {
			continue
		}

		assert.Equalf(t, test.ExpectedCode, res.StatusCode, test.Description)

		body, err := io.ReadAll(res.Body)

		assert.Nilf(t, err, test.Description)

		assert.Equalf(t, test.ExpectedBody, string(body), test.Description)

		if t.Failed() {
			fmt.Println("Test", i, ":", test.Description, ": FAIL")
		} else {
			fmt.Println("Test", i, ":", test.Description, ": PASS")
		}
	}
}
