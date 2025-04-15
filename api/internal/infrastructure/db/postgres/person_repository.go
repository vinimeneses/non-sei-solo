package postgres

import (
	"api/internal/domain/entities"
	"api/internal/domain/repositories"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type GormPersonRepository struct {
	db *gorm.DB
}

func NewGormPersonRepository(db *gorm.DB) repositories.PersonRepository {
	return &GormPersonRepository{db: db}
}

func (g *GormPersonRepository) Create(person *entities.ValidatedPerson) (*entities.Person, error) {

	dbPerson := toDBPerson(person)
	if err := g.db.Create(dbPerson).Error; err != nil {
		return nil, err
	}

	return toDomainPerson(dbPerson), nil
}

func (g *GormPersonRepository) FindById(id uint) (*entities.Person, error) {
	var dbPerson Person

	if err := g.db.First(&dbPerson, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("person with ID %d not found", id)
		}
		return nil, err
	}

	return toDomainPerson(&dbPerson), nil
}

func (g *GormPersonRepository) FindAll() ([]*entities.Person, error) {
	var dbPersons []Person

	if err := g.db.Find(&dbPersons).Error; err != nil {
		return nil, err
	}

	result := make([]*entities.Person, len(dbPersons))
	for i, p := range dbPersons {
		person := p
		result[i] = toDomainPerson(&person)
	}

	return result, nil
}

func (g *GormPersonRepository) Update(vp *entities.ValidatedPerson) (*entities.Person, error) {
	dbPerson := toDBPerson(vp)

	if err := g.db.
		Model(&dbPerson).
		Select("*").
		Where("id = ?", dbPerson.ID).
		Updates(dbPerson).
		Error; err != nil {
		return nil, err
	}

	if err := g.db.
		Where("id = ?", dbPerson.ID).
		First(&dbPerson).
		Error; err != nil {
		return nil, err
	}

	return toDomainPerson(dbPerson), nil
}

func (g *GormPersonRepository) Delete(id uint) error {
	result := g.db.Delete(&Person{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("person with ID %d not found", id)
	}

	return nil
}
