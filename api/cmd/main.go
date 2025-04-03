package main

import (
	"api/internal/config"
	"api/internal/domain/entities"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Validator = &config.CustomValidator{Validator: validator.New()}
	e.Use(middleware.Logger())
	err := godotenv.Load(".env")
	if err != nil {
		e.Logger.Fatal(err)
	}
	dsn := os.Getenv("DB_URL")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		e.Logger.Fatal(err)
	}

	err = db.AutoMigrate(&entities.Family{}, &entities.Person{}, &entities.Attachment{})
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.Logger.Fatal(e.Start(":1323"))

}
