package sqlite_test

import (
	"api/internal/domain/entities"
	"api/internal/infrastructure/db/postgres"
	"gorm.io/gorm/logger"
	"testing"
	"time"

	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) (*gorm.DB, func()) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	require.NoError(t, err)

	err = db.AutoMigrate(&postgres.Person{})
	require.NoError(t, err)

	return db, func() {
		sqlDB, err := db.DB()
		if err == nil {
			_ = sqlDB.Close()
		}
	}
}

func createValidatedPerson(name, surname, citizenship string, birthday time.Time, ids ...uint) *entities.ValidatedPerson {
	person := entities.NewPerson(
		name,
		surname,
		birthday,
		citizenship,
		"TEST123",
		nil,
		"123-456-7890",
		"123 Test St",
		100.0,
	)

	if len(ids) > 0 {
		person.ID = ids[0]
	}

	validatedPerson, err := entities.NewValidatedPerson(person)
	if err != nil {
		panic(err)
	}
	return validatedPerson
}

func parseDate(dateStr string) time.Time {
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		panic("Invalid date: " + err.Error())
	}
	return t
}

func TestGormPersonRepository_Create(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	repo := postgres.NewGormPersonRepository(db)

	birthday := parseDate("1990-01-01")
	validatedPerson := createValidatedPerson("John", "Doe", "USA", birthday)

	createdPerson, err := repo.Create(validatedPerson)
	require.NoError(t, err)
	require.NotZero(t, createdPerson.ID)

	assert.Equal(t, "John", createdPerson.Name)
	assert.Equal(t, "Doe", createdPerson.Surname)
	assert.Equal(t, birthday.Format("2006-01-02"), createdPerson.Birthday.Format("2006-01-02"))
	assert.Equal(t, "USA", createdPerson.Citizenship)
	assert.Equal(t, "TEST123", createdPerson.TaxCode)
	assert.Equal(t, "123-456-7890", createdPerson.PhoneNumber)
	assert.Equal(t, "123 Test St", createdPerson.Address)
	assert.Equal(t, 100.0, createdPerson.Contribution)
}

func TestGormPersonRepository_FindById(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	repo := postgres.NewGormPersonRepository(db)

	birthday := parseDate("1985-05-15")
	validatedPerson := createValidatedPerson("Jane", "Smith", "Canada", birthday)
	createdPerson, err := repo.Create(validatedPerson)
	require.NoError(t, err)
	personID := createdPerson.ID

	t.Run("Existing person", func(t *testing.T) {
		foundPerson, err := repo.FindById(personID)
		require.NoError(t, err)
		require.NotNil(t, foundPerson)

		assert.Equal(t, personID, foundPerson.ID)
		assert.Equal(t, "Jane", foundPerson.Name)
		assert.Equal(t, birthday.Format("2006-01-02"), foundPerson.Birthday.Format("2006-01-02"))
	})

	t.Run("Non-existent person", func(t *testing.T) {
		nonExistentID := personID + 999
		foundPerson, err := repo.FindById(nonExistentID)

		assert.Error(t, err)
		assert.Nil(t, foundPerson)
		assert.Contains(t, err.Error(), "not found")
	})
}

func TestGormPersonRepository_FindAll(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	repo := postgres.NewGormPersonRepository(db)

	t.Run("Empty database", func(t *testing.T) {
		persons, err := repo.FindAll()
		assert.NoError(t, err)
		assert.Empty(t, persons)
	})

	b1 := parseDate("1990-01-01")
	b2 := parseDate("1985-05-15")
	v1 := createValidatedPerson("John", "Doe", "USA", b1)
	v2 := createValidatedPerson("Jane", "Smith", "Canada", b2)

	_, err := repo.Create(v1)
	require.NoError(t, err)
	_, err = repo.Create(v2)
	require.NoError(t, err)

	t.Run("Multiple persons", func(t *testing.T) {
		persons, err := repo.FindAll()
		assert.NoError(t, err)
		assert.Len(t, persons, 2)

		names := []string{persons[0].Name, persons[1].Name}
		assert.Contains(t, names, "John")
		assert.Contains(t, names, "Jane")
	})
}

func TestGormPersonRepository_Update(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	repo := postgres.NewGormPersonRepository(db)

	birthday := parseDate("1990-01-01")
	vp := createValidatedPerson("John", "Doe", "USA", birthday)
	created, err := repo.Create(vp)
	require.NoError(t, err)
	personID := created.ID

	updatedBirthday := parseDate("1991-02-02")
	p := entities.NewPerson(
		"Johnny",
		"Doeson",
		updatedBirthday,
		"Canada",
		"XYZ789",
		nil,
		"987-654-3210",
		"456 Other St",
		200.0,
	)
	p.ID = personID
	vpToUpdate, err := entities.NewValidatedPerson(p)
	require.NoError(t, err)

	updatedPerson, err := repo.Update(vpToUpdate)
	require.NoError(t, err)
	require.NotNil(t, updatedPerson)

	assert.Equal(t, personID, updatedPerson.ID)
	assert.Equal(t, "Johnny", updatedPerson.Name)
	assert.Equal(t, "Doeson", updatedPerson.Surname)
	assert.Equal(t,
		updatedBirthday.Format("2006-01-02"),
		updatedPerson.Birthday.Format("2006-01-02"),
	)
	assert.Equal(t, "Canada", updatedPerson.Citizenship)
	assert.Equal(t, "XYZ789", updatedPerson.TaxCode)
	assert.Equal(t, "987-654-3210", updatedPerson.PhoneNumber)
	assert.Equal(t, "456 Other St", updatedPerson.Address)
	assert.Equal(t, 200.0, updatedPerson.Contribution)

	retrieved, err := repo.FindById(personID)
	require.NoError(t, err)
	assert.Equal(t, "Johnny", retrieved.Name)
	assert.Equal(t, "Doeson", retrieved.Surname)
	assert.Equal(t,
		updatedBirthday.Format("2006-01-02"),
		retrieved.Birthday.Format("2006-01-02"),
	)
	assert.Equal(t, "Canada", retrieved.Citizenship)
	assert.Equal(t, "XYZ789", retrieved.TaxCode)
	assert.Equal(t, "987-654-3210", retrieved.PhoneNumber)
	assert.Equal(t, "456 Other St", retrieved.Address)
	assert.Equal(t, 200.0, retrieved.Contribution)
}

func TestGormPersonRepository_Delete(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	repo := postgres.NewGormPersonRepository(db)

	birthday := parseDate("1990-01-01")
	vp := createValidatedPerson("John", "Doe", "USA", birthday)
	created, err := repo.Create(vp)
	require.NoError(t, err)
	personID := created.ID

	t.Run("Delete existing person", func(t *testing.T) {
		err := repo.Delete(personID)
		assert.NoError(t, err)

		_, err = repo.FindById(personID)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "not found")
	})

	t.Run("Delete non-existent person", func(t *testing.T) {
		nonExistentID := personID + 999
		err := repo.Delete(nonExistentID)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "not found")
	})
}
