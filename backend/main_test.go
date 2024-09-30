package main

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/Slowers-team/Slowers-App/db"
)

type FlowersRouteTestSuite struct {
	suite.Suite
	DbClient *db.DatabaseClient
}

func (suite *FlowersRouteTestSuite) SetupTest() {
	databaseURI := os.Getenv("MONGODB_URI")
	if databaseURI == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable.")
	}

	var err error
	suite.DbClient, err = db.Connect(databaseURI, "SlowersTest")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Clear()
	if err != nil {
		log.Fatal(err)
	}
}

func (suite *FlowersRouteTestSuite) TestFlowersRoute() {
	tests := []struct {
		description   string
		route         string
		expectedError bool
		expectedCode  int
		expectedBody  string
	}{
		{
			description:   "GET flowers",
			route:         "/api/flowers",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "OK",
		},
	}

	app := Setup()

	for _, test := range tests {
		req, _ := http.NewRequest(
			"GET",
			test.route,
			nil,
		)

		res, err := app.Test(req, -1)

		suite.Equalf(test.expectedError, err != nil, test.description)

		if test.expectedError {
			continue
		}

		suite.Equalf(test.expectedCode, res.StatusCode, test.description)
	}
}

func (suite *FlowersRouteTestSuite) TearDownTest() {
	if err := db.Clear(); err != nil {
		log.Fatal(err)
	}

	if err := db.Disconnect(suite.DbClient); err != nil {
		log.Fatal(err)
	}
}

func TestFlowersRouteTestSuite(t *testing.T) {
	suite.Run(t, new(FlowersRouteTestSuite))
}
