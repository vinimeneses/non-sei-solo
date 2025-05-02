package repositories

import "api/internal/domain/entities"

type PersonRepository interface {
	Create(person *entities.ValidatedPerson) (*entities.Person, error)
	FindById(id uint) (*entities.Person, error)
	FindAll() ([]*entities.Person, error)
	Update(person *entities.ValidatedPerson) (*entities.Person, error)
	Delete(id uint) error
}
