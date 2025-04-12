package config

import (
	"api/internal/domain/repositories"
	postgres2 "api/internal/infrastructure/db/postgres"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Repositories struct {
	PersonRepository repositories.PersonRepository
	//FamilyRepository repositories.FamilyRepository
}

func Setup() *Repositories {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dsn := os.Getenv("DB_URL")

	db, err := postgres2.NewConnection(dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(
		&postgres2.Person{},
		&postgres2.Family{},
		&postgres2.Attachment{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	return &Repositories{
		PersonRepository: postgres2.NewGormPersonRepository(db),
		//FamilyRepository: postgres2.NewGormFamilyRepository(db),
	}
}
