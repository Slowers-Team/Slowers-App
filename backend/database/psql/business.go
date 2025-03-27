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

func (pDb SQLDatabase) GetBusinessByUserID(ctx context.Context, businessID int) (*Business, error) {
	business := new(Business)

	query := `SELECT B.id, B.name,	B.type, B.phone_number, B.email, B.postal_code, B.city, B.address, B.business_id_code, B.created_at::TEXT, B.additional_info
			FROM Businesses B JOIN Memberships M ON B.id = M.business_id WHERE B.id = $1`
	err := pDb.pool.QueryRow(ctx, query, businessID).Scan(
		&business.ID,
		&business.BusinessName,
		&business.Type,
		&business.PhoneNumber,
		&business.Email,
		&business.PostalCode,
		&business.City,
		&business.Address,
		&business.BusinessIdCode,
		&business.CreatedAt,
		&business.AdditionalInfo,
	)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return business, nil
}
