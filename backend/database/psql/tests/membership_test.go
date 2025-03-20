//go:build sql
// +build sql

package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"

	database "github.com/Slowers-team/Slowers-App/database/psql"
	"github.com/Slowers-team/Slowers-App/testdataPsql"
	"github.com/Slowers-team/Slowers-App/testutils"
	"github.com/Slowers-team/Slowers-App/utils"
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
}

func (s *DbMembershipTestSuite) AddTestBusinessToDatabase() *database.Business { // TODO: this implementation is dependent on working business db logic
	bus := testdataPsql.GetBusinesses()[0]
	business := database.Business{
		BusinessName: bus.BusinessName,
		Type:         bus.Type,
		PhoneNumber:  bus.PhoneNumber,
		Email:        bus.Email,
		PostAddress:  bus.PostAddress,
		PostalCode:   bus.PostalCode,
		City:         bus.City,
		Notes:        bus.Notes,
	}
	newBusiness, err := s.Db.CreateBusiness(context.Background(), business)

	if err != nil {
		fmt.Println("Error in test business creation: ", err) // TODO: Better error handling
		return nil
	}
	return newBusiness
}

func (s *DbMembershipTestSuite) SetupTest() {
	s.TestBusiness = *s.AddTestBusinessToDatabase()
	s.AddTestUserToDatabase()
}

func (s *DbMembershipTestSuite) AddTestUserToDatabase() { // TODO: this implementation is dependent on working business db logic
	hashedPassword, _ := utils.HashPassword(s.TestUser.Password)
	user := database.User{
		Username: s.TestUser.Username,
		Email:    s.TestUser.Email,
		Password: hashedPassword,
		IsActive: s.TestUser.IsActive,
		IsAdmin:  s.TestUser.IsAdmin,
	}
	_, err := s.Db.CreateUser(context.Background(), user)
	if err != nil {
		fmt.Println("Error in test user creation") // TODO: Better error handling
	}
}

func (s *DbMembershipTestSuite) TestAddMembership() {
	membership := database.Membership{
		UserEmail:   s.TestUser.Email,
		BusinessID:  s.TestBusiness.ID,
		Designation: "owner",
	}
	newMembership, err := s.Db.AddMembership(context.Background(), membership)

	s.NoError(
		err,
		"AddMembership() should not return an error",
	)
	s.NotZero(
		newMembership.ID,
		"new membership should have non-zero ID",
	)
}

func (s *DbMembershipTestSuite) TearDownTest() {
	s.Db.Clear()
}

func (s *DbMembershipTestSuite) TearDownSuite() {
	testutils.DisconnectSQLDB(s.Db)
}

func TestDbMembershipTestSuite(t *testing.T) {
	suite.Run(t, new(DbMembershipTestSuite))
}
