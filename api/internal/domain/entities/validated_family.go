package entities

type ValidatedFamily struct {
	Family
	isValidated bool
}

func (vf *ValidatedFamily) IsValid() bool {
	return vf.isValidated
}

func NewValidatedFamily(family *Family) (*ValidatedFamily, error) {
	if err := family.validate(); err != nil {
		return nil, err
	}

	return &ValidatedFamily{
		Family:      *family,
		isValidated: true,
	}, nil
}
