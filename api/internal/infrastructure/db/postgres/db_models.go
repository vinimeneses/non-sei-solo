package postgres

import (
	"time"

	"gorm.io/gorm"
)

type Attachment struct {
	ID       uint `gorm:"primaryKey"`
	Path     string
	FamilyID uint
}

type Family struct {
	gorm.Model
	Name        string
	Members     []Person     `gorm:"foreignKey:FamilyID"`
	Attachments []Attachment `gorm:"foreignKey:FamilyID"`
	Description string       `gorm:"-"`
}

type Person struct {
	gorm.Model
	Name         string
	Surname      string
	Birthday     time.Time
	Citizenship  string
	TaxCode      string
	FamilyID     *uint
	Family       Family `gorm:"foreignKey:FamilyID"`
	PhoneNumber  string
	Address      string
	Contribution float64
}
