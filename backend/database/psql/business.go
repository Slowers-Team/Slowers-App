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
	Delivery       string
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
							additional_info,
							delivery)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
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
		newBusiness.Delivery,
	).Scan(&newBusiness.ID)

	if err != nil {
		return nil, err
	}

	return &newBusiness, nil
}

func (pDb SQLDatabase) GetBusinessByUserID(ctx context.Context, userID int) (*Business, error) {
	business := new(Business)

	query := `
		SELECT
			Businesses.id,
			Businesses.name,
			Businesses.type,
			Businesses.phone_number,
			Businesses.email,
			Businesses.postal_code,
			Businesses.city,
			Businesses.address,
			Businesses.business_id_code,
			Businesses.created_at::TEXT,
			Businesses.additional_info,
			Businesses.delivery
		FROM
			Businesses
			INNER JOIN Memberships ON Businesses.id = Memberships.business_id
			INNER JOIN Users       ON Users.email = Memberships.user_email
		WHERE
			Users.id = $1`
	err := pDb.pool.QueryRow(ctx, query, userID).Scan(
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
		&business.Delivery,
	)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return business, nil
}
