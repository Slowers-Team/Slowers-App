package apitests

import (
	"context"
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
	TestUser database.User
}

func (s *UsersAPITestSuite) SetupSuite() {
	s.TestUser = testdata.GetUser()
}

func (s *UsersAPITestSuite) TestCreatingUser() {
	testutils.RunTest(s.T(), testutils.TestCase{
		Description: "POST /api/register",
		Route:       "/api/register",
		Method:      "POST",
		Body: utils.UserToJSON(
			database.User{
				Username: s.TestUser.Username,
				Email:    s.TestUser.Email,
				Password: s.TestUser.Password,
			},
		),
		ExpectedCode: 201,
		ExpectedBody: "Created",
		SetupMocks: func(db *mocks.Database) {
			db.EXPECT().CountUsersWithEmail(
				mock.Anything, s.TestUser.Email,
			).Return(
				0, nil,
			).Once()
			db.EXPECT().CreateUser(
				mock.Anything, mock.Anything,
			).RunAndReturn(func(ctx context.Context, user database.User) error {
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
				return nil
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
		Body: utils.LogInToJSON(
			database.LogIn{
				Email:    s.TestUser.Email,
				Password: s.TestUser.Password,
			},
		),
		ExpectedCode: 200,
		ExpectedBodyFunc: func(body string) bool {
			// TODO: Check that the token is valid
			return true
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

func TestUsersAPITestSuite(t *testing.T) {
	suite.Run(t, new(UsersAPITestSuite))
}
