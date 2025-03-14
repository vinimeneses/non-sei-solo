package domain

import "time"

type AttachmentType string

const (
	DocumentAttachment AttachmentType = "DOCUMENT"
	ImageAttachment    AttachmentType = "IMAGE"
)

type Attachment struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	DocumentID uint           `json:"document_id,omitempty" gorm:"index"`
	Type       AttachmentType `json:"type" validate:"required,oneof=DOCUMENT IMAGE"`
	Path       string         `json:"path" validate:"required"`
	CreatedAt  time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
}
