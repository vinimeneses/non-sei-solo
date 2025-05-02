package postgres

import (
	"time"

	"gorm.io/gorm"
)

type Family struct {
	gorm.Model
	Name        string
	Members     []Person     `gorm:"foreignKey:FamilyID;constraint:OnDelete:CASCADE"`
	Attachments []Attachment `gorm:"foreignKey:FamilyID;constraint:OnDelete:CASCADE"`
	Description string       `gorm:"-"` // ignored by gorm
}

type Person struct {
	gorm.Model
	Name         string
	Surname      string
	Birthday     time.Time
	Citizenship  string
	TaxCode      string
	FamilyID     *uint
	Family       Family
	PhoneNumber  string
	Address      string
	Contribution float64
}

type Attachment struct {
	ID       uint `gorm:"primaryKey"`
	Path     string
	FamilyID uint
}
