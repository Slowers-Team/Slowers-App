package enums

import "errors"

type Designation struct {
	slug string
}

func (r Designation) String() string {
	return r.slug
}

var (
	NoDesignation = Designation{""}
	Owner         = Designation{"owner"}
	Employee      = Designation{"employee"}
)

func DesignationFromString(s string) (Designation, error) {
	switch s {
	case Owner.slug:
		return Owner, nil
	case Employee.slug:
		return Employee, nil
	}
	return NoDesignation, errors.New("unknown designation: " + s)
}
