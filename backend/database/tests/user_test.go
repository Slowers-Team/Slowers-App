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
	s.TestUser = testdata.GetUsers()[0]
}

func (s *DbUserTestSuite) TestCreateUser() {
	hashedPassword, _ := utils.HashPassword(s.TestUser.Password)
	user := database.User{
		Username: s.TestUser.Username,
		Email:    s.TestUser.Email,
		Password: hashedPassword,
	}
	err := s.Db.CreateUser(context.Background(), user)

	s.NoError(
		err,
		"CreateUser() should not return an error",
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
	s.Db.CreateUser(context.Background(), user)

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

func (s *DbUserTestSuite) TestCreateAndGetUserByID() {
	hashedPassword, _ := utils.HashPassword(s.TestUser.Password)
	user := database.User{
		Username: s.TestUser.Username,
		Email:    s.TestUser.Email,
		Password: hashedPassword,
		Role:     s.TestUser.Role,
	}
	s.Db.CreateUser(context.Background(), user)

	createdUser, _ := s.Db.GetUserByEmail(context.Background(), s.TestUser.Email)

	fetchedUser, err := s.Db.GetUserByID(context.Background(), createdUser.ID)

	s.NoError(
		err,
		"GetUserByID() should not return an error",
	)
	s.NotZero(
		fetchedUser.ID,
		"fetched user should have non-zero ID",
	)
	s.Equal(
		s.TestUser.Username,
		fetchedUser.Username,
		"wrong username for fetched user",
	)
	s.Equal(
		s.TestUser.Email,
		fetchedUser.Email,
		"wrong email for fetched user",
	)
	s.Equal(
		"",
		fetchedUser.Password,
		"fetched user should have empty password",
	)
}

func (s *DbUserTestSuite) TestCreateUserAndChangeRole() {
	hashedPassword, _ := utils.HashPassword(s.TestUser.Password)
	user := database.User{
		Username: s.TestUser.Username,
		Email:    s.TestUser.Email,
		Password: hashedPassword,
		Role:     s.TestUser.Role,
	}
	s.Db.CreateUser(context.Background(), user)

	createdUser, _ := s.Db.GetUserByEmail(context.Background(), s.TestUser.Email)

	err := s.Db.SetUserRole(context.Background(), createdUser.ID, "retailer")

	s.NoError(
		err,
		"SetUserRole() should not return an error",
	)

	editedUser, _ := s.Db.GetUserByEmail(context.Background(), s.TestUser.Email)

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
