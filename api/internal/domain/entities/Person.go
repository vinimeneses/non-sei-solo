package entities

import (
	"fmt"
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	Name          string  `json:"name" validate:"required,min=3"`
	Surname       string  `json:"surname" validate:"required"`
	Birthday      string  `json:"birthday" validate:"required"`
	Citizenship   string  `json:"citizenship" validate:"required"`
	TaxCode       string  `json:"tax_code"`
	FamilyID      *uint   `json:"family_id,omitempty" gorm:"index"`
	Family        *Family `json:"family,omitempty" gorm:"foreignKey:FamilyID"`
	PhoneNumber   string  `json:"phone_number"`
	Address       string  `json:"address"`
	Contribuition float64 `json:"contribuition"`
}

func (p *Person) validate() error {
	if p.Name == "" {
		return fmt.Errorf("name is required")
	}
	if len(p.Name) < 3 {
		return fmt.Errorf("name must be at least 3 characters long")
	}
	return nil
}
