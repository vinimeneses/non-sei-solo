package entities

import (
	"fmt"
	"net/url"
)

type Attachment struct {
	Path string `json:"path" validate:"required"`
}

func (a *Attachment) validate() error {
	if a.Path == "" {
		return fmt.Errorf("path is required")
	}

	_, err := url.ParseRequestURI(a.Path)
	if err != nil {
		return fmt.Errorf("invalid URL: %v", err)
	}
	return nil
}
