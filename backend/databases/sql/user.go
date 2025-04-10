package sql

import (
	"context"
	"fmt"
	"strconv"

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
	query := `
	INSERT INTO users (username, password, email, is_active, is_admin)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id`

	err := pDb.pool.QueryRow(ctx, query,
		newUser.Username,
		newUser.Password,
		newUser.Email,
		newUser.IsActive,
		newUser.IsAdmin,
	).Scan(&newUser.ID)

	if err != nil {
		return nil, err
	}

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

func (pDb SQLDatabase) GetUserByID(ctx context.Context, userID string) (*User, error) {
	user := new(User)
	parsedUserID, err := strconv.Atoi(userID)
	if err != nil {
		return nil, err
	}

	query := `SELECT id, created_at::TEXT, last_modified::TEXT, last_login::TEXT, username, password, email, is_active, is_admin FROM users WHERE id=$1`
	err = pDb.pool.QueryRow(ctx, query, parsedUserID).Scan(
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
