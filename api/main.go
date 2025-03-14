package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	dsn := os.Getenv("DB_URL")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		e.Logger.Fatal(err)
	}

	type People struct {
		ID    uint64 `json:"id" gorm:"primaryKey"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	// Auto Migrate the schema
	err = db.AutoMigrate(&People{})
	if err != nil {
		return
	}

	e.POST("/person", func(c echo.Context) error {
		person := new(People)
		if err := c.Bind(person); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}

		if result := db.Create(&person); result.Error != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": result.Error.Error(),
			})
		}

		return c.JSON(http.StatusCreated, person)
	})

	e.GET("/person", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":1323"))
	// add some comment
}
