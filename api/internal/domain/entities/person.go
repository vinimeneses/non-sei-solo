package entities

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Person struct {
	gorm.Model
	Name         string    `json:"name" validate:"required,min=3"`
	Surname      string    `json:"surname" validate:"required"`
	Birthday     time.Time `json:"birthday" validate:"required"`
	Citizenship  string    `json:"citizenship" validate:"required"`
	TaxCode      string    `json:"tax_code"`
	FamilyID     *uint     `json:"family_id,omitempty" gorm:"index"`
	Family       *Family   `json:"family,omitempty" gorm:"foreignKey:FamilyID"`
	PhoneNumber  string    `json:"phone_number"`
	Address      string    `json:"address"`
	Contribution float64   `json:"contribution"`
}

func (p *Person) validate() error {
	if p.Name == "" {
		return fmt.Errorf("name is required")
	}
	if len(p.Name) < 3 {
		return fmt.Errorf("name must be at least 3 characters long")
	}
	if p.Surname == "" {
		return fmt.Errorf("surname is required")
	}
	if p.Birthday.IsZero() {
		return fmt.Errorf("birthday is required")
	}
	if p.Birthday.After(time.Now()) {
		return fmt.Errorf("birthday cannot be in the future")
	}
	if p.Citizenship == "" {
		return fmt.Errorf("citizenship is required")
	}

	if p.Contribution < 0 {
		return fmt.Errorf("contribution cannot be negative")
	}
	return nil
}

func NewPerson(name string, surname string, birthday time.Time, citizenship string,
	taxCode string, familyID *uint, phoneNumber string, address string, contribution float64) *Person {
	return &Person{
		Name:         name,
		Surname:      surname,
		Birthday:     birthday,
		Citizenship:  citizenship,
		TaxCode:      taxCode,
		FamilyID:     familyID,
		PhoneNumber:  phoneNumber,
		Address:      address,
		Contribution: contribution,
	}
}

func (p *Person) UpdateName(name string) error {
	p.Name = name
	p.UpdatedAt = time.Now()

	return p.validate()
}

func (p *Person) UpdateSurname(surname string) error {
	p.Surname = surname
	p.UpdatedAt = time.Now()

	return p.validate()
}

func (p *Person) UpdateBirthday(birthday time.Time) error {
	p.Birthday = birthday
	p.UpdatedAt = time.Now()

	return p.validate()
}

func (p *Person) UpdateCitizenship(citizenship string) error {
	p.Citizenship = citizenship
	p.UpdatedAt = time.Now()

	return p.validate()
}

func (p *Person) UpdateTaxCode(taxCode string) error {
	p.TaxCode = taxCode
	p.UpdatedAt = time.Now()

	return p.validate()
}

func (p *Person) UpdateFamilyID(familyID *uint) error {
	p.FamilyID = familyID
	p.UpdatedAt = time.Now()

	return p.validate()
}

func (p *Person) UpdatePhoneNumber(phoneNumber string) error {
	p.PhoneNumber = phoneNumber
	p.UpdatedAt = time.Now()

	return p.validate()
}

func (p *Person) UpdateAddress(address string) error {
	p.Address = address
	p.UpdatedAt = time.Now()

	return p.validate()
}

func (p *Person) UpdateContribution(contribution float64) error {
	p.Contribution = contribution
	p.UpdatedAt = time.Now()

	return p.validate()
}
