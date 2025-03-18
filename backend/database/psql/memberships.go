package database

type Membership struct {
	UserID      int
	BusinessID  int
	Designation string
}

func (Db SQLDatabase) AddMembership(user_email string, business_email string, designation string) error {
	return nil
}
