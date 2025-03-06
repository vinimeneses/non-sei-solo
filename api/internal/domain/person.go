package domain

import "time"

type Person struct {
	ID           uint       `json:"id" gorm:"primaryKey"`
	Name         string     `json:"name" validate:"required"`
	Surname      string     `json:"surname" validate:"required"`
	Birthday     time.Time  `json:"birthday" validate:"required"`
	Citizenship  string     `json:"citizenship"`
	TaxCode      string     `json:"tax_code"`
	FamilyID     *uint      `json:"family_id,omitempty" gorm:"index"`
	Family       *Family    `json:"family,omitempty" gorm:"foreignKey:FamilyID"`
	PhoneNumber  string     `json:"phone_number"`
	Address      string     `json:"address"`
	Documents    []Document `json:"documents,omitempty" gorm:"foreignKey:PersonID"`
	CreatedBy    uint       `json:"created_by"`
	CreatedAt    time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	OfertaStatus string     `json:"oferta_status"`
}
