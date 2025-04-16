package sqlite_test

import (
	"api/internal/domain/entities"
	"api/internal/infrastructure/db/postgres"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
	"time"

	_ "github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTestDB(t *testing.T) (*gorm.DB, func()) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	require.NoError(t, err)

	err = db.AutoMigrate(&postgres.Family{}, &postgres.Person{}, &postgres.Attachment{})
	require.NoError(t, err)

	return db, func() {
		sqlDB, err := db.DB()
		if err == nil {
			_ = sqlDB.Close()
		}
	}
}

func createSampleValidatedFamily(t *testing.T) *entities.ValidatedFamily {
	birthday := time.Now()

	person1 := entities.NewPerson(
		"Carlos", "Silva", birthday, "Brazil", "TAX123", nil,
		"999-111-1111", "Street A", 100.0,
	)
	validatedPerson1, err := entities.NewValidatedPerson(person1)
	require.NoError(t, err)

	person2 := entities.NewPerson(
		"Ana", "Silva", birthday, "Brazil", "TAX456", nil,
		"999-222-2222", "Street B", 200.0,
	)
	validatedPerson2, err := entities.NewValidatedPerson(person2)
	require.NoError(t, err)

	attachment := &entities.Attachment{
		Path: "https://drive.google.com/file/d/abc123",
	}
	validatedAttachment, err := entities.NewValidatedAttachment(attachment)
	require.NoError(t, err)

	family, err := entities.NewFamily(
		"Silva",
		[]*entities.ValidatedPerson{validatedPerson1, validatedPerson2},
		[]*entities.ValidatedAttachment{validatedAttachment},
		"",
	)
	require.NoError(t, err)

	vf, err := entities.NewValidatedFamily(family)

	if err != nil {
		t.Fatalf("failed to create validated family: %v", err)
	}

	return vf
}

func TestGormFamilyRepository_Create(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	repo := postgres.NewGormFamilyRepository(db)
	validatedFamily := createSampleValidatedFamily(t)

	created, err := repo.Create(validatedFamily)
	require.NoError(t, err)
	require.NotNil(t, created)

	assert.Equal(t, "Silva", created.Name)
	assert.Len(t, created.Members, 2)
	assert.Len(t, created.Attachments, 1)
}

func TestGormFamilyRepository_FindById(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	repo := postgres.NewGormFamilyRepository(db)
	validatedFamily := createSampleValidatedFamily(t)

	created, err := repo.Create(validatedFamily)
	require.NoError(t, err)

	found, err := repo.FindById(created.ID)
	require.NoError(t, err)
	assert.Equal(t, created.ID, found.ID)
}

func TestGormFamilyRepository_FindByPersonId(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	repo := postgres.NewGormFamilyRepository(db)
	validatedFamily := createSampleValidatedFamily(t)

	created, err := repo.Create(validatedFamily)
	require.NoError(t, err)

	personID := created.Members[0].ID
	found, err := repo.FindByPersonId(personID)
	require.NoError(t, err)
	assert.Equal(t, created.ID, found.ID)
}

func TestGormFamilyRepository_FindAll(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	repo := postgres.NewGormFamilyRepository(db)

	families, err := repo.FindAll()
	require.NoError(t, err)
	assert.Empty(t, families)

	_, err = repo.Create(createSampleValidatedFamily(t))
	require.NoError(t, err)
	_, err = repo.Create(createSampleValidatedFamily(t))
	require.NoError(t, err)

	families, err = repo.FindAll()
	require.NoError(t, err)
	assert.Len(t, families, 2)
}

func TestGormFamilyRepository_Update(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	repo := postgres.NewGormFamilyRepository(db)
	original := createSampleValidatedFamily(t)

	created, err := repo.Create(original)
	require.NoError(t, err)

	var validatedMembers []*entities.ValidatedPerson
	for _, p := range created.Members {
		pCopy := p
		vp, err := entities.NewValidatedPerson(&pCopy)
		require.NoError(t, err)
		validatedMembers = append(validatedMembers, vp)
	}

	var validatedAttachments []*entities.ValidatedAttachment
	for _, a := range created.Attachments {
		aCopy := a
		va, err := entities.NewValidatedAttachment(&aCopy)
		require.NoError(t, err)
		validatedAttachments = append(validatedAttachments, va)
	}

	updated, err := entities.NewFamily(
		"Oliveira",
		validatedMembers,
		validatedAttachments,
		"",
	)
	require.NoError(t, err)

	validatedUpdated, err := entities.NewValidatedFamily(updated)
	require.NoError(t, err)
	validatedUpdated.ID = created.ID

	result, err := repo.Update(validatedUpdated)
	require.NoError(t, err)
	assert.Equal(t, "Oliveira", result.Name)
}

func TestGormFamilyRepository_Delete(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	repo := postgres.NewGormFamilyRepository(db)
	validatedFamily := createSampleValidatedFamily(t)

	created, err := repo.Create(validatedFamily)
	require.NoError(t, err)

	err = repo.Delete(created.ID)
	require.NoError(t, err)

	_, err = repo.FindById(created.ID)
	assert.Error(t, err)
}
