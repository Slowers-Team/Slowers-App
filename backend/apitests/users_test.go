package apitests

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"

	"github.com/Slowers-team/Slowers-App/database"
	"github.com/Slowers-team/Slowers-App/mocks"
	"github.com/Slowers-team/Slowers-App/testdata"
	"github.com/Slowers-team/Slowers-App/testutils"
	"github.com/Slowers-team/Slowers-App/utils"
)

type UsersAPITestSuite struct {
	suite.Suite
	TestUser  database.User
	TestUsers []database.User
}

func (s *UsersAPITestSuite) SetupSuite() {
	s.TestUser = testdata.GetUsers()[0]
	s.TestUsers = testdata.GetUsers()
}

func (s *UsersAPITestSuite) TestCreatingUser() {
	testutils.RunTest(s.T(), testutils.TestCase{
		Description: "POST /api/register",
		Route:       "/api/register",
		Method:      "POST",
		ContentType: "application/json",
		Body: utils.ToJSON(
			database.User{
				Username: s.TestUser.Username,
				Email:    s.TestUser.Email,
				Password: s.TestUser.Password,
				Role:     s.TestUser.Role,
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
				s.TestUser.Role,
				"tried to add wrong role to database",
			)
			s.NotEmpty(
				response.Token,
				"token should not be empty",
			)
		},

		SetupMocks: func(db *mocks.Database) {
			db.EXPECT().CountUsersWithEmail(
				mock.Anything, s.TestUser.Email,
			).Return(
				0, nil,
			).Once()
			db.EXPECT().CreateUser(
				mock.Anything, mock.Anything,
			).RunAndReturn(func(ctx context.Context, user database.User) (*database.User, error) {
				s.Equal(
					user.Username,
					s.TestUser.Username,
					"tried to add wrong username to database",
				)
				s.Equal(
					user.Email,
					s.TestUser.Email,
					"tried to add wrong email to database",
				)
				s.NoError(
					bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(s.TestUser.Password)),
					"tried to add wrong Username to database",
				)
				s.Equal(
					user.Role,
					s.TestUser.Role,
					"tried to add wrong role to database",
				)
				return &s.TestUser, nil
			}).Once()
		},
	})
}

func (s *UsersAPITestSuite) TestLoggingIn() {
	hashedPassword, _ := utils.HashPassword(s.TestUser.Password)

	testutils.RunTest(s.T(), testutils.TestCase{
		Description: "POST /api/login",
		Route:       "/api/login",
		Method:      "POST",
		ContentType: "application/json",
		Body: utils.ToJSON(
			database.LogIn{
				Email:    s.TestUser.Email,
				Password: s.TestUser.Password,
			},
		),
		ExpectedCode: 200,
		ExpectedBodyFunc: func(body []byte) {
			// TODO: Check here that the token is valid
		},
		SetupMocks: func(db *mocks.Database) {
			db.EXPECT().GetUserByEmail(
				mock.Anything, s.TestUser.Email,
			).Return(
				&database.User{
					ID:       s.TestUser.ID,
					Username: s.TestUser.Username,
					Email:    s.TestUser.Email,
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
		ExpectedBody: utils.ToJSON(s.TestUser),
		SetupMocks: func(db *mocks.Database) {
			db.EXPECT().GetUserByID(
				mock.Anything, s.TestUser.ID,
			).Return(
				&s.TestUser, nil,
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
				mock.Anything, s.TestUser.ID, role,
			).Return(
				nil,
			).Once()
		},
	})
}

func (s *UsersAPITestSuite) TestCreatingUsersWithAllAvailableRoles() { //TODO: This is currently modified copy-paste of TestCreatingUser()
	for i := 0; i < len(s.TestUsers); i++ {
		testutils.RunTest(s.T(), testutils.TestCase{
			Description: "POST /api/register",
			Route:       "/api/register",
			Method:      "POST",
			ContentType: "application/json",
			Body: utils.ToJSON(
				database.User{
					Username: s.TestUsers[i].Username,
					Email:    s.TestUsers[i].Email,
					Password: s.TestUsers[i].Password,
					Role:     s.TestUsers[i].Role,
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
					s.TestUsers[i].Role,
					"tried to add wrong role to database",
				)
				s.NotEmpty(
					response.Token,
					"token should not be empty",
				)
			},

			SetupMocks: func(db *mocks.Database) {
				db.EXPECT().CountUsersWithEmail(
					mock.Anything, s.TestUsers[i].Email,
				).Return(
					0, nil,
				).Once()
				db.EXPECT().CreateUser(
					mock.Anything, mock.Anything,
				).RunAndReturn(func(ctx context.Context, user database.User) (*database.User, error) {
					s.Equal(
						user.Username,
						s.TestUsers[i].Username,
						"tried to add wrong username to database",
					)
					s.Equal(
						user.Email,
						s.TestUsers[i].Email,
						"tried to add wrong email to database",
					)
					s.NoError(
						bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(s.TestUsers[i].Password)),
						"tried to add wrong Username to database",
					)
					s.Equal(
						user.Role,
						s.TestUsers[i].Role,
						"tried to add wrong role to database",
					)
					return &s.TestUsers[i], nil
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
				Username: s.TestUser.Username,
				Email:    s.TestUser.Email,
				Password: s.TestUser.Password,
				Role:     "superadmin",
			},
		),
		ExpectedCode: 400,
		ExpectedBody: []byte("unknown role: superadmin"),
		SetupMocks: func(db *mocks.Database) {
		},
	})
}

func TestUsersAPITestSuite(t *testing.T) {
	suite.Run(t, new(UsersAPITestSuite))
}
