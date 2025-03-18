package database

import "context"

type Business struct {
	ID                  int
	CreatedAt           string
	LastModified        string
	BusinessName        string
	BusinessType        string
	BusinessPhoneNumber string
	BusinessEmail       string
	BusinessAddress     string
	PostalCode          int
	City                string
	Notes               string
}

func (pDb SQLDatabase) CreateBusiness(ctx context.Context, newBusiness Business) (*Business, error) {

	query := `
	INSERT INTO Businesses (
							business_name,
							business_type,
							business_phone_number,
							business_email,
							business_address,
							postal_code,
							city,
							notes)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	REURNING id`

	err := pDb.pool.QueryRow(
		ctx,
		query,
		newBusiness.BusinessName,
		newBusiness.BusinessType,
		newBusiness.BusinessPhoneNumber,
		newBusiness.BusinessEmail,
		newBusiness.BusinessAddress,
		newBusiness.PostalCode,
		newBusiness.City,
		newBusiness.Notes,
	).Scan(&newBusiness.ID)

	if err != nil {
		return nil, err
	}

	return &newBusiness, nil
}
