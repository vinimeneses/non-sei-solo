package repositories

import "api/internal/domain/entities"

type FamilyRepository interface {
	Create(family *entities.ValidatedFamily) (*entities.Family, error)
	FindById(id uint) (*entities.Family, error)
	FindByPersonId(personId uint) (*entities.Family, error)
	FindAll() ([]*entities.Family, error)
	Update(family *entities.ValidatedFamily) (*entities.Family, error)
	Delete(id uint) error
}
