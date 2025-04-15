package postgres

import (
	"api/internal/domain/entities"
	"api/internal/domain/repositories"
	"gorm.io/gorm"
)

type FamilyRepository struct {
	db *gorm.DB
}

func NewGormFamilyRepository(db *gorm.DB) repositories.FamilyRepository {
	return &FamilyRepository{db: db}
}

func (f FamilyRepository) Create(validatedFamily *entities.ValidatedFamily) (*entities.Family, error) {
	dbFamily := toDBFamily(validatedFamily)
	result := f.db.Create(dbFamily)
	if result.Error != nil {
		return nil, result.Error
	}

	domainFamily, err := toDomainFamily(dbFamily)
	if err != nil {
		return nil, err
	}
	return domainFamily, nil
}

func (f FamilyRepository) FindById(id uint) (*entities.Family, error) {
	var dbFamily Family
	result := f.db.Preload("Members").Preload("Attachments").First(&dbFamily, id)
	if result.Error != nil {
		return nil, result.Error
	}

	domainFamily, err := toDomainFamily(&dbFamily)
	if err != nil {
		return nil, err
	}
	return domainFamily, nil
}

func (f FamilyRepository) FindByPersonId(personId uint) (*entities.Family, error) {
	var dbPerson Person
	if err := f.db.First(&dbPerson, personId).Error; err != nil {
		return nil, err
	}

	var dbFamily Family
	if err := f.db.Preload("Members").Preload("Attachments").
		First(&dbFamily, dbPerson.FamilyID).Error; err != nil {
		return nil, err
	}

	domainFamily, err := toDomainFamily(&dbFamily)
	if err != nil {
		return nil, err
	}
	return domainFamily, nil
}

func (f FamilyRepository) FindAll() ([]*entities.Family, error) {
	var dbFamilies []*Family
	result := f.db.Preload("Members").Find(&dbFamilies)
	if result.Error != nil {
		return nil, result.Error
	}

	var families []*entities.Family
	for _, dbFamily := range dbFamilies {
		domainFamily, err := toDomainFamily(dbFamily)
		if err != nil {
			return nil, err
		}
		families = append(families, domainFamily)
	}
	return families, nil
}

func (f FamilyRepository) Update(validatedFamily *entities.ValidatedFamily) (*entities.Family, error) {
	dbFamily := toDBFamily(validatedFamily)

	err := f.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(dbFamily).Error; err != nil {
			return err
		}

		if err := tx.Model(dbFamily).Association("Members").Replace(dbFamily.Members); err != nil {
			return err
		}

		if err := tx.Model(dbFamily).Association("Attachments").Replace(dbFamily.Attachments); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	domainFamily, err := toDomainFamily(dbFamily)
	if err != nil {
		return nil, err
	}
	return domainFamily, nil
}

func (f FamilyRepository) Delete(id uint) error {
	return f.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&Family{}, id).Error; err != nil {
			return err
		}
		return nil
	})
}
