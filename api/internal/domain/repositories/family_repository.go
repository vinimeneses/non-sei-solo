package repositories

import "api/internal/domain/entities"

type FamilyRepository interface {
	Create(family *entities.Family) (*entities.Family, error)
	FindById(id string) (*entities.Family, error)
	FindByPersonId(personId string) ([]*entities.Family, error)
	FindAll() ([]*entities.Family, error)
	Update(family *entities.Family) (*entities.Family, error)
	Delete(id string) error
}
