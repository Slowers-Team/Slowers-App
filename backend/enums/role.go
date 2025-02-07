package enums

import "errors"

type Role struct {
	slug string
}

func (r Role) String() string {
	return r.slug
}

var (
	Unknown       = Role{""}
	Grower        = Role{"grower"}
	Retailer      = Role{"retailer"}
	GrowerOwner   = Role{"growerowner"}
	RetailerOwner = Role{"retailerowner"}
)

func RoleFromString(s string) (Role, error) {
	switch s {
	case Grower.slug:
		return Grower, nil
	case Retailer.slug:
		return Retailer, nil
	case GrowerOwner.slug:
		return GrowerOwner, nil
	case RetailerOwner.slug:
		return RetailerOwner, nil
	}
	return Unknown, errors.New("unknown role: " + s)
}
