package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/Slowers-team/Slowers-App/database"
	"github.com/Slowers-team/Slowers-App/handlers"
)

func flowersToJSON(flowers []database.Flower) string {
	flowersJSON, err := json.Marshal(flowers)
	if err != nil {
		log.Fatal(err)
	}
	return string(flowersJSON)
}

func idToJSON(id string) string {
	return "{\"id\": \"" + id + "\"}"
}

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
			expectedBody:  flowersToJSON(testFlowers),
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
			body:          idToJSON(testID),
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
		app := SetupAppAndSetAuthTo(false)

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
	}
}
