package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/Slowers-team/Slowers-App/databases/sql"
	"github.com/Slowers-team/Slowers-App/testdata"
	"github.com/Slowers-team/Slowers-App/testutils"
	"github.com/Slowers-team/Slowers-App/utils"
)

type DbMembershipTestSuite struct {
	suite.Suite
	SqlDb          sql.Database
	TestUser       sql.User
	TestBusiness   sql.Business
	TestMembership sql.Membership
}

func (s *DbMembershipTestSuite) SetupSuite() {
	s.SqlDb = testutils.ConnectSqlDB()
	s.SqlDb.Clear()
	s.TestUser = testdata.GetUsers()[0]
}

func (s *DbMembershipTestSuite) AddTestBusinessToDatabase() *sql.Business {
	bus := testdata.GetBusinesses()[0]
	business := sql.Business{
		BusinessName:   bus.BusinessName,
		BusinessIdCode: bus.BusinessIdCode,
		Type:           bus.Type,
		PhoneNumber:    bus.PhoneNumber,
		Email:          bus.Email,
		Address:        bus.Address,
		PostalCode:     bus.PostalCode,
		City:           bus.City,
		AdditionalInfo: bus.AdditionalInfo,
		Delivery:       bus.Delivery,
	}
	newBusiness, err := s.SqlDb.CreateBusiness(context.Background(), business)

	if err != nil {
		fmt.Println("Error in test business creation: ", err) // TODO: Better error handling
		return nil
	}
	return newBusiness
}

func (s *DbMembershipTestSuite) AddTestUserToDatabase() {
	hashedPassword, _ := utils.HashPassword(s.TestUser.Password)
	user := sql.User{
		Username: s.TestUser.Username,
		Email:    s.TestUser.Email,
		Password: hashedPassword,
		IsActive: s.TestUser.IsActive,
		IsAdmin:  s.TestUser.IsAdmin,
	}
	_, err := s.SqlDb.CreateUser(context.Background(), user)
	if err != nil {
		fmt.Println("Error in test user creation") // TODO: Better error handling
	}
}

func (s *DbMembershipTestSuite) SetupTest() {
	s.TestBusiness = *s.AddTestBusinessToDatabase()
	s.AddTestUserToDatabase()
}

func (s *DbMembershipTestSuite) TestAddMembership() {
	membership := sql.Membership{
		UserEmail:   s.TestUser.Email,
		BusinessID:  s.TestBusiness.ID,
		Designation: "owner",
	}
	newMembership, err := s.SqlDb.AddMembership(context.Background(), membership)

	s.NoError(
		err,
		"AddMembership() should not return an error",
	)
	s.NotZero(
		newMembership.ID,
		"new membership should have non-zero ID",
	)
	s.Equal(
		newMembership.UserEmail,
		s.TestUser.Email,
		"wrong user email for new membership",
	)
	s.Equal(
		newMembership.BusinessID,
		s.TestBusiness.ID,
		"wrong business id for new membership",
	)
	s.Equal(
		newMembership.Designation,
		"owner",
		"wrong membership designation for new membership",
	)
}

func (s *DbMembershipTestSuite) TestGetMembershipByUserId() {
	existingMembership := sql.Membership{
		UserEmail:   s.TestUser.Email,
		BusinessID:  s.TestBusiness.ID,
		Designation: "owner",
	}
	_, err := s.SqlDb.AddMembership(context.Background(), existingMembership)

	membership, err := s.SqlDb.GetMembershipByUserId(context.Background(), s.TestUser.ID)

	s.NoError(
		err,
		"CheckMembership() should not return an error",
	)
	s.NotZero(
		membership.ID,
		"membership should have non-zero ID",
	)
	s.Equal(
		membership.UserEmail,
		s.TestUser.Email,
		"wrong user email for membership",
	)
	s.Equal(
		membership.BusinessID,
		s.TestBusiness.ID,
		"wrong business id for membership",
	)
	s.Equal(
		membership.Designation,
		"owner",
		"wrong membership designation for membership",
	)
	s.Equal(
		membership.BusinessName,
		s.TestBusiness.BusinessName,
		"wrong business name for membership",
	)
}

// func (s *DbMembershipTestSuite) TestGetMembershipByUserEmailWorksWhenUserEmailHasNoUser() {
// 	existingMembership := sql.Membership{
// 		UserEmail:   "nonexistent@email.com",
// 		BusinessID:  s.TestBusiness.ID,
// 		Designation: "owner",
// 	}
// 	_, err := s.SqlDb.AddMembership(context.Background(), existingMembership)

// 	membership, err := s.SqlDb.GetMembershipByUserEmail(context.Background(), "nonexistent@email.com")

// 	s.NoError(
// 		err,
// 		"CheckMembership() should not return an error",
// 	)
// 	s.NotZero(
// 		membership.ID,
// 		"membership should have non-zero ID",
// 	)
// 	s.Equal(
// 		membership.UserEmail,
// 		"nonexistent@email.com",
// 		"wrong user email for membership",
// 	)
// 	s.Equal(
// 		membership.BusinessID,
// 		s.TestBusiness.ID,
// 		"wrong business id for membership",
// 	)
// 	s.Equal(
// 		membership.Designation,
// 		"owner",
// 		"wrong membership designation for membership",
// 	)
// 	s.Equal(
// 		membership.BusinessName,
// 		s.TestBusiness.BusinessName,
// 		"wrong business name for membership",
// 	)
// }

func (s *DbMembershipTestSuite) TearDownTest() {
	s.SqlDb.Clear()
}

func (s *DbMembershipTestSuite) TearDownSuite() {
	testutils.DisconnectSqlDB(s.SqlDb)
}

func TestDbMembershipTestSuite(t *testing.T) {
	suite.Run(t, new(DbMembershipTestSuite))
}
