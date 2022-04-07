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

type GetStatisticsResponseDTO struct {
	Hits       int            `json:"hits"`
	Parameters GetFizzbuzzDTO `json:"parameters"`
}

// @Produce      json
// @Success      200			{object}  	GetStatisticsResponseDTO
// @Failure      500
// @Router       /statistics [get]
func (h *StatisticsHandler) Statistics(c echo.Context) (err error) {
	key, hits, err := getBestFizzbuzz(h.Cache)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var parameters GetFizzbuzzDTO
	err = json.Unmarshal([]byte(key), &parameters)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, GetStatisticsResponseDTO{
		Hits:       hits,
		Parameters: parameters,
	})
}
