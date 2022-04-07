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

func TestFizzBuzz1(t *testing.T) {
	mockedCache := services.NewMockedCache()
	fizzbuzzHandler := NewFizzbuzzHandler(mockedCache)

	e := echo.New()
	e.Validator = handlers_validator.NewCustomValidator()

	req := httptest.NewRequest(http.MethodGet, "/fizzbuzz?limit=10&int1=3&int2=5&str1=Fizz&str2=Buzz", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, fizzbuzzHandler.Fizzbuzz(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "application/json; charset=UTF-8", rec.Header().Get(echo.HeaderContentType))
		assert.Equal(t, `["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz"]`, strings.TrimSpace(rec.Body.String()))
	}
}
