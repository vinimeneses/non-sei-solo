package command

import (
	"api/internal/application/common"
	"time"
)

type CreatePersonCommand struct {
	Name         string
	Surname      string
	Birthday     time.Time
	Citizenship  string
	TaxCode      string
	PhoneNumber  string
	Address      string
	Contribution float64
}

type CreatePersonCommandResult struct {
	Result *common.PersonResult
}
