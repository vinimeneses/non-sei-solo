package entities

import (
	"strings"
	"testing"
	"time"
)

func parseDate(dateStr string) time.Time {
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		panic("Invalid date: " + err.Error())
	}
	return t
}

func validValidatedPerson() *ValidatedPerson {
	p := NewPerson(
		"John",
		"Doe",
		parseDate("1990-01-01"),
		"USA",
		"123456789",
		nil, // familyID
		"123-4567",
		"123 Street",
		100,
	)
	vp, err := NewValidatedPerson(p)
	if err != nil {
		panic("failed to validate person: " + err.Error())
	}
	return vp
}

func validatedPersonWithID(id uint) *ValidatedPerson {
	p := NewPerson(
		"John",
		"Doe",
		parseDate("1990-01-01"),
		"USA",
		"123456789",
		nil,
		"123-4567",
		"123 Street",
		100,
	)
	p.ID = id
	vp, err := NewValidatedPerson(p)
	if err != nil {
		panic("failed to validate person: " + err.Error())
	}
	return vp
}

func TestFamily_Validate_Valid(t *testing.T) {
	f := Family{
		Name:    "Smith Family",
		Members: []Person{validatedPersonWithID(1).Person},
	}
	if err := f.validate(); err != nil {
		t.Fatalf("expected valid family, got error: %v", err)
	}
}

func TestFamily_Validate_Invalid(t *testing.T) {
	cases := []struct {
		name    string
		family  Family
		wantErr string
	}{
		{"empty name", Family{Name: "", Members: []Person{validatedPersonWithID(1).Person}}, "name is required"},
		{"short name", Family{Name: "AB", Members: []Person{validatedPersonWithID(1).Person}}, "name must have at least 3 characters"},
		{"long name", Family{Name: string(make([]byte, 101)), Members: []Person{validatedPersonWithID(1).Person}}, "name must have less than 100 characters"},
		{"duplicate member ID", Family{
			Name:    "ValidName",
			Members: []Person{validatedPersonWithID(1).Person, validatedPersonWithID(1).Person},
		}, "duplicate member ID"},
	}

	for _, c := range cases {
		err := c.family.validate()
		if err == nil || err.Error() == "" || !strings.Contains(err.Error(), c.wantErr) {
			t.Errorf("case %q: expected error to contain %q, got %v", c.name, c.wantErr, err)
		}
	}
}

func TestNewFamily(t *testing.T) {
	name := "Test Family"
	members := []*ValidatedPerson{validatedPersonWithID(1), validatedPersonWithID(2)}

	attachment, err := NewAttachment("https://example.com/document.pdf")
	if err != nil {
		t.Fatalf("failed to create test attachment: %v", err)
	}
	vAttachment, err := NewValidatedAttachment(attachment)
	if err != nil {
		t.Fatalf("failed to validate attachment: %v", err)
	}
	attachments := []*ValidatedAttachment{vAttachment}

	description := "Family description text"

	family, err := NewFamily(name, members, attachments, description)
	if err != nil {
		t.Fatalf("failed to create family: %v", err)
	}

	if family.Name != name {
		t.Errorf("expected name %q, got %q", name, family.Name)
	}

	if len(family.Members) != len(members) {
		t.Errorf("expected %d members, got %d", len(members), len(family.Members))
	}

	memberIDs := map[uint]bool{}
	for _, m := range family.Members {
		memberIDs[m.ID] = true
	}
	if !memberIDs[1] || !memberIDs[2] {
		t.Error("not all expected members were added correctly")
	}

	if len(family.Attachments) != 1 {
		t.Errorf("expected 1 attachment, got %d", len(family.Attachments))
	}
	// Assuming Attachment has a field 'Path'
	if family.Attachments[0].Path != attachment.Path {
		t.Errorf("expected attachment path %q, got %q", attachment.Path, family.Attachments[0].Path)
	}

	if family.Description != description {
		t.Errorf("expected description %q, got %q", description, family.Description)
	}
}

func TestFamily_AddMember(t *testing.T) {
	members := []*ValidatedPerson{validatedPersonWithID(1)}
	family, err := NewFamily("Smith", members, nil, "")
	if err != nil {
		t.Fatalf("failed to create family: %v", err)
	}

	newVP := validatedPersonWithID(2)
	err = family.AddMember(newVP)
	if err != nil {
		t.Fatalf("expected to add valid member, got error: %v", err)
	}
	if len(family.Members) != 2 {
		t.Fatalf("expected 2 members, got %d", len(family.Members))
	}

	found := false
	for _, member := range family.Members {
		if member.ID == 2 {
			found = true
			break
		}
	}
	if !found {
		t.Fatal("added member not found in family members")
	}
}

func TestFamily_AddMember_Invalid(t *testing.T) {
	members := []*ValidatedPerson{validatedPersonWithID(1)}
	family, err := NewFamily("Smith", members, nil, "")
	if err != nil {
		t.Fatalf("failed to create family: %v", err)
	}

	vp := validValidatedPerson()
	vp.Person.Name = "" // Invalidate the person by clearing the name

	err = family.AddMember(vp)
	if err == nil {
		t.Fatal("expected error for invalid member, got nil")
	}
	if !strings.Contains(err.Error(), "name is required") {
		t.Fatalf("expected error to mention invalid name, got: %v", err)
	}
}

func TestFamily_RemoveMemberByID(t *testing.T) {
	// Create a family with 3 members.
	members := []*ValidatedPerson{
		validatedPersonWithID(1),
		validatedPersonWithID(2),
		validatedPersonWithID(3),
	}
	family, err := NewFamily("Test Family", members, nil, "")
	if err != nil {
		t.Fatalf("failed to create family: %v", err)
	}

	if len(family.Members) != 3 {
		t.Fatalf("setup failed: expected 3 members, got %d", len(family.Members))
	}

	removed := family.RemoveMemberByID(2)
	if !removed {
		t.Fatal("expected member to be removed, got false")
	}

	if len(family.Members) != 2 {
		t.Fatalf("expected 2 members after removal, got %d", len(family.Members))
	}

	for _, m := range family.Members {
		if m.ID == 2 {
			t.Fatal("member with ID 2 was not removed")
		}
	}

	foundIDs := map[uint]bool{}
	for _, m := range family.Members {
		foundIDs[m.ID] = true
	}
	if !foundIDs[1] || !foundIDs[3] {
		t.Fatal("expected members with IDs 1 and 3 to remain")
	}
}

func TestFamily_RemoveMemberByID_NotFound(t *testing.T) {
	members := []*ValidatedPerson{validatedPersonWithID(1), validatedPersonWithID(2)}
	family, err := NewFamily("Test", members, nil, "")
	if err != nil {
		t.Fatalf("failed to create family: %v", err)
	}

	removed := family.RemoveMemberByID(999) // Non-existent ID
	if removed {
		t.Fatal("expected removal to return false for nonexistent member")
	}

	if len(family.Members) != 2 {
		t.Fatalf("expected still 2 members, got %d", len(family.Members))
	}
}

func TestFamily_UpdateName(t *testing.T) {
	family, err := NewFamily("OldName", []*ValidatedPerson{}, nil, "")
	if err != nil {
		t.Fatalf("failed to create family: %v", err)
	}
	err = family.UpdateName("NewName")
	if err != nil {
		t.Fatalf("expected successful update, got error: %v", err)
	}
	if family.Name != "NewName" {
		t.Fatalf("expected name to be updated, got: %s", family.Name)
	}
}

func TestFamily_UpdateName_Invalid(t *testing.T) {
	family, err := NewFamily("OldName", []*ValidatedPerson{}, nil, "")
	if err != nil {
		t.Fatalf("failed to create family: %v", err)
	}

	err = family.UpdateName("AB")
	if err == nil {
		t.Fatal("expected error for short name, got nil")
	}
	if !strings.Contains(err.Error(), "must have at least 3 characters") {
		t.Fatalf("unexpected error message: %v", err)
	}

	if family.Name != "OldName" {
		t.Fatalf("name should not have been updated, got: %s", family.Name)
	}
}

func TestFamily_AddAttachment(t *testing.T) {
	family, err := NewFamily("Souza", []*ValidatedPerson{}, nil, "")
	if err != nil {
		t.Fatalf("failed to create family: %v", err)
	}
	attachment, err := NewAttachment("https://drive.google.com/file/d/abc123")
	if err != nil {
		t.Fatalf("expected valid attachment, got error: %v", err)
	}
	vAttachment, err := NewValidatedAttachment(attachment)
	if err != nil {
		t.Fatalf("failed to validate attachment: %v", err)
	}
	family.AddAttachment(vAttachment)
	if len(family.Attachments) != 1 {
		t.Fatalf("expected 1 attachment, got %d", len(family.Attachments))
	}
}
