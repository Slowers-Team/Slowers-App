package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"

	"github.com/Slowers-team/Slowers-App/database"
	"github.com/Slowers-team/Slowers-App/testdata"
	"github.com/Slowers-team/Slowers-App/testutils"
)

type DbUserTestSuite struct {
	suite.Suite
	Db   database.Database
	User database.User
}

func (s *DbUserTestSuite) SetupSuite() {
	s.Db = testutils.ConnectDB()
	s.Db.Clear()
	s.User = testdata.GetUsers()[0]
}

func (s *DbUserTestSuite) TestCreateUser() {
	user := s.User
	userToAdd := testdata.PrepareUserForAdding(user)
	addedUser, err := s.Db.CreateUser(context.Background(), userToAdd)

	s.Require().NoError(
		err,
		"CreateUser() should not return an error",
	)
	s.NotZero(
		addedUser.ID,
		"new user should have non-zero ID",
	)
	s.Equal(
		user.Username,
		addedUser.Username,
		"wrong username for new user",
	)
	s.Equal(
		user.Email,
		addedUser.Email,
		"wrong email for new user",
	)
	s.NoError(
		bcrypt.CompareHashAndPassword(
			[]byte(addedUser.Password), []byte(user.Password),
		),
		"wrong password for new user",
	)
}

func (s *DbUserTestSuite) TestCountUsers() {
	count, err := s.Db.CountUsersWithEmail(context.Background(), s.User.Email)

	s.NoError(
		err,
		"CountUsersWithEmail() should not return an error",
	)
	s.Equal(
		int64(0),
		count,
		"CountUsersWithEmail() should return 0 for an empty database",
	)
}

func (s *DbUserTestSuite) TestCreateAndGetUser() {
	user := s.User
	userToAdd := testdata.PrepareUserForAdding(user)
	addedUser, _ := s.Db.CreateUser(context.Background(), userToAdd)

	fetchedUser, err := s.Db.GetUserByEmail(context.Background(), user.Email)

	s.Require().NoError(
		err,
		"GetUserByEmail() should not return an error",
	)
	s.Equal(
		addedUser.ID,
		fetchedUser.ID,
		"wrong ID for fetched user",
	)
	s.Equal(
		user.Username,
		fetchedUser.Username,
		"wrong username for fetched user",
	)
	s.Equal(
		user.Email,
		fetchedUser.Email,
		"wrong email for fetched user",
	)
	s.NoError(
		bcrypt.CompareHashAndPassword(
			[]byte(fetchedUser.Password), []byte(user.Password),
		),
		"wrong password for fetched user",
	)
}

func (s *DbUserTestSuite) TestCreateAndGetUserByID() {
	user := s.User
	userToAdd := testdata.PrepareUserForAdding(user)
	s.Db.CreateUser(context.Background(), userToAdd)

	addedUser, _ := s.Db.GetUserByEmail(context.Background(), user.Email)

	fetchedUser, err := s.Db.GetUserByID(context.Background(), addedUser.ID)

	s.Require().NoError(
		err,
		"GetUserByID() should not return an error",
	)
	s.Equal(
		addedUser.ID,
		fetchedUser.ID,
		"wrong ID for fetched user",
	)
	s.Equal(
		user.Username,
		fetchedUser.Username,
		"wrong username for fetched user",
	)
	s.Equal(
		user.Email,
		fetchedUser.Email,
		"wrong email for fetched user",
	)
	s.Equal(
		"",
		fetchedUser.Password,
		"fetched user should have empty password",
	)
	s.Equal(
		user.Role,
		fetchedUser.Role,
		"wrong role for fetched user",
	)
}

func (s *DbUserTestSuite) TestCreateUserAndChangeRole() {
	user := s.User
	userToAdd := testdata.PrepareUserForAdding(user)
	s.Db.CreateUser(context.Background(), userToAdd)

	fetchedUser, _ := s.Db.GetUserByEmail(context.Background(), user.Email)

	err := s.Db.SetUserRole(context.Background(), fetchedUser.ID, "retailer")

	s.Require().NoError(
		err,
		"SetUserRole() should not return an error",
	)

	editedUser, _ := s.Db.GetUserByEmail(context.Background(), user.Email)

	s.Equal(
		"retailer",
		editedUser.Role,
		"role should have changed to \"retailer\" for the user",
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
