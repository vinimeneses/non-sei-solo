package entities

type ValidatedPerson struct {
	Person
	isValidated bool
}

func (vp *ValidatedPerson) IsValid() bool {
	return vp.isValidated
}

func NewValidatedPerson(p *Person) (*ValidatedPerson, error) {
	if err := p.validate(); err != nil {
		return nil, err
	}

	return &ValidatedPerson{
		Person:      *p,
		isValidated: true,
	}, nil
}
