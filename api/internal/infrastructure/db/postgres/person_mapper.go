package postgres

import (
	"api/internal/domain/entities"
	"gorm.io/gorm"
)

func toDBPerson(validatedPerson *entities.ValidatedPerson) *Person {
	return &Person{
		Model:        gorm.Model{},
		Name:         validatedPerson.Name,
		Surname:      validatedPerson.Surname,
		Birthday:     validatedPerson.Birthday,
		Citizenship:  validatedPerson.Citizenship,
		TaxCode:      validatedPerson.TaxCode,
		FamilyID:     validatedPerson.FamilyID,
		PhoneNumber:  validatedPerson.PhoneNumber,
		Address:      validatedPerson.Address,
		Contribution: validatedPerson.Contribution,
	}
}

func toDomainPerson(dbPerson *Person) *entities.Person {
	return &entities.Person{
		ID:           dbPerson.ID,
		Name:         dbPerson.Name,
		Surname:      dbPerson.Surname,
		Birthday:     dbPerson.Birthday,
		Citizenship:  dbPerson.Citizenship,
		TaxCode:      dbPerson.TaxCode,
		FamilyID:     dbPerson.FamilyID,
		PhoneNumber:  dbPerson.PhoneNumber,
		Address:      dbPerson.Address,
		Contribution: dbPerson.Contribution,
	}
}
