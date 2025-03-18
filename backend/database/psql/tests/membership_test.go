//go:build sql
// +build sql

package tests

import (
	"testing"

	"github.com/stretchr/testify/suite"

	database "github.com/Slowers-team/Slowers-App/database/psql"
	"github.com/Slowers-team/Slowers-App/testdataPsql"
	"github.com/Slowers-team/Slowers-App/testutils"
)

type DbMembershipTestSuite struct {
	suite.Suite
	Db             database.Database
	TestUser       database.User
	TestBusiness   database.Business
	TestMembership database.Membership
}

func (s *DbMembershipTestSuite) SetupSuite() {
	s.Db = testutils.ConnectSQLDB()
	s.Db.Clear()
	s.TestUser = testdataPsql.GetUsers()[0]
	s.TestBusiness = testdataPsql.GetBusinesses()[0]
}

// Here some tests

func (s *DbMembershipTestSuite) TearDownTest() {
	s.Db.Clear()
}

func (s *DbMembershipTestSuite) TearDownSuite() {
	testutils.DisconnectSQLDB(s.Db)
}

func TestDbMembershipTestSuite(t *testing.T) {
	suite.Run(t, new(DbMembershipTestSuite))
}
