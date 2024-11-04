package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"

	"github.com/Slowers-team/Slowers-App/database"
	"github.com/Slowers-team/Slowers-App/testdata"
	"github.com/Slowers-team/Slowers-App/testutils"
	"github.com/Slowers-team/Slowers-App/utils"
)

type DbUserTestSuite struct {
	suite.Suite
	Db       database.Database
	TestUser database.User
}

func (s *DbUserTestSuite) SetupSuite() {
	s.Db = testutils.ConnectDB()
	s.Db.Clear()
	s.TestUser = testdata.GetUser()
}

func (s *DbUserTestSuite) TestCreateUser() {
	hashedPassword, _ := utils.HashPassword(s.TestUser.Password)
	user := database.User{
		Username: s.TestUser.Username,
		Email:    s.TestUser.Email,
		Password: hashedPassword,
	}
	newUser, err := s.Db.CreateUser(context.Background(), user)

	s.NoError(
		err,
		"CreateUser() should not return an error",
	)
	s.NotNil(
		newUser,
		"CreateUser() should return a non-nil user",
	)
}

func (s *DbUserTestSuite) TestCountUsers() {
	count, err := s.Db.CountUsersWithEmail(context.Background(), s.TestUser.Email)

	s.NoError(
		err,
		"CountUsersWithEmail() should not return an error",
	)
	s.Equal(
		count,
		int64(0),
		"CountUsersWithEmail() should return 0 for an empty database",
	)
}

func (s *DbUserTestSuite) TestCreateAndGetUser() {
	hashedPassword, _ := utils.HashPassword(s.TestUser.Password)
	user := database.User{
		Username: s.TestUser.Username,
		Email:    s.TestUser.Email,
		Password: hashedPassword,
	}

	newUser, err := s.Db.CreateUser(context.Background(), user)

	s.NoError(
		err,
		"CreateUser() should not return an error",
	)
	s.NotNil(
		newUser,
		"CreateUser() should return a non-nil user",
	)

	fetchedUser, err := s.Db.GetUserByEmail(context.Background(), s.TestUser.Email)

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
		s.TestUser.Username,
		"wrong username for fetched user",
	)
	s.Equal(
		fetchedUser.Email,
		s.TestUser.Email,
		"wrong email for fetched user",
	)
	s.NoError(
		bcrypt.CompareHashAndPassword([]byte(fetchedUser.Password), []byte(s.TestUser.Password)),
		"wrong password for fetched user",
	)
}

func (s *DbUserTestSuite) TearDownTest() {
	s.Db.Clear()
}

func (s *DbUserTestSuite) TearDownSuite() {
	testutils.DisconnectDB(s.Db)
}

func TestDbUserTestSuite(t *testing.T) {
	suite.Run(t, new(DbUserTestSuite))
}
