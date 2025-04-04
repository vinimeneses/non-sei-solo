package entities

import (
	"fmt"
	"gorm.io/gorm"
)

type Family struct {
	gorm.Model
	members     []Person
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

func (f *Family) validate() error {

	if f.Name == "" {
		return fmt.Errorf("name is required")
	}

	for i, member := range f.members {
		if err := member.validate(); err != nil {
			return fmt.Errorf("invalid member at index %d: %s", i, err.Error())
		}
	}
	return nil
}
