package domain

import "time"

type Family struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description"`
	Members     []Person  `json:"members,omitempty" gorm:"foreignKey:FamilyID"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
