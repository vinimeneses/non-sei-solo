package entities

import (
	"fmt"
)

type Family struct {
	ID          uint
	Members     []Person
	Name        string
	Attachments []Attachment
	Description string
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

func NewFamily(name string, members []Person, attachments []Attachment, description string) *Family {
	return &Family{
		Name:        name,
		Members:     members,
		Attachments: attachments,
		Description: description,
	}
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
