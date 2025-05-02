package entities

type ValidatedAttachment struct {
	Attachment
	isValidated bool
}

func (va *ValidatedAttachment) IsValid() bool {
	return va.isValidated
}

func NewValidatedAttachment(attachment *Attachment) (*ValidatedAttachment, error) {
	if err := attachment.Validate(); err != nil {
		return nil, err
	}

	return &ValidatedAttachment{
		Attachment:  *attachment,
		isValidated: true,
	}, nil
}
