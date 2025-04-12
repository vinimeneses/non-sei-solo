package entities_test

import (
	"api/internal/domain/entities"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func parseDate(dateStr string) time.Time {
	t, err := time.Parse("2006-02-01", dateStr)
	if err != nil {
		panic("Invalid date: " + err.Error())
	}
	return t
}

func formatDate(t time.Time) string {
	return t.Format("2006-02-01")
}

func TestNewPerson(t *testing.T) {
	name := "John"
	surname := "Doe"
	birthday := parseDate("1990-01-01")
	citizenship := "USA"
	taxCode := "ABC123"
	var familyID uint = 1
	phoneNumber := "1234567890"
	address := "123 Main St"
	contribution := 100.0

	person := entities.NewPerson(name, surname, birthday, citizenship, taxCode, &familyID, phoneNumber, address, contribution)

	assert.Equal(t, name, person.Name)
	assert.Equal(t, surname, person.Surname)
	assert.Equal(t, citizenship, person.Citizenship)
	assert.Equal(t, taxCode, person.TaxCode)
	assert.Equal(t, &familyID, person.FamilyID)
	assert.Equal(t, phoneNumber, person.PhoneNumber)
	assert.Equal(t, address, person.Address)
	assert.Equal(t, contribution, person.Contribution)
}

func TestNewPersonWithNilFamilyID(t *testing.T) {
	person := entities.NewPerson("John", "Doe", parseDate("1990-01-01"), "USA", "ABC123", nil, "1234567890", "123 Main St", 100.0)
	assert.Nil(t, person.FamilyID)
}

func TestPersonValidation(t *testing.T) {
	validPerson := entities.NewPerson(
		"John", "Doe",
		parseDate("1990-01-01"),
		"USA", "ABC123",
		nil,
		"1234567890", "123 Main St",
		100.0,
	)

	tests := []struct {
		name          string
		modifyPerson  func(*entities.Person)
		expectedError string
	}{
		{
			name: "Valid person",
			modifyPerson: func(p *entities.Person) {
				// No change
			},
			expectedError: "",
		},
		{
			name: "Empty name",
			modifyPerson: func(p *entities.Person) {
				p.Name = ""
			},
			expectedError: "name is required",
		},
		{
			name: "Name too short",
			modifyPerson: func(p *entities.Person) {
				p.Name = "Jo"
			},
			expectedError: "name must be at least 3 characters long",
		},
		{
			name: "Empty surname",
			modifyPerson: func(p *entities.Person) {
				p.Surname = ""
			},
			expectedError: "surname is required",
		},
		{
			name: "Zero birthday",
			modifyPerson: func(p *entities.Person) {
				p.Birthday = time.Time{}
			},
			expectedError: "birthday is required",
		},
		{
			name: "Future birthday",
			modifyPerson: func(p *entities.Person) {
				p.Birthday = parseDate(fmt.Sprintf("%d-01-01", time.Now().Year()+1))
			},
			expectedError: "birthday cannot be in the future",
		},
		{
			name: "Empty citizenship",
			modifyPerson: func(p *entities.Person) {
				p.Citizenship = ""
			},
			expectedError: "citizenship is required",
		},
		{
			name: "Negative contribution",
			modifyPerson: func(p *entities.Person) {
				p.Contribution = -10.0
			},
			expectedError: "contribution cannot be negative",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testPerson := *validPerson
			tt.modifyPerson(&testPerson)

			err := testPerson.UpdateName(testPerson.Name)

			if tt.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
			}
		})
	}
}

func TestUpdateMethods(t *testing.T) {
	basePerson := entities.NewPerson(
		"John", "Doe", parseDate("1990-01-01"),
		"USA", "ABC123", nil, "1234567890", "123 Main St", 100.0,
	)

	t.Run("UpdateName", func(t *testing.T) {
		person := *basePerson

		err := person.UpdateName("Jane")
		assert.NoError(t, err)
		assert.Equal(t, "Jane", person.Name)

		err = person.UpdateName("")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "name is required")
	})

	t.Run("UpdateSurname", func(t *testing.T) {
		person := *basePerson

		err := person.UpdateSurname("Smith")
		assert.NoError(t, err)
		assert.Equal(t, "Smith", person.Surname)

		err = person.UpdateSurname("")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "surname is required")
	})

	t.Run("UpdateBirthday", func(t *testing.T) {
		person := *basePerson

		err := person.UpdateBirthday(parseDate("1985-05-05"))
		assert.NoError(t, err)
		assert.Equal(t, "1985-05-05", formatDate(person.Birthday))

		err = person.UpdateBirthday(time.Time{})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "birthday is required")

		err = person.UpdateBirthday(parseDate(fmt.Sprintf("%d-01-01", time.Now().Year()+1)))
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "birthday cannot be in the future")
	})

	t.Run("UpdateCitizenship", func(t *testing.T) {
		person := *basePerson

		err := person.UpdateCitizenship("Canada")
		assert.NoError(t, err)
		assert.Equal(t, "Canada", person.Citizenship)

		err = person.UpdateCitizenship("")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "citizenship is required")
	})

	t.Run("UpdateTaxCode", func(t *testing.T) {
		person := *basePerson

		err := person.UpdateTaxCode("XYZ789")
		assert.NoError(t, err)
		assert.Equal(t, "XYZ789", person.TaxCode)

		err = person.UpdateTaxCode("")
		assert.NoError(t, err)
		assert.Equal(t, "", person.TaxCode)
	})

	t.Run("UpdateFamilyID", func(t *testing.T) {
		person := *basePerson

		var familyID uint = 2
		err := person.UpdateFamilyID(&familyID)
		assert.NoError(t, err)
		assert.Equal(t, &familyID, person.FamilyID)

		err = person.UpdateFamilyID(nil)
		assert.NoError(t, err)
		assert.Nil(t, person.FamilyID)
	})

	t.Run("UpdatePhoneNumber", func(t *testing.T) {
		person := *basePerson

		err := person.UpdatePhoneNumber("9876543210")
		assert.NoError(t, err)
		assert.Equal(t, "9876543210", person.PhoneNumber)

		err = person.UpdatePhoneNumber("")
		assert.NoError(t, err)
		assert.Equal(t, "", person.PhoneNumber)
	})

	t.Run("UpdateAddress", func(t *testing.T) {
		person := *basePerson

		err := person.UpdateAddress("456 Oak Ave")
		assert.NoError(t, err)
		assert.Equal(t, "456 Oak Ave", person.Address)

		err = person.UpdateAddress("")
		assert.NoError(t, err)
		assert.Equal(t, "", person.Address)
	})

	t.Run("UpdateContribution", func(t *testing.T) {
		person := *basePerson

		err := person.UpdateContribution(200.0)
		assert.NoError(t, err)
		assert.Equal(t, 200.0, person.Contribution)

		err = person.UpdateContribution(0.0)
		assert.NoError(t, err)
		assert.Equal(t, 0.0, person.Contribution)

		err = person.UpdateContribution(-50.0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "contribution cannot be negative")
	})
}
