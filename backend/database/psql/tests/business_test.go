//go:build sql
// +build sql

package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	database "github.com/Slowers-team/Slowers-App/database/psql"
	"github.com/Slowers-team/Slowers-App/testdataPsql"
	"github.com/Slowers-team/Slowers-App/testutils"
)

type DbBusinessTestSuite struct {
	suite.Suite
	Db           database.Database
	TestBusiness database.Business
}

func (s *DbBusinessTestSuite) SetupSuite() {
	s.Db = testutils.ConnectSQLDB()
	s.Db.Clear()
	s.TestBusiness = testdataPsql.GetBusinesses()[0]
}

func (s *DbBusinessTestSuite) TestCreateBusiness() {
	business := database.Business{
		BusinessName: s.TestBusiness.BusinessName,
		Type:         s.TestBusiness.Type,
		PhoneNumber:  s.TestBusiness.PhoneNumber,
		Email:        s.TestBusiness.Email,
		PostAddress:  s.TestBusiness.PostAddress,
		PostalCode:   s.TestBusiness.PostalCode,
		City:         s.TestBusiness.City,
		Notes:        s.TestBusiness.Notes,
	}
	newBusiness, err := s.Db.CreateBusiness(context.Background(), business)

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
		newBusiness.PostAddress,
		s.TestBusiness.PostAddress,
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
		newBusiness.Notes,
		s.TestBusiness.Notes,
		"wrong notes for new business",
	)
}

func (s *DbBusinessTestSuite) TearDownTest() {
	s.Db.Clear()
}

func (s *DbBusinessTestSuite) TearDownSuite() {
	testutils.DisconnectSQLDB(s.Db)
}

func TestDbBusinessTestSuite(t *testing.T) {
	suite.Run(t, new(DbBusinessTestSuite))
}
