package database

import (
	"context"
	"fmt"

	_ "github.com/lib/pq"
)

type User struct {
	ID           int
	CreatedAt    string
	LastModified string
	LastLogin    string
	Username     string
	Password     string
	Email        string
	IsActive     bool
	IsAdmin      bool
}

type LogIn struct {
	Email    string
	Password string
}

func (pDb SQLDatabase) CreateUser(ctx context.Context, newUser User) (*User, error) {
	fmt.Println("This works!")
	query := `
	INSERT INTO users (last_login, username, password, email, is_active, is_admin)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id`

	err := pDb.pool.QueryRow(ctx, query,
		newUser.LastLogin,
		newUser.Username,
		newUser.Password,
		newUser.Email,
		newUser.IsActive,
		newUser.IsAdmin,
	).Scan(&newUser.ID)

	if err != nil {
		return nil, err
	}

	fmt.Println("Tämä printtautuu, jos homma toimii")

	return &newUser, nil
}

func (pDb SQLDatabase) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	user := new(User)
	query := `SELECT id, created_at::TEXT, last_modified::TEXT, last_login::TEXT, username, password, email, is_active, is_admin FROM users WHERE email=$1`
	err := pDb.pool.QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.CreatedAt, &user.LastModified,
		&user.LastLogin,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.IsActive,
		&user.IsAdmin,
	)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return user, nil
}

// func (pDb SQLDatabase) SetUserRole(ctx context.Context, userID ObjectID, role string) error {
// 	update := bson.M{"$set": bson.M{"role": role}}
// 	_, err := db.Collection("users").UpdateByID(ctx, userID, update)

// 	return err
// }

// func (pDb SQLDatabase) GetUserByID(ctx context.Context, userID ObjectID) (*User, error) {
// 	user := new(User)
// 	filter := bson.M{"_id": userID}
// 	err := db.Collection("users").FindOne(ctx, filter).Decode(&user)
// 	if err != nil {
// 		return nil, err
// 	}
// 	user.Password = ""
// 	return user, nil
// }
