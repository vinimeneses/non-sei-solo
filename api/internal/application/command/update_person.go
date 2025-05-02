package command

import (
	"api/internal/application/common"
	"time"
)

type UpdatePersonCommand struct {
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

type UpdatePersonCommandResult struct {
	Result *common.PersonResult
}
