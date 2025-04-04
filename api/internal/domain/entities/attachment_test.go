package entities

import (
	"strings"
	"testing"
)

func TestAttachmentValidate(t *testing.T) {
	tests := []struct {
		name           string
		attachment     Attachment
		expectError    bool
		expectedErrMsg string
	}{
		{
			name: "Empty path",
			attachment: Attachment{
				Path: "",
			},
			expectError:    true,
			expectedErrMsg: "path is required",
		},
		{
			name: "Invalid URL",
			attachment: Attachment{
				Path: "invalid-url",
			},
			expectError:    true,
			expectedErrMsg: "invalid URL",
		},
		{
			name: "Valid URL",
			attachment: Attachment{
				Path: "https://www.google.com/images/branding/googlelogo/1x/googlelogo_color_272x92dp.png",
			},
			expectError: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.attachment.validate()
			if tc.expectError {
				if err == nil {
					t.Errorf("expected error, but got none")
				} else if !strings.HasPrefix(err.Error(), tc.expectedErrMsg) {
					t.Errorf("expected error message to start with '%s', but got '%s'", tc.expectedErrMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("did not expect error, but got: %s", err.Error())
				}
			}
		})
	}
}
