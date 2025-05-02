package entities_test

import (
	"api/internal/domain/entities"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func createValidPerson() entities.Person {
	return *entities.NewPerson(
		"John",
		"Doe",
		time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
		"USA",
		"123456789",
		nil,
		"555-1234",
		"123 Main St",
		100.0,
	)
}

func createValidFamily() *entities.Family {
	return &entities.Family{
		Name:    "The Smiths",
		Members: []entities.Person{createValidPerson()},
	}
}

func TestNewValidatedFamily_WithValidFamily(t *testing.T) {
	family := createValidFamily()

	validatedFamily, err := entities.NewValidatedFamily(family)

	assert.NoError(t, err)
	assert.NotNil(t, validatedFamily)
	assert.True(t, validatedFamily.IsValid())
	assert.Equal(t, "The Smiths", validatedFamily.Name)
	assert.Len(t, validatedFamily.Members, 1)
}

func TestNewValidatedFamily_WithInvalidFamily(t *testing.T) {
	family := &entities.Family{
		Name:    "", // invalid name
		Members: []entities.Person{createValidPerson()},
	}

	validatedFamily, err := entities.NewValidatedFamily(family)

	assert.Error(t, err)
	assert.Nil(t, validatedFamily)
}

func TestNewValidatedFamily_WithDuplicateMemberIDs(t *testing.T) {
	person := createValidPerson()
	person.ID = 1

	family := &entities.Family{
		Name:    "Duplicate IDs",
		Members: []entities.Person{person, person}, // same ID twice
	}

	validatedFamily, err := entities.NewValidatedFamily(family)

	assert.Error(t, err)
	assert.Nil(t, validatedFamily)
}
