package testdataPsql

import (
	database "github.com/Slowers-team/Slowers-App/database/psql"
)

func GetBusinesses() []database.Business {
	return []database.Business{
		{
			ID:             1,
			BusinessName:   "Test Business",
			BusinessIdCode: "1234567-8",
			Type:           "grower",
			PhoneNumber:    "010234567",
			Email:          "tester@test.fi",
			Address:        "Imaginary road 1",
			PostalCode:     "98765",
			City:           "Flowertown",
			AdditionalInfo: "No notes",
		},
	}
}
