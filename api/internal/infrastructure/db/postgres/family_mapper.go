package postgres

import (
	"api/internal/domain/entities"
	"gorm.io/gorm"
)

func toDBFamily(family *entities.ValidatedFamily) *Family {
	dbMembers := make([]Person, len(family.Members))
	for i, m := range family.Members {
		dbMembers[i] = Person{
			Model:        gorm.Model{ID: m.ID},
			Name:         m.Name,
			Surname:      m.Surname,
			Birthday:     m.Birthday,
			Citizenship:  m.Citizenship,
			TaxCode:      m.TaxCode,
			FamilyID:     m.FamilyID,
			PhoneNumber:  m.PhoneNumber,
			Address:      m.Address,
			Contribution: m.Contribution,
		}
	}

	dbAttachments := make([]Attachment, len(family.Attachments))
	for i, a := range family.Attachments {
		dbAttachments[i] = Attachment{
			Path: a.Path,
		}
	}

	return &Family{
		Model:       gorm.Model{},
		Members:     dbMembers,
		Name:        family.Name,
		Attachments: dbAttachments,
		Description: family.Description,
	}
}

func toDomainFamily(db *Family) (*entities.Family, error) {
	domainMembers := make([]entities.Person, len(db.Members))
	for i, dbm := range db.Members {
		domainMembers[i] = entities.Person{
			ID:           dbm.ID,
			Name:         dbm.Name,
			Surname:      dbm.Surname,
			Birthday:     dbm.Birthday,
			Citizenship:  dbm.Citizenship,
			TaxCode:      dbm.TaxCode,
			FamilyID:     dbm.FamilyID,
			PhoneNumber:  dbm.PhoneNumber,
			Address:      dbm.Address,
			Contribution: dbm.Contribution,
		}
	}

	domainAttachments := make([]entities.Attachment, len(db.Attachments))
	for i, a := range db.Attachments {
		domainAttachments[i] = entities.Attachment{
			Path: a.Path,
		}
	}

	domain := &entities.Family{
		ID:          db.ID,
		Members:     domainMembers,
		Name:        db.Name,
		Attachments: domainAttachments,
		Description: db.Description,
	}

	return domain, nil
}
