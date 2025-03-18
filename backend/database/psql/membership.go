package database

import "context"

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
