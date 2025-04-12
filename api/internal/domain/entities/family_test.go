package entities

import (
	"strings"
	"testing"
	"time"
)

// Helper function to create a valid person for testing
func validPerson() *Person {
	return NewPerson(
		"John",
		"Doe",
		parseDate("1990-01-01"),
		"USA",
		"123456789",
		nil, // familyID
		"123-4567",
		"123 Street",
		100, // contribution
	)
}

// Helper function to create a person with specific ID
func personWithID(id uint) Person {
	person := *validPerson()
	person.ID = id
	return person
}

func parseDate(dateStr string) time.Time {
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		panic("Invalid date: " + err.Error())
	}
	return t
}

func TestFamily_Validate_Valid(t *testing.T) {
	f := Family{
		Name:    "Smith Family",
		Members: []Person{personWithID(1)},
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
		{"empty name", Family{Name: "", Members: []Person{personWithID(1)}}, "name is required"},
		{"short name", Family{Name: "AB", Members: []Person{personWithID(1)}}, "name must have at least 3 characters"},
		{"long name", Family{Name: string(make([]byte, 101)), Members: []Person{personWithID(1)}}, "name must have less than 100 characters"},
		{"duplicate member ID", Family{
			Name:    "ValidName",
			Members: []Person{personWithID(1), personWithID(1)},
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
	members := []Person{personWithID(1), personWithID(2)}

	attachment, err := NewAttachment("https://example.com/document.pdf")
	if err != nil {
		t.Fatalf("Failed to create test attachment: %v", err)
	}
	attachments := []Attachment{*attachment}

	description := "Family description text"

	family := NewFamily(name, members, attachments, description)

	if family.Name != name {
		t.Errorf("Expected name %q, got %q", name, family.Name)
	}

	if len(family.Members) != len(members) {
		t.Errorf("Expected %d members, got %d", len(members), len(family.Members))
	}

	memberIDs := map[uint]bool{}
	for _, m := range family.Members {
		memberIDs[m.ID] = true
	}

	if !memberIDs[1] || !memberIDs[2] {
		t.Error("Not all expected members were added correctly")
	}

	if len(family.Attachments) != 1 {
		t.Errorf("Expected 1 attachment, got %d", len(family.Attachments))
	}

	if family.Attachments[0].Path != attachment.Path {
		t.Errorf("Expected attachment path %q, got %q", attachment.Path, family.Attachments[0].Path)
	}

	if family.Description != description {
		t.Errorf("Expected description %q, got %q", description, family.Description)
	}
}

func TestFamily_AddMember(t *testing.T) {
	f := Family{Name: "Smith", Members: []Person{personWithID(1)}}

	// Creating a new person with different ID
	newPerson := personWithID(2)

	err := f.AddMember(newPerson)
	if err != nil {
		t.Fatalf("expected to add valid member, got error: %v", err)
	}
	if len(f.Members) != 2 {
		t.Fatalf("expected 2 members, got %d", len(f.Members))
	}

	found := false
	for _, member := range f.Members {
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
	f := Family{Name: "Smith", Members: []Person{personWithID(1)}}

	invalidPerson := NewPerson(
		"", // Empty name makes the person invalid
		"Doe",
		parseDate("1990-01-01"),
		"USA",
		"123456789",
		nil,
		"123-4567",
		"123 Street",
		100,
	)

	err := f.AddMember(*invalidPerson)
	if err == nil {
		t.Fatal("expected error for invalid member, got nil")
	}
	if !strings.Contains(err.Error(), "name is required") {
		t.Fatalf("expected error to mention invalid name, got: %v", err)
	}
}

func TestFamily_RemoveMemberByID(t *testing.T) {
	f := Family{
		Name: "Test Family",
		Members: []Person{
			personWithID(1),
			personWithID(2),
			personWithID(3),
		},
	}

	if len(f.Members) != 3 {
		t.Fatalf("setup failed: expected 3 members, got %d", len(f.Members))
	}

	removed := f.RemoveMemberByID(2)
	if !removed {
		t.Fatal("expected member to be removed, got false")
	}

	if len(f.Members) != 2 {
		t.Fatalf("expected 2 members after removal, got %d", len(f.Members))
	}

	for _, m := range f.Members {
		if m.ID == 2 {
			t.Fatal("member with ID 2 was not removed")
		}
	}

	foundIDs := map[uint]bool{}
	for _, m := range f.Members {
		foundIDs[m.ID] = true
	}
	if !foundIDs[1] || !foundIDs[3] {
		t.Fatal("expected members with IDs 1 and 3 to remain")
	}
}

func TestFamily_RemoveMemberByID_NotFound(t *testing.T) {
	f := Family{
		Name:    "Test",
		Members: []Person{personWithID(1), personWithID(2)},
	}

	removed := f.RemoveMemberByID(999) // Non-existent ID
	if removed {
		t.Fatal("expected removal to return false for nonexistent member")
	}

	if len(f.Members) != 2 {
		t.Fatalf("expected still 2 members, got %d", len(f.Members))
	}
}

func TestFamily_UpdateName(t *testing.T) {
	f := Family{Name: "OldName"}
	err := f.UpdateName("NewName")
	if err != nil {
		t.Fatalf("expected successful update, got error: %v", err)
	}
	if f.Name != "NewName" {
		t.Fatalf("expected name to be updated, got: %s", f.Name)
	}
}

func TestFamily_UpdateName_Invalid(t *testing.T) {
	f := Family{Name: "OldName"}

	err := f.UpdateName("AB")
	if err == nil {
		t.Fatal("expected error for short name, got nil")
	}
	if !strings.Contains(err.Error(), "must have at least 3 characters") {
		t.Fatalf("unexpected error message: %v", err)
	}

	if f.Name != "OldName" {
		t.Fatalf("name should not have been updated, got: %s", f.Name)
	}
}

func TestFamily_AddAttachment(t *testing.T) {
	f := Family{Name: "Souza"}

	a, err := NewAttachment("https://drive.google.com/file/d/abc123")
	if err != nil {
		t.Fatalf("expected valid attachment, got error: %v", err)
	}

	f.AddAttachment(*a)

	if len(f.Attachments) != 1 {
		t.Fatalf("expected 1 attachment, got %d", len(f.Attachments))
	}
}
