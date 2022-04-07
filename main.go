package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/quantumsheep/fizzbuzz-rest/docs"
	"github.com/quantumsheep/fizzbuzz-rest/handlers"
	"github.com/quantumsheep/fizzbuzz-rest/handlers/handlers_validator"
	services "github.com/quantumsheep/fizzbuzz-rest/services/cache"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Fizzbuzz REST API
// @version 1.0
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()
	e.Validator = handlers_validator.NewCustomValidator()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	cache := services.NewRedisCache(os.Getenv("REDIS_HOST"))

	fizzbuzzHandler := handlers.NewFizzbuzzHandler(cache)
	e.GET("/fizzbuzz", fizzbuzzHandler.Fizzbuzz)

	statisticsHandler := handlers.NewStatisticsHandler(cache)
	e.GET("/statistics", statisticsHandler.Statistics)

	e.Logger.Fatal(e.Start(":1323"))
}
