package database

import (
	"context"
	"fmt"
)

type Business struct {
	ID             int
	CreatedAt      string
	LastModified   string
	BusinessName   string
	BusinessIdCode string
	Type           string
	PhoneNumber    string
	Email          string
	Address        string
	PostalCode     string
	City           string
	AdditionalInfo string
}

func (pDb SQLDatabase) CreateBusiness(ctx context.Context, newBusiness Business) (*Business, error) {
	fmt.Println("kysely")
	query := `
	INSERT INTO Businesses (
							name,
							business_id_code,
							type,
							phone_number,
							email,
							address,
							postal_code,
							city,
							additional_info)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	RETURNING id`

	err := pDb.pool.QueryRow(
		ctx,
		query,
		newBusiness.BusinessName,
		newBusiness.BusinessIdCode,
		newBusiness.Type,
		newBusiness.PhoneNumber,
		newBusiness.Email,
		newBusiness.Address,
		newBusiness.PostalCode,
		newBusiness.City,
		newBusiness.AdditionalInfo,
	).Scan(&newBusiness.ID)

	if err != nil {
		return nil, err
	}

	return &newBusiness, nil
}
