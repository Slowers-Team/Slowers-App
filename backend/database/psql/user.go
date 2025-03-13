package database

import (
	"context"

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
	).Scan(&newUser.ID, &newUser.CreatedAt, &newUser.LastModified)

	if err != nil {
		return nil, err
	}

	return &newUser, nil
}
