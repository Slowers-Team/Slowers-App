package testdata

import "github.com/Slowers-team/Slowers-App/databases/sql"

func GetUsers() []sql.User {
	return []sql.User{
		{
			ID:       1,
			Username: "testuser",
			Password: "testpassword",
			Email:    "testuser@test.com",
			IsActive: false,
			IsAdmin:  false,
		},
		{
			ID:       2,
			Username: "testuser2",
			Email:    "testuser2@test.com",
			Password: "testpassword2",
			IsActive: false,
			IsAdmin:  false,
		},
		{
			ID:       3,
			Username: "testuser3",
			Email:    "testuser3@test.com",
			Password: "testpassword",
			IsActive: false,
			IsAdmin:  false,
		},
		{
			ID:       4,
			Username: "testuser4",
			Email:    "testuser4@test.com",
			Password: "testpassword",
			IsActive: false,
			IsAdmin:  false,
		},
		{
			ID:       5,
			Username: "testuser5",
			Email:    "testuser5@test.com",
			Password: "testpassword",
			IsActive: false,
			IsAdmin:  false,
		},
	}
}
