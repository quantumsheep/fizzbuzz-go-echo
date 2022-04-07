package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/quantumsheep/fizzbuzz-rest/handlers/handlers_validator"
	services "github.com/quantumsheep/fizzbuzz-rest/services/cache"
	"github.com/stretchr/testify/assert"
)

func TestStatisticsMostUsed(t *testing.T) {
	mockedCache := services.NewMockedCache()
	fizzbuzzHandler := NewFizzbuzzHandler(mockedCache)
	statisticsHandler := NewStatisticsHandler(mockedCache)

	e := echo.New()
	e.Validator = handlers_validator.NewCustomValidator()

	// Execute one request one time
	for i := 0; i < 1; i++ {
		req := httptest.NewRequest(http.MethodGet, "/fizzbuzz?limit=5&int1=6&int2=7&str1=Hi&str2=Bob", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, fizzbuzzHandler.Fizzbuzz(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	}

	// Execute another request two times
	for i := 0; i < 2; i++ {
		req := httptest.NewRequest(http.MethodGet, "/fizzbuzz?limit=10&int1=3&int2=5&str1=Fizz&str2=Buzz", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, fizzbuzzHandler.Fizzbuzz(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	}

	// Fetch statistics
	req := httptest.NewRequest(http.MethodGet, "/statistics", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, statisticsHandler.Statistics(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "application/json; charset=UTF-8", rec.Header().Get(echo.HeaderContentType))
		assert.Equal(t, `{"hits":2,"parameters":{"int1":3,"int2":5,"limit":10,"str1":"Fizz","str2":"Buzz"}}`, strings.TrimSpace(rec.Body.String()))
	}
}
