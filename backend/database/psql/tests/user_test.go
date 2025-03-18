package tests

import (
	"context"
	//"testing"

	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"

	"github.com/Slowers-team/Slowers-App/database"
	"github.com/Slowers-team/Slowers-App/testdata"
	"github.com/Slowers-team/Slowers-App/testutils"
	"github.com/Slowers-team/Slowers-App/utils"
)

// NÄMÄ ASIAT OTETTU SUORAAN MONGOA KÄYTTÄVÄSTÄ USER_TESTISTÄ
// TÄYTYY MUOKATA FUNKTIOITA

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
}
