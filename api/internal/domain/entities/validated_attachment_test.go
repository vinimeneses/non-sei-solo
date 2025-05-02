package entities_test

import (
	"api/internal/domain/entities"
	"testing"
)

func TestNewValidatedAttachment(t *testing.T) {
	t.Run("valid attachment", func(t *testing.T) {
		a := &entities.Attachment{Path: "https://drive.google.com/file/d/abc123"}
		validated, err := entities.NewValidatedAttachment(a)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if validated == nil || !validated.IsValid() {
			t.Fatal("expected validated attachment to be valid")
		}
	})

	t.Run("empty path", func(t *testing.T) {
		a := &entities.Attachment{Path: ""}
		_, err := entities.NewValidatedAttachment(a)
		if err == nil {
			t.Fatal("expected error for empty path, got nil")
		}
	})

	t.Run("invalid URL", func(t *testing.T) {
		a := &entities.Attachment{Path: "not-a-valid-url"}
		_, err := entities.NewValidatedAttachment(a)
		if err == nil {
			t.Fatal("expected error for invalid URL, got nil")
		}
	})
}
