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
	BusinessName string
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

func (pDb SQLDatabase) GetMembershipByUserId(ctx context.Context, userID int) (*Membership, error) {
	membership := new(Membership)
	query := `
	SELECT
		memberships.id,
		memberships.created_at::TEXT,
		memberships.last_modified::TEXT,
		memberships.user_email,
		memberships.business_id,
		memberships.designation
	FROM
		memberships
	JOIN
		users
	ON
		memberships.user_email = users.email
	WHERE
		users.id=$1`

	err := pDb.pool.QueryRow(ctx, query, userID).Scan(
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

func (pDd SQLDatabase) DeleteMemberByUserID(ctx context.Context,user_email string, business_id int) error {
	membership := new(Membership)
	query := `
	DELETE FROM Memberships 
	WHERE user_email = $1
	AND business_id = $2
	`

	err := pDd.pool.QueryRow(ctx, query, user_email, business_id).Scan(
		ctx,
		query,
		&membership.UserEmail,
		&membership.BusinessID,
	)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil

}

// func (pDb SQLDatabase) GetDesignationByEmail(ctx context.Context, userEmail string) (*Membership, error) {
// 	membership := new(Membership)
// 	query := `
// 			SELECT
// 				Memberships.designation
// 			FROM
// 				Memberships
// 				INNER JOIN Users ON Memberships.user_email = Users.email
// 			WHERE
// 				User.email = 1$`

// 	err := pDb.pool.QueryRow(ctx, query, userEmail).Scan(
// 		&membership.Designation,
// 	)

// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return nil, err
// 	}

// 	return membership, nil
// }
