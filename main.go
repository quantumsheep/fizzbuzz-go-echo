package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/quantumsheep/leboncoin-fizzbuzz/handlers"
	"github.com/quantumsheep/leboncoin-fizzbuzz/handlers/handlers_validator"
)

func main() {
	e := echo.New()
	e.Validator = handlers_validator.NewCustomValidator()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/fizzbuzz", handlers.Fizzbuzz)

	e.Logger.Fatal(e.Start(":1323"))
}
