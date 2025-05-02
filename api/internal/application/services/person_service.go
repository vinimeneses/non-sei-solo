package services

import (
	"api/internal/application/command"
	"api/internal/application/interfaces"
	"api/internal/application/mapper"
	"api/internal/application/query"
	"api/internal/domain/entities"
	"api/internal/domain/repositories"
	"errors"
)

type PersonService struct {
	repo repositories.PersonRepository
}

func NewPersonService(repo repositories.PersonRepository) interfaces.PersonService {
	return &PersonService{repo: repo}
}

func (p PersonService) CreatePerson(person *command.CreatePersonCommand) (*command.CreatePersonCommandResult, error) {

	var familyId *uint = nil
	var newPerson = entities.NewPerson(
		person.Name,
		person.Surname,
		person.Birthday,
		person.Citizenship,
		person.TaxCode,
		familyId,
		person.PhoneNumber,
		person.Address,
		person.Contribution,
	)

	validatedPerson, err := entities.NewValidatedPerson(newPerson)
	if err != nil {
		return nil, err
	}

	_, err = p.repo.Create(validatedPerson)
	if err != nil {
		return nil, err
	}

	result := command.CreatePersonCommandResult{
		Result: mapper.NewPersonResultFromValidatedEntity(validatedPerson),
	}

	return &result, nil
}

func (p PersonService) FindAllPersons() (*query.PersonQueryListResult, error) {
	storedPersons, err := p.repo.FindAll()
	if err != nil {
		return nil, err
	}

	var queryResult query.PersonQueryListResult
	for _, person := range storedPersons {
		queryResult.Result = append(queryResult.Result, mapper.NewPersonResultFromEntity(person))
	}
	return &queryResult, nil
}

func (p PersonService) FindPersonById(id uint) (*query.PersonQueryResult, error) {
	person, err := p.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	if person == nil {
		return nil, nil
	}

	result := &query.PersonQueryResult{
		Result: mapper.NewPersonResultFromEntity(person),
	}

	return result, nil
}

func (p PersonService) UpdatePerson(person *command.UpdatePersonCommand) (*command.UpdatePersonCommandResult, error) { //TODO implement me
	storedPerson, err := p.repo.FindById(person.ID)

	if err != nil {
		return nil, errors.New("person not found")
	}

	if storedPerson == nil {
		return nil, nil
	}

	if err := storedPerson.UpdateName(person.Name); err != nil {
		return nil, err
	}

	if err := storedPerson.UpdateSurname(person.Surname); err != nil {
		return nil, err
	}

	if err := storedPerson.UpdateBirthday(person.Birthday); err != nil {
		return nil, err
	}

	if err := storedPerson.UpdateCitizenship(person.Citizenship); err != nil {
		return nil, err
	}

	if err := storedPerson.UpdateTaxCode(person.TaxCode); err != nil {
		return nil, err
	}

	if err := storedPerson.UpdateFamilyID(person.FamilyID); err != nil {
		return nil, err
	}

	if err := storedPerson.UpdatePhoneNumber(person.PhoneNumber); err != nil {
		return nil, err
	}

	if err := storedPerson.UpdateAddress(person.Address); err != nil {
		return nil, err
	}

	if err := storedPerson.UpdateContribution(person.Contribution); err != nil {
		return nil, err
	}

	validatedUpdatedPerson, err := entities.NewValidatedPerson(storedPerson)
	if err != nil {
		return nil, err
	}

	_, err = p.repo.Update(validatedUpdatedPerson)
	if err != nil {
		return nil, err
	}
	result := command.UpdatePersonCommandResult{
		Result: mapper.NewPersonResultFromValidatedEntity(validatedUpdatedPerson),
	}

	return &result, nil
}

func (p PersonService) DeletePerson(person *command.DeletePersonCommand) error {
	err := p.repo.Delete(person.ID)
	if err != nil {
		return err
	}
	return nil
}
