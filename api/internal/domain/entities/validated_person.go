package entities

type ValidatedPerson struct {
	Person
	IsValidated bool
}

func (vp *ValidatedPerson) IsValid() bool {
	return vp.IsValidated
}

func NewValidatedPerson(p *Person) (*ValidatedPerson, error) {
	if err := p.validate(); err != nil {
		return nil, err
	}

	return &ValidatedPerson{
		Person:      *p,
		IsValidated: true,
	}, nil
}
