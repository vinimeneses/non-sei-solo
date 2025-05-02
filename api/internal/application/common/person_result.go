package common

import "time"

type PersonResult struct {
	ID           uint
	Name         string
	Surname      string
	Birthday     time.Time
	Citizenship  string
	TaxCode      string
	FamilyID     *uint
	PhoneNumber  string
	Address      string
	Contribution float64
}
