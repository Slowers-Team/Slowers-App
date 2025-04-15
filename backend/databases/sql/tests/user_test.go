package tests

import (
	"context"
	"strconv"
	"testing"

	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"

	"github.com/Slowers-team/Slowers-App/databases/sql"
	"github.com/Slowers-team/Slowers-App/testdata"
	"github.com/Slowers-team/Slowers-App/testutils"
	"github.com/Slowers-team/Slowers-App/utils"
)

type DbUserTestSuite struct {
	suite.Suite
	SqlDb    sql.Database
	TestUser sql.User
}

func (s *DbUserTestSuite) SetupSuite() {
	s.SqlDb = testutils.ConnectSqlDB()
	s.SqlDb.Clear()
	s.TestUser = testdata.GetUsers()[0]
}

func (s *DbUserTestSuite) TestCreateUser() {
	hashedPassword, _ := utils.HashPassword(s.TestUser.Password)
	user := sql.User{
		Username: s.TestUser.Username,
		Email:    s.TestUser.Email,
		Password: hashedPassword,
		IsActive: s.TestUser.IsActive,
		IsAdmin:  s.TestUser.IsAdmin,
	}
	newUser, err := s.SqlDb.CreateUser(context.Background(), user)

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

func (s *DbUserTestSuite) TestGetUserByEmail() {
	hashedPassword, _ := utils.HashPassword(s.TestUser.Password)
	user := sql.User{
		Username: s.TestUser.Username,
		Email:    s.TestUser.Email,
		Password: hashedPassword,
		IsActive: s.TestUser.IsActive,
		IsAdmin:  s.TestUser.IsAdmin,
	}
	newUser, err := s.SqlDb.CreateUser(context.Background(), user)
	s.NoError(
		err,
		"CreateUser() should not return an error",
	)

	fetchedUser, err := s.SqlDb.GetUserByEmail(context.Background(), newUser.Email)
	s.NoError(
		err,
		"GetUserByEmail() should not return an error",
	)
	s.NotZero(
		fetchedUser.ID,
		"fetched user should have non-zero ID",
	)
	s.Equal(
		fetchedUser.Username,
		newUser.Username,
		"wrong username for fetched user",
	)
	s.Equal(
		fetchedUser.Email,
		newUser.Email,
		"wrong email for fetched user",
	)
	s.NoError(
		bcrypt.CompareHashAndPassword([]byte(fetchedUser.Password), []byte(s.TestUser.Password)),
		"wrong password for fetched user",
	)
	s.Equal(
		fetchedUser.IsActive,
		newUser.IsActive,
		"wrong active status for fetched user",
	)
	s.Equal(
		fetchedUser.IsAdmin,
		newUser.IsAdmin,
		"wrong admin status for fetched user",
	)
}

func (s *DbUserTestSuite) TestGetUserByID() {
	hashedPassword, _ := utils.HashPassword(s.TestUser.Password)
	user := sql.User{
		Username: s.TestUser.Username,
		Email:    s.TestUser.Email,
		Password: hashedPassword,
		IsActive: s.TestUser.IsActive,
		IsAdmin:  s.TestUser.IsAdmin,
	}
	newUser, err := s.SqlDb.CreateUser(context.Background(), user)
	s.NoError(
		err,
		"CreateUser() should not return an error",
	)
	parsedUserID := strconv.Itoa(newUser.ID)
	fetchedUser, err := s.SqlDb.GetUserByID(context.Background(), parsedUserID)
	s.NoError(
		err,
		"GetUserByID() should not return an error",
	)
	s.NotZero(
		fetchedUser.ID,
		"fetched user should have non-zero ID",
	)
	s.Equal(
		fetchedUser.Username,
		newUser.Username,
		"wrong username for fetched user",
	)
	s.Equal(
		fetchedUser.Email,
		newUser.Email,
		"wrong email for fetched user",
	)
	s.NoError(
		bcrypt.CompareHashAndPassword([]byte(fetchedUser.Password), []byte(s.TestUser.Password)),
		"wrong password for fetched user",
	)
	s.Equal(
		fetchedUser.IsActive,
		newUser.IsActive,
		"wrong active status for fetched user",
	)
	s.Equal(
		fetchedUser.IsAdmin,
		newUser.IsAdmin,
		"wrong admin status for fetched user",
	)
}

func (s *DbUserTestSuite) TearDownTest() {
	s.SqlDb.Clear()
}

func (s *DbUserTestSuite) TearDownSuite() {
	testutils.DisconnectSqlDB(s.SqlDb)
}

func TestDbUserTestSuite(t *testing.T) {
	suite.Run(t, new(DbUserTestSuite))
}
