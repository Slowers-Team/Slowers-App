package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/Slowers-team/Slowers-App/application"
	"github.com/Slowers-team/Slowers-App/database"
	"github.com/Slowers-team/Slowers-App/handlers"
	"github.com/Slowers-team/Slowers-App/utils"
)

type testCase struct {
	description   string
	route         string
	method        string
	body          string
	expectedError bool
	expectedCode  int
	expectedBody  string
	setupMocks    func(db *database.MockDatabase)
}

func TestFlowersRoute(t *testing.T) {
	testFlowers := []database.Flower{
		{
			Name: "sunflower",
			LatinName: "Helianthus annuus",
			AddedTime: time.Date(2024, 9, 30, 21, 11, 54, 0, time.UTC),
		},
	}

	testID := "842af389e234e768923475bc"

	tests := []testCase{
		{
			description:   "\"GET /api/flowers\" without error",
			route:         "/api/flowers",
			method:        "GET",
			body:          "",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  utils.FlowersToJSON(testFlowers),
			setupMocks:    func(db *database.MockDatabase) {
				db.On(
					"GetFlowers", mock.Anything,
				).Return(
					testFlowers, nil,
				).Once()
			},
		},
		{
			description:   "\"GET /api/flowers\" with error",
			route:         "/api/flowers",
			method:        "GET",
			body:          "",
			expectedError: false,
			expectedCode:  500,
			expectedBody:  "Database error",
			setupMocks:    func(db *database.MockDatabase) {
				db.On(
					"GetFlowers", mock.Anything,
				).Return(
					[]database.Flower{}, errors.New("Database error"),
				).Once()
			},
		},
		{
			description:   "DELETE /api/flowers/<id>",
			route:         "/api/flowers/" + testID,
			method:        "DELETE",
			body:          utils.IDToJSON(testID),
			expectedError: false,
			expectedCode:  204,
			expectedBody:  "",
			setupMocks:    func(db *database.MockDatabase) {
				db.On(
					"DeleteFlower", mock.Anything, testID,
				).Return(
					true, nil,
				).Once()
			},
		},

	}

	runTests(t, tests)
}

func runTests(t *testing.T, tests []testCase) {
	for i, test := range tests {
		app := application.SetupAndSetAuthTo(false)

		db := new(database.MockDatabase)
		handlers.SetDatabase(db)

		test.setupMocks(db)

		req, _ := http.NewRequest(test.method, test.route, strings.NewReader(test.body))

		res, err := app.Test(req, -1)

		db.AssertExpectations(t)

		assert.Equalf(t, test.expectedError, err != nil, test.description)

		if test.expectedError {
			continue
		}

		assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)

		body, err := io.ReadAll(res.Body)

		assert.Nilf(t, err, test.description)

		assert.Equalf(t, test.expectedBody, string(body), test.description)

		if t.Failed() {
			fmt.Println("Test", i, ":", test.description, ": FAIL")
		} else {
			fmt.Println("Test", i, ":", test.description, ": PASS")
		}
	}
}
