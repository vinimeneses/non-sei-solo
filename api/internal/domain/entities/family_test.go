package entities

import (
	"testing"
)

func TestFamilyValidate(t *testing.T) {
	tests := []struct {
		name           string
		family         Family
		expectError    bool
		expectedErrMsg string
	}{
		{
			name: "Empty family name",
			family: Family{
				Name: "",
			},
			expectError:    true,
			expectedErrMsg: "name is required",
		},
		{
			name: "Valid family with no members",
			family: Family{
				Name: "Rossi Family",
			},
			expectError: false,
		},
		{
			name: "Valid family with valid members",
			family: Family{
				Name: "Rossi Family",
				members: []Person{
					{
						Name:        "Giovanni",
						Surname:     "Rossi",
						Birthday:    "1980-05-12",
						Citizenship: "Italian",
					},
					{
						Name:        "Maria",
						Surname:     "Rossi",
						Birthday:    "1985-03-22",
						Citizenship: "Italian",
					},
				},
			},
			expectError: false,
		},
		{
			name: "Family with an invalid member",
			family: Family{
				Name: "Rossi Family",
				members: []Person{
					{
						Name:        "Gi", // Invalid: less than 3 characters
						Surname:     "Rossi",
						Birthday:    "1980-05-12",
						Citizenship: "Italian",
					},
				},
			},
			expectError:    true,
			expectedErrMsg: "invalid member at index 0: name must be at least 3 characters long",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.family.validate()
			if tc.expectError {
				if err == nil {
					t.Errorf("expected error, but got none")
				} else if err.Error() != tc.expectedErrMsg {
					t.Errorf("expected error message '%s', but got '%s'", tc.expectedErrMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("did not expect error, but got: %s", err.Error())
				}
			}
		})
	}
}
