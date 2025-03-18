//go:build sql
// +build sql

package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"

	database "github.com/Slowers-team/Slowers-App/database/psql"
	"github.com/Slowers-team/Slowers-App/testdataPsql"
	"github.com/Slowers-team/Slowers-App/testutils"
	"github.com/Slowers-team/Slowers-App/utils"
)

type DbUserTestSuite struct {
	suite.Suite
	Db       database.Database
	TestUser database.User
}

func (s *DbUserTestSuite) SetupSuite() {
	s.Db = testutils.ConnectSQLDB()
	s.Db.Clear()
	s.TestUser = testdataPsql.GetUsers()[0]
}

func (s *DbUserTestSuite) TestCreateUser() {
	hashedPassword, _ := utils.HashPassword(s.TestUser.Password)
	user := database.User{
		Username: s.TestUser.Username,
		Email:    s.TestUser.Email,
		Password: hashedPassword,
		IsActive: s.TestUser.IsActive,
		IsAdmin:  s.TestUser.IsAdmin,
	}
	newUser, err := s.Db.CreateUser(context.Background(), user)

	s.NoError(
		err,
		"CreateUser() should not return an error",
	)
	s.NotZero(
		newUser.ID,
		"new user should have non-zero ID",
	)
	s.Equal(
		newUser.Username,
		s.TestUser.Username,
		"wrong username for new user",
	)
	s.Equal(
		newUser.Email,
		s.TestUser.Email,
		"wrong email for new user",
	)
	s.NoError(
		bcrypt.CompareHashAndPassword([]byte(newUser.Password), []byte(s.TestUser.Password)),
		"wrong password for new user",
	)
	s.Equal(
		newUser.IsActive,
		s.TestUser.IsActive,
		"wrong active status for new user",
	)
	s.Equal(
		newUser.IsAdmin,
		s.TestUser.IsAdmin,
		"wrong admin status for new user",
	)
}

func (s *DbUserTestSuite) TearDownTest() {
	s.Db.Clear()
}

func (s *DbUserTestSuite) TearDownSuite() {
	testutils.DisconnectSQLDB(s.Db)
}

func TestDbUserTestSuite(t *testing.T) {
	suite.Run(t, new(DbUserTestSuite))
}
