package testutils

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Slowers-team/Slowers-App/application"
	"github.com/Slowers-team/Slowers-App/databases/mongo"
	"github.com/Slowers-team/Slowers-App/databases/sql"
	"github.com/Slowers-team/Slowers-App/handlers"
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
	SetupSql         func(sql sql.Database)
	SetupMongo       func(mongo mongo.Database)
}

func RunTest(t *testing.T, test TestCase) {
	app := application.SetupAndSetAuthTo(false) //TODO: Add Psql toggle
	mongo := ConnectMongoDB()
	sql := ConnectSqlDB()

	handlers.SetDatabases(mongo, sql)
	test.SetupMongo(mongo)
	test.SetupSql(sql)

	req, _ := http.NewRequest(
		test.Method,
		test.Route,
		bytes.NewReader(test.Body),
	)
	req.Header.Add("Content-Type", test.ContentType)
	res, err := app.Test(req, -1)

	assert.Equal(t, test.ExpectedCode, res.StatusCode, test.Description)

	body, err := io.ReadAll(res.Body)
	assert.NoError(t, err, test.Description)
	if test.ExpectedBodyFunc == nil {
		assert.Equal(t, test.ExpectedBody, body, test.Description)
	} else {
		test.ExpectedBodyFunc(body)
	}
}
