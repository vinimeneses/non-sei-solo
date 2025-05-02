package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection(dns string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dns), &gorm.Config{})
}
