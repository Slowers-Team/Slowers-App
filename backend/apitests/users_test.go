//go:build api
// +build api

package apitests

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"

	"github.com/Slowers-team/Slowers-App/databases/sql"
	"github.com/Slowers-team/Slowers-App/mocks"
	"github.com/Slowers-team/Slowers-App/testdata"
	"github.com/Slowers-team/Slowers-App/testutils"
	"github.com/Slowers-team/Slowers-App/utils"
)

type UsersAPITestSuite struct {
	suite.Suite
	User  sql.User
	Users []sql.User
}

func (s *UsersAPITestSuite) SetupSuite() {
	s.User = testdata.GetUsers()[0]
}

func (s *UsersAPITestSuite) TestCreatingUser() {
	userToAdd := s.User
	userToAdd.ID = database.NilObjectID

	testutils.RunTest(s.T(), testutils.TestCase{
		Description:  "POST /api/register",
		Route:        "/api/register",
		Method:       "POST",
		ContentType:  "application/json",
		Body:         utils.ToJSON(userToAdd),
		ExpectedCode: 201,
		ExpectedBodyFunc: func(body []byte) {
			var response struct {
				Role  string `json:"role"`
				Token string `json:"token"`
			}
			err := json.Unmarshal([]byte(body), &response)
			s.Require().NoError(
				err,
				"response body should contain a role and a token",
			)
			s.Equal(
				s.User.Role,
				response.Role,
				"tried to add wrong role to database",
			)
			s.NotEmpty(
				response.Token,
				"token should not be empty",
			)
		},

		SetupMocks: func(db *mocks.Database) {
			db.EXPECT().CountUsersWithEmail(
				mock.Anything, s.User.Email,
			).Return(
				0, nil,
			).Once()
			db.EXPECT().CreateUser(
				mock.Anything, mock.Anything,
			).RunAndReturn(func(ctx context.Context, user database.User) (*database.User, error) {
				s.Equal(
					s.User.Username,
					user.Username,
					"tried to add wrong username to database",
				)
				s.Equal(
					s.User.Email,
					user.Email,
					"tried to add wrong email to database",
				)
				s.NoError(
					bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(s.User.Password)),
					"tried to add wrong Username to database",
				)
				s.Equal(
					s.User.Role,
					user.Role,
					"tried to add wrong role to database",
				)
				return &s.User, nil
			}).Once()
		},
	})
}

func (s *UsersAPITestSuite) TestLoggingIn() {
	hashedPassword, _ := utils.HashPassword(s.User.Password)

	testutils.RunTest(s.T(), testutils.TestCase{
		Description: "POST /api/login",
		Route:       "/api/login",
		Method:      "POST",
		ContentType: "application/json",
		Body: utils.ToJSON(
			database.LogIn{
				Email:    s.User.Email,
				Password: s.User.Password,
			},
		),
		ExpectedCode: 200,
		ExpectedBodyFunc: func(body []byte) {
			// TODO: Check here that the token is valid
		},
		SetupMocks: func(db *mocks.Database) {
			db.EXPECT().GetUserByEmail(
				mock.Anything, s.User.Email,
			).Return(
				&database.User{
					ID:       s.User.ID,
					Username: s.User.Username,
					Email:    s.User.Email,
					Password: hashedPassword,
				},
				nil,
			).Once()
		},
	})
}

func (s *UsersAPITestSuite) TestFetchingUser() {
	testutils.RunTest(s.T(), testutils.TestCase{
		Description:  "GET /api/user",
		Route:        "/api/user",
		Method:       "GET",
		ContentType:  "application/json",
		Body:         []byte{},
		ExpectedCode: 200,
		ExpectedBody: utils.ToJSON(s.User),
		SetupMocks: func(db *mocks.Database) {
			db.EXPECT().GetUserByID(
				mock.Anything, s.User.ID,
			).Return(
				&s.User, nil,
			).Once()
		},
	})
}

func (s *UsersAPITestSuite) TestChangingRole() {
	role := "retailer"
	roleJSON := "\"" + role + "\""

	testutils.RunTest(s.T(), testutils.TestCase{
		Description:  "POST /api/user/role",
		Route:        "/api/user/role",
		Method:       "POST",
		ContentType:  "application/json",
		Body:         []byte(roleJSON),
		ExpectedCode: 201,
		ExpectedBody: []byte(roleJSON),
		SetupMocks: func(db *mocks.Database) {
			db.EXPECT().SetUserRole(
				mock.Anything, s.User.ID, role,
			).Return(
				nil,
			).Once()
		},
	})
}

func (s *UsersAPITestSuite) TestCreatingUsersWithAllAvailableRoles() { //TODO: This is currently modified copy-paste of TestCreatingUser()
	for i := 0; i < len(s.Users); i++ {
		testutils.RunTest(s.T(), testutils.TestCase{
			Description: "POST /api/register",
			Route:       "/api/register",
			Method:      "POST",
			ContentType: "application/json",
			Body: utils.ToJSON(
				database.User{
					Username: s.Users[i].Username,
					Email:    s.Users[i].Email,
					Password: s.Users[i].Password,
					Role:     s.Users[i].Role,
				},
			),
			ExpectedCode: 201,
			ExpectedBodyFunc: func(body []byte) {
				var response struct {
					Role  string `json:"role"`
					Token string `json:"token"`
				}
				err := json.Unmarshal([]byte(body), &response)
				s.NoError(err, "response body should contain a role and a token")
				s.Equal(
					response.Role,
					s.Users[i].Role,
					"tried to add wrong role to database",
				)
				s.NotEmpty(
					response.Token,
					"token should not be empty",
				)
			},

			SetupMocks: func(db *mocks.Database) {
				db.EXPECT().CountUsersWithEmail(
					mock.Anything, s.Users[i].Email,
				).Return(
					0, nil,
				).Once()
				db.EXPECT().CreateUser(
					mock.Anything, mock.Anything,
				).RunAndReturn(func(ctx context.Context, user database.User) (*database.User, error) {
					s.Equal(
						user.Username,
						s.Users[i].Username,
						"tried to add wrong username to database",
					)
					s.Equal(
						user.Email,
						s.Users[i].Email,
						"tried to add wrong email to database",
					)
					s.NoError(
						bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(s.Users[i].Password)),
						"tried to add wrong Username to database",
					)
					s.Equal(
						user.Role,
						s.Users[i].Role,
						"tried to add wrong role to database",
					)
					return &s.Users[i], nil
				}).Once()
			},
		})
	}
}

func (s *UsersAPITestSuite) TestCreatingUserWithNonavailableRole() {
	testutils.RunTest(s.T(), testutils.TestCase{
		Description: "POST /api/register",
		Route:       "/api/register",
		Method:      "POST",
		ContentType: "application/json",
		Body: utils.ToJSON(
			database.User{
				Username: s.User.Username,
				Email:    s.User.Email,
				Password: s.User.Password,
				Role:     "superadmin",
			},
		),
		ExpectedCode: 400,
		ExpectedBody: []byte("unknown role: superadmin"),
		SetupMocks: func(db *mocks.Database) {
		},
	})
}

func (s *UsersAPITestSuite) TestChangingRoleWithAllAvailableRoles() {
	roles := [4]string{"retailer", "grower", "retailerowner", "growerowner"}

	for i := 0; i < len(roles); i++ {
		roleJSON := "\"" + roles[i] + "\""
		testutils.RunTest(s.T(), testutils.TestCase{
			Description:  "POST /api/user/role",
			Route:        "/api/user/role",
			Method:       "POST",
			ContentType:  "application/json",
			Body:         []byte(roleJSON),
			ExpectedCode: 201,
			ExpectedBody: []byte(roleJSON),
			SetupMocks: func(db *mocks.Database) {
				db.EXPECT().SetUserRole(
					mock.Anything, s.User.ID, roles[i],
				).Return(
					nil,
				).Once()
			},
		})
	}
}

func (s *UsersAPITestSuite) TestChangingRoleWithNonavailableRole() {
	role := "superadmin"
	roleJSON := "\"" + role + "\""

	testutils.RunTest(s.T(), testutils.TestCase{
		Description:  "POST /api/user/role",
		Route:        "/api/user/role",
		Method:       "POST",
		ContentType:  "application/json",
		Body:         []byte(roleJSON),
		ExpectedCode: 400,
		ExpectedBody: []byte("unknown role: superadmin"),
		SetupMocks: func(db *mocks.Database) {
		},
	})
}

func TestUsersAPITestSuite(t *testing.T) {
	suite.Run(t, new(UsersAPITestSuite))
}
