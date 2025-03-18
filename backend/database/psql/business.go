package database

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
