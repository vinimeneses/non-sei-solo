package postgres

import (
	"api/internal/domain/entities"
	"api/internal/domain/repositories"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strconv"
)

type GormPersonRepository struct {
	db *gorm.DB
}

func NewGormPersonRepository(db *gorm.DB) repositories.PersonRepository {
	return &GormPersonRepository{db: db}
}

func (g *GormPersonRepository) Create(person *entities.Person) (*entities.Person, error) {
	validatedPerson, err := entities.NewValidatedPerson(person)
	if err != nil {
		return nil, err
	}

	dbPerson := toDBPerson(validatedPerson)
	if err := g.db.Create(dbPerson).Error; err != nil {
		return nil, err
	}

	return toDomainPerson(dbPerson), nil
}

func (g *GormPersonRepository) FindById(id string) (*entities.Person, error) {
	var dbPerson Person

	uintID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %v", err)
	}

	if err := g.db.First(&dbPerson, uint(uintID)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("person with ID %s not found", id)
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

func (g *GormPersonRepository) Update(person *entities.Person) (*entities.Person, error) {
	validatedPerson, err := entities.NewValidatedPerson(person)
	if err != nil {
		return nil, err
	}

	dbPerson := toDBPerson(validatedPerson)
	if err := g.db.Save(dbPerson).Error; err != nil {
		return nil, err
	}

	return toDomainPerson(dbPerson), nil
}

func (g *GormPersonRepository) Delete(id string) error {
	uintID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return fmt.Errorf("invalid ID format: %v", err)
	}

	result := g.db.Delete(&Person{}, uint(uintID))
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("person with ID %s not found", id)
	}

	return nil
}
