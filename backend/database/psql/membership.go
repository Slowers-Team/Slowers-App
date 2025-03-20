package database

import (
	"context"
	"fmt"
)

type Membership struct {
	ID           int
	CreatedAt    string
	LastModified string
	UserEmail    string
	BusinessID   int
	Designation  string
}

func (pDb SQLDatabase) AddMembership(ctx context.Context, newMembership Membership) (*Membership, error) {
	query := `
	INSERT INTO Memberships (
							user_email,
							business_id,
							designation)
	VALUES ($1, $2, $3)
	RETURNING id`

	err := pDb.pool.QueryRow(
		ctx,
		query,
		newMembership.UserEmail,
		newMembership.BusinessID,
		newMembership.Designation,
	).Scan(&newMembership.ID)

	if err != nil {
		return nil, err
	}

	return &newMembership, nil
}

func (pDb SQLDatabase) CheckMembership(ctx context.Context, userEmail string) (*Membership, error) {
	membership := new(Membership)
	query := `SELECT id, created_at::TEXT, last_modified::TEXT, user_email, business_id, designation FROM memberships WHERE user_email=$1`
	err := pDb.pool.QueryRow(ctx, query, userEmail).Scan(
		&membership.ID,
		&membership.CreatedAt,
		&membership.LastModified,
		&membership.UserEmail,
		&membership.BusinessID,
		&membership.Designation,
	)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return membership, nil
}
