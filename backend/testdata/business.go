package testdata

import "github.com/Slowers-team/Slowers-App/databases/sql"

func GetBusinesses() []sql.Business {
	return []sql.Business{
		{
			ID:             1,
			BusinessName:   "Test Business",
			BusinessIdCode: "1234567-8",
			Type:           "grower",
			PhoneNumber:    "0101234567",
			Email:          "tester@test.fi",
			Address:        "Imaginary road 1",
			PostalCode:     "98765",
			City:           "Flowertown",
			AdditionalInfo: "No notes",
		},
	}
}
