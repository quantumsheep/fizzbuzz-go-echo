package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	services "github.com/quantumsheep/fizzbuzz-rest/services/cache"
)

type StatisticsHandler struct {
	Cache services.Cache
}

func NewStatisticsHandler(cache services.Cache) *StatisticsHandler {
	return &StatisticsHandler{
		Cache: cache,
	}
}

func (h *StatisticsHandler) Statistics(c echo.Context) (err error) {
	key, hits, err := getBestFizzbuzz(h.Cache)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	parameters := make(map[string]interface{})
	json.Unmarshal([]byte(key), &parameters)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"parameters": parameters,
		"hits":       hits,
	})
}
