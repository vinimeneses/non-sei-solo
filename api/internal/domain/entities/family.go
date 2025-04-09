package entities

import (
	"fmt"
	"gorm.io/gorm"
)

type Family struct {
	gorm.Model
	Members     []Person     `json:"members" validate:"required,min=1"`
	Name        string       `json:"name" validate:"required"`
	Attachments []Attachment `json:"attachments"`
	Description string       `json:"description" gorm:"-"`
}

func (f *Family) validate() error {
	if f.Name == "" {
		return fmt.Errorf("name is required")
	}
	if len(f.Name) < 3 {
		return fmt.Errorf("name must have at least 3 characters")
	}
	if len(f.Name) > 100 {
		return fmt.Errorf("name must have less than 100 characters")
	}
	if len(f.Members) == 0 {
		return fmt.Errorf("family must have at least one member")
	}

	seen := make(map[uint]bool)
	for _, member := range f.Members {
		if member.ID != 0 {
			if seen[member.ID] {
				return fmt.Errorf("duplicate member ID: %d", member.ID)
			}
			seen[member.ID] = true
		}
	}
	return nil
}

func (f *Family) AddMember(p Person) error {
	if err := p.validate(); err != nil {
		return fmt.Errorf("cannot add invalid person: %v", err)
	}
	f.Members = append(f.Members, p)
	return nil
}

func (f *Family) RemoveMemberByID(id uint) bool {
	for i, m := range f.Members {
		if m.ID == id {
			f.Members = append(f.Members[:i], f.Members[i+1:]...)
			return true
		}
	}
	return false
}

func (f *Family) UpdateName(newName string) error {
	if len(newName) < 3 {
		return fmt.Errorf("name must have at least 3 characters")
	}
	f.Name = newName
	return nil
}

func (f *Family) AddAttachment(a Attachment) {
	f.Attachments = append(f.Attachments, a)
}
