package entities

import (
	"fmt"
	"net/url"
)

type Attachment struct {
	Path string
}

func (a *Attachment) Validate() error {
	if a.Path == "" {
		return fmt.Errorf("URL is required")
	}

	_, err := url.ParseRequestURI(a.Path)
	if err != nil {
		return fmt.Errorf("invalid URL: %v", err)
	}
	return nil
}

func NewAttachment(urlStr string) (*Attachment, error) {
	parsed, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return nil, fmt.Errorf("invalid URL: %v", err)
	}

	return &Attachment{
		Path: parsed.String(),
	}, nil
}
