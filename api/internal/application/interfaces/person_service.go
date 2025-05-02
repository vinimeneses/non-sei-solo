package interfaces

import (
	"api/internal/application/command"
	"api/internal/application/query"
)

type PersonService interface {
	CreatePerson(person *command.CreatePersonCommand) (*command.CreatePersonCommandResult, error)
	FindAllPersons() (*query.PersonQueryListResult, error)
	FindPersonById(id uint) (*query.PersonQueryResult, error)
	UpdatePerson(person *command.UpdatePersonCommand) (*command.UpdatePersonCommandResult, error)
	DeletePerson(person *command.DeletePersonCommand) error
}
