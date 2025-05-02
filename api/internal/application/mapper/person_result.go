package mapper

import (
	"api/internal/application/common"
	"api/internal/domain/entities"
)

func NewPersonResultFromValidatedEntity(person *entities.ValidatedPerson) *common.PersonResult {
	return NewPersonResultFromEntity(&person.Person)
}

func NewPersonResultFromEntity(person *entities.Person) *common.PersonResult {
	if person == nil {
		return nil
	}

	return &common.PersonResult{
		ID:           person.ID,
		Name:         person.Name,
		Surname:      person.Surname,
		Birthday:     person.Birthday,
		Citizenship:  person.Citizenship,
		TaxCode:      person.TaxCode,
		FamilyID:     person.FamilyID,
		PhoneNumber:  person.PhoneNumber,
		Address:      person.Address,
		Contribution: person.Contribution,
	}
}
