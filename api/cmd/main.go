package main

import (
	"api/internal/config"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Validator = &config.CustomValidator{Validator: validator.New()}
	e.Use(middleware.Logger())
	config.Setup()
	e.Logger.Fatal(e.Start(":1323"))
}
