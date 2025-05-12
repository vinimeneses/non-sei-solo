package services

import (
	"api/internal/application/command"
	"api/internal/domain/entities"
	"testing"
	"time"
)

// MockPersonRepository is a mock implementation of the PersonRepository interface
type MockPersonRepository struct {
	person []*entities.ValidatedPerson
}

func (m *MockPersonRepository) Create(person *entities.ValidatedPerson) (*entities.Person, error) {
	m.person = append(m.person, person)
	return &person.Person, nil
}

func (m *MockPersonRepository) FindById(id uint) (*entities.Person, error) {
	for _, p := range m.person {
		if p.ID == id {
			return &p.Person, nil
		}
	}
	return nil, nil
}

func (m *MockPersonRepository) FindAll() ([]*entities.Person, error) {
	var persons []*entities.Person
	for _, p := range m.person {
		persons = append(persons, &p.Person)
	}
	return persons, nil
}

func (m *MockPersonRepository) Update(person *entities.ValidatedPerson) (*entities.Person, error) {
	for i, p := range m.person {
		if p.ID == person.ID {
			m.person[i] = person
			return &person.Person, nil
		}
	}
	return nil, nil
}

func (m *MockPersonRepository) Delete(id uint) error {
	for i, p := range m.person {
		if p.ID == id {
			m.person = append(m.person[:i], m.person[i+1:]...)
			return nil
		}
	}
	return nil
}

func TestPersonService_CreatePerson(t *testing.T) {
	personRepo := &MockPersonRepository{}
	service := NewPersonService(personRepo)

	// Create Person
	person := createPersistedPerson(t, personRepo)
	personCommand := getCreatePersonCommand(person)
	_, err := service.CreatePerson(personCommand)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
}

func getCreatePersonCommand(person *entities.ValidatedPerson) *command.CreatePersonCommand {
	return &command.CreatePersonCommand{
		Name:         person.Name,
		Surname:      person.Surname,
		Birthday:     person.Birthday,
		Citizenship:  person.Citizenship,
		TaxCode:      person.TaxCode,
		PhoneNumber:  person.PhoneNumber,
		Address:      person.Address,
		Contribution: person.Contribution,
	}
}

func createPersistedPerson(t *testing.T, personRepo *MockPersonRepository) *entities.ValidatedPerson {
	person := entities.NewPerson("John", "Doe", mustParseDate("1990-01-01"), "USA",
		"ABC123", nil, "1234567890", "123 Main St", 100.0)
	validatedPerson, err := entities.NewValidatedPerson(person)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	_, err = personRepo.Create(validatedPerson)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	return validatedPerson
}

// helper para converter string para time.Time
func mustParseDate(value string) time.Time {
	t, err := time.Parse("2006-01-02", value)
	if err != nil {
		panic(err)
	}
	return t
}
