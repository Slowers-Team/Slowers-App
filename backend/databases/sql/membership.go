package sql

import (
	"context"
	"fmt"
	"strconv"
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

func (pDb SQLDatabase) GetMembershipByUserId(ctx context.Context, userID string) (*Membership, error) {
	membership := new(Membership)
	parsedUserID, err := strconv.Atoi(userID)
	if err != nil {
		return nil, err
	}

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

	err = pDb.pool.QueryRow(ctx, query, parsedUserID).Scan(
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

func (pDd SQLDatabase) DeleteMembership(ctx context.Context, userEmail string, businessId int) error {
	query := `
	DELETE FROM Memberships 
	WHERE user_email = $1
	AND business_id = $2
	`

	_, err := pDd.pool.Exec(ctx, query, userEmail, businessId)

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

func (pDb SQLDatabase) GetAllMembersInBusiness(ctx context.Context, businessID int) ([]Membership, error) {
	query := `
	SELECT
		user_email,
		designation
	FROM
		Memberships
	WHERE
		business_id=$1
	ORDER BY
		user_email;`

	rows, err := pDb.pool.Query(ctx, query, businessID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var memberships []Membership

	for rows.Next() {
		var member Membership
		err := rows.Scan(
			&member.UserEmail,
			&member.Designation,
		)
		if err != nil {
			return nil, err
		}
		memberships = append(memberships, member)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return memberships, nil
}

func (pDb SQLDatabase) EditMembership(ctx context.Context, membership Membership) error {
	query := `
	UPDATE 
		Memberships
	SET 
		designation=$3
	WHERE 
		user_email=$1 AND
		business_id=$2
	`
	_, err := pDb.pool.Exec(
		ctx,
		query,
		membership.UserEmail,
		membership.BusinessID,
		membership.Designation,
	)

	if err != nil {
		return err
	}

	return nil
}
