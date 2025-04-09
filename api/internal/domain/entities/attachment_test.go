package entities_test

import (
	"api/internal/domain/entities"
	"strings"
	"testing"
)

func TestNewValidatedAttachment_Valid(t *testing.T) {
	attachment := &entities.Attachment{
		Path: "https://drive.google.com/file/d/12345/view",
	}

	validated, err := entities.NewValidatedAttachment(attachment)
	if err != nil {
		t.Fatalf("expected valid attachment, got error: %v", err)
	}

	if !validated.IsValid() {
		t.Fatal("expected attachment to be marked as valid")
	}

	if validated.Path != attachment.Path {
		t.Errorf("expected path %q, got %q", attachment.Path, validated.Path)
	}
}

func TestNewValidatedAttachment_EmptyPath(t *testing.T) {
	attachment := &entities.Attachment{Path: ""}
	_, err := entities.NewValidatedAttachment(attachment)
	if err == nil {
		t.Fatal("expected error for empty path, got nil")
	}
}

func TestNewValidatedAttachment_InvalidURL(t *testing.T) {
	attachment := &entities.Attachment{Path: "not a valid url"}
	_, err := entities.NewValidatedAttachment(attachment)
	if err == nil {
		t.Fatal("expected error for invalid URL, got nil")
	}
}

func TestValidatedAttachment_IsValid(t *testing.T) {
	attachment := &entities.Attachment{Path: "https://drive.google.com/file/d/12345/view"}
	validated, err := entities.NewValidatedAttachment(attachment)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !validated.IsValid() {
		t.Error("expected attachment to be valid")
	}
}

func TestAttachment_ValidateMethod(t *testing.T) {
	a := &entities.Attachment{Path: "https://example.com/file"}
	err := a.Validate()
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	a2 := &entities.Attachment{Path: ""}
	err = a2.Validate()
	if err == nil || err.Error() != "URL is required" {
		t.Errorf("expected 'URL is required' error, got %v", err)
	}

	a3 := &entities.Attachment{Path: "ht!tp"}
	err = a3.Validate()
	if err == nil || !strings.Contains(err.Error(), "invalid URL:") {
		t.Errorf("expected invalid URL error, got %v", err)
	}
}

func TestNewAttachment_ValidURL(t *testing.T) {
	a, err := entities.NewAttachment("https://example.com/test")
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if a.Path != "https://example.com/test" {
		t.Errorf("expected path to be %q, got %q", "https://example.com/test", a.Path)
	}
}

func TestNewAttachment_InvalidURL(t *testing.T) {
	_, err := entities.NewAttachment("ht!tp://invalid-url")
	if err == nil || !strings.Contains(err.Error(), "invalid URL:") {
		t.Fatalf("expected invalid URL error, got: %v", err)
	}
}
