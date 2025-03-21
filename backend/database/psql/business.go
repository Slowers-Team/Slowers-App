package database

import "context"

type Business struct {
	ID           int
	CreatedAt    string
	LastModified string
	BusinessName string
	Type         string
	PhoneNumber  string
	Email        string
	PostAddress  string
	PostalCode   string
	City         string
	Notes        string
}

func (pDb SQLDatabase) CreateBusiness(ctx context.Context, newBusiness Business) (*Business, error) {

	query := `
	INSERT INTO Businesses (
							name,
							type,
							phone_number,
							email,
							post_address,
							postal_code,
							city,
							notes)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING id`

	err := pDb.pool.QueryRow(
		ctx,
		query,
		newBusiness.BusinessName,
		newBusiness.Type,
		newBusiness.PhoneNumber,
		newBusiness.Email,
		newBusiness.PostAddress,
		newBusiness.PostalCode,
		newBusiness.City,
		newBusiness.Notes,
	).Scan(&newBusiness.ID)

	if err != nil {
		return nil, err
	}

	return &newBusiness, nil
}
