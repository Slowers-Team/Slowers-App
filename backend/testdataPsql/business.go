package testdataPsql

import (
	database "github.com/Slowers-team/Slowers-App/database/psql"
)

func GetBusinesses() []database.Business {
	return []database.Business{
		{
			ID:           1,
			BusinessName: "Test Business",
			Type:         "growing",
			PhoneNumber:  "010234567",
			Email:        "tester@test.fi",
			PostAddress:  "Imaginary road 1",
			PostalCode:   "98765",
			City:         "Flowertown",
			Notes:        "No notes",
		},
	}
}
