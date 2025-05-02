package entities_test

import (
	"api/internal/domain/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewValidatedPerson_WithValidPerson_ReturnsValidatedPerson(t *testing.T) {
	person := entities.NewPerson(
		"John",
		"Doe",
		parseDate("1990-01-01"),
		"USA",
		"123-45-6789",
		nil,
		"555-1234",
		"123 Main St",
		100.0,
	)

	validatedPerson, err := entities.NewValidatedPerson(person)

	assert.NoError(t, err)
	assert.NotNil(t, validatedPerson)
	assert.True(t, validatedPerson.IsValid())
	assert.Equal(t, "John", validatedPerson.Name)
	assert.Equal(t, "Doe", validatedPerson.Surname)
}

func TestNewValidatedPerson_WithInvalidPerson_ReturnsError(t *testing.T) {
	invalidPerson := entities.NewPerson(
		"", // Invalid name
		"Doe",
		parseDate("1990-01-01"),
		"USA",
		"123-45-6789",
		nil,
		"555-1234",
		"123 Main St",
		100.0,
	)

	validatedPerson, err := entities.NewValidatedPerson(invalidPerson)

	assert.Nil(t, validatedPerson)
	assert.Error(t, err)
}
