package tests

import (
	"context"
	"strconv"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/Slowers-team/Slowers-App/databases/sql"
	"github.com/Slowers-team/Slowers-App/testdata"
	"github.com/Slowers-team/Slowers-App/testutils"
)

type DbBusinessTestSuite struct {
	suite.Suite
	SqlDb          sql.Database
	TestBusiness   sql.Business
	TestUser       sql.User
	TestMembership sql.Membership
}

func (s *DbBusinessTestSuite) SetupSuite() {
	s.SqlDb = testutils.ConnectSqlDB()
	s.SqlDb.Clear()
	s.TestBusiness = testdata.GetBusinesses()[0]
}

func (s *DbBusinessTestSuite) TestCreateBusiness() {
	business := sql.Business{
		BusinessName:   s.TestBusiness.BusinessName,
		BusinessIdCode: s.TestBusiness.BusinessIdCode,
		Type:           s.TestBusiness.Type,
		PhoneNumber:    s.TestBusiness.PhoneNumber,
		Email:          s.TestBusiness.Email,
		Address:        s.TestBusiness.Address,
		PostalCode:     s.TestBusiness.PostalCode,
		City:           s.TestBusiness.City,
		AdditionalInfo: s.TestBusiness.AdditionalInfo,
		Delivery:       s.TestBusiness.Delivery,
	}
	newBusiness, err := s.SqlDb.CreateBusiness(context.Background(), business)

	s.NoError(
		err,
		"CreateBusiness() should not return an error",
	)
	s.NotZero(
		newBusiness.ID,
		"new user should have non-zero ID",
	)
	s.Equal(
		newBusiness.BusinessName,
		s.TestBusiness.BusinessName,
		"wrong name for new business",
	)
	s.Equal(
		newBusiness.BusinessIdCode,
		s.TestBusiness.BusinessIdCode,
		"wrong business id code for new business",
	)
	s.Equal(
		newBusiness.Type,
		s.TestBusiness.Type,
		"wrong business type for new business",
	)
	s.Equal(
		newBusiness.PhoneNumber,
		s.TestBusiness.PhoneNumber,
		"wrong phone number for new business",
	)
	s.Equal(
		newBusiness.Email,
		s.TestBusiness.Email,
		"wrong email for new business",
	)
	s.Equal(
		newBusiness.Address,
		s.TestBusiness.Address,
		"wrong post address for new business",
	)
	s.Equal(
		newBusiness.PostalCode,
		s.TestBusiness.PostalCode,
		"wrong postal code for new business",
	)
	s.Equal(
		newBusiness.City,
		s.TestBusiness.City,
		"wrong city for new business",
	)
	s.Equal(
		newBusiness.AdditionalInfo,
		s.TestBusiness.AdditionalInfo,
		"wrong additional info for new business",
	)
	s.Equal(
		newBusiness.Delivery,
		s.TestBusiness.Delivery,
		"wrong delivery option for new business",
	)
}

func (s *DbBusinessTestSuite) TestGetBusinessByUserID() {
	createdBusiness := sql.Business{
		BusinessName:   s.TestBusiness.BusinessName,
		BusinessIdCode: s.TestBusiness.BusinessIdCode,
		Type:           s.TestBusiness.Type,
		PhoneNumber:    s.TestBusiness.PhoneNumber,
		Email:          s.TestBusiness.Email,
		Address:        s.TestBusiness.Address,
		PostalCode:     s.TestBusiness.PostalCode,
		City:           s.TestBusiness.City,
		AdditionalInfo: s.TestBusiness.AdditionalInfo,
		Delivery:       s.TestBusiness.Delivery,
	}
	_, err := s.SqlDb.CreateBusiness(context.Background(), createdBusiness)
	s.NoError(
		err,
		"CreateBusiness() should not return an error",
	)

	membership := sql.Membership{
		UserEmail:   s.TestUser.Email,
		BusinessID:  s.TestBusiness.ID,
		Designation: "owner",
	}
	_, err = s.SqlDb.AddMembership(context.Background(), membership)
	s.NoError(
		err,
		"AddMembership() should not return an error",
	)
	parsedUserID := strconv.Itoa(s.TestUser.ID)
	business, err := s.SqlDb.GetBusinessByUserID(context.Background(), parsedUserID)
	s.NoError(
		err,
		"GetBusinessByUserID() should not return an error",
	)
	s.Equal(
		business.BusinessName,
		createdBusiness.BusinessName,
		"wrong name for fetched business",
	)
	s.Equal(
		business.BusinessIdCode,
		createdBusiness.BusinessIdCode,
		"wrong business id code for fetched business",
	)
	s.Equal(
		business.Type,
		createdBusiness.Type,
		"wrong business type for fetched business",
	)
	s.Equal(
		business.PhoneNumber,
		createdBusiness.PhoneNumber,
		"wrong phone number for fetched business",
	)
	s.Equal(
		business.Email,
		createdBusiness.Email,
		"wrong email for fetched business",
	)
	s.Equal(
		business.Address,
		createdBusiness.Address,
		"wrong post address for fetched business",
	)
	s.Equal(
		business.PostalCode,
		createdBusiness.PostalCode,
		"wrong postal code for fetched business",
	)
	s.Equal(
		business.City,
		createdBusiness.City,
		"wrong city for fetched business",
	)
	s.Equal(
		business.AdditionalInfo,
		createdBusiness.AdditionalInfo,
		"wrong additional info for fetched business",
	)
	s.Equal(
		business.Delivery,
		createdBusiness.Delivery,
		"wrong delivery option for fetched business",
	)
}

func (s *DbBusinessTestSuite) TearDownTest() {
	s.SqlDb.Clear()
}

func (s *DbBusinessTestSuite) TearDownSuite() {
	testutils.DisconnectSqlDB(s.SqlDb)
}

func TestDbBusinessTestSuite(t *testing.T) {
	suite.Run(t, new(DbBusinessTestSuite))
}
