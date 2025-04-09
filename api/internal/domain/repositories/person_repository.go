package repositories

import "api/internal/domain/entities"

type PersonRepository interface {
	Create(person *entities.Person) (*entities.Person, error)
	FindById(id string) (*entities.Person, error)
	FindAll() ([]*entities.Person, error)
	Update(person *entities.Person) (*entities.Person, error)
	Delete(id string) error
}
