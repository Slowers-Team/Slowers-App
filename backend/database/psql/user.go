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
		fmt.Println("Errorin sisällä:", err)
		// täällä näkyy, että hashed on liian pitkä, kun skeemassa on rajana 50 merkkiä
		// ilman hashia toimii
		return nil, err
	}

	fmt.Println("Tämä printtautuu, jos homma toimii")

	return &newUser, nil
}
