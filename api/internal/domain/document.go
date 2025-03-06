package domain

import "time"

type Document struct {
	ID         uint       `json:"id" gorm:"primaryKey"`
	PersonID   uint       `json:"person_id" gorm:"index"`
	Type       string     `json:"type" validate:"required"`
	Path       string     `json:"path" validate:"required"`
	Attachment Attachment `json:"attachment,omitempty" gorm:"foreignKey:DocumentID"`
	CreatedAt  time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
}
