package entities

import (
	"testing"
)

func TestPersonValidate(t *testing.T) {
	tests := []struct {
		name           string
		person         Person
		expectError    bool
		expectedErrMsg string
	}{
		{
			name: "Empty name",
			person: Person{
				Name:        "",
				Surname:     "Rossi",
				Birthday:    "1980-05-12",
				Citizenship: "Italian",
			},
			expectError:    true,
			expectedErrMsg: "name is required",
		},
		{
			name: "Name with less than 3 characters",
			person: Person{
				Name:        "Gi",
				Surname:     "Rossi",
				Birthday:    "1980-05-12",
				Citizenship: "Italian",
			},
			expectError:    true,
			expectedErrMsg: "name must be at least 3 characters long",
		},
		{
			name: "Valid name",
			person: Person{
				Name:        "Giovanni",
				Surname:     "Rossi",
				Birthday:    "1980-05-12",
				Citizenship: "Italian",
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.person.validate()
			if tt.expectError && err == nil {
				t.Errorf("expected error, but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("did not expect error, but got: %s", err.Error())
			}
			if tt.expectError && err != nil && err.Error() != tt.expectedErrMsg {
				t.Errorf("expected error message '%s', but got '%s'", tt.expectedErrMsg, err.Error())
			}
		})
	}
}
