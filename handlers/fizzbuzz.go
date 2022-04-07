package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	services "github.com/quantumsheep/fizzbuzz-rest/services/cache"
)

type FizzbuzzHandler struct {
	Cache services.Cache
}

func NewFizzbuzzHandler(cache services.Cache) *FizzbuzzHandler {
	return &FizzbuzzHandler{
		Cache: cache,
	}
}

type GetFizzbuzzDTO struct {
	Limit int    `param:"limit" query:"limit" form:"limit" json:"limit" xml:"limit" validate:"required"`
	Int1  int    `param:"int1" query:"int1" form:"int1" json:"int1" xml:"int1" validate:"required"`
	Int2  int    `param:"int2" query:"int2" form:"int2" json:"int2" xml:"int2" validate:"required"`
	Str1  string `param:"str1" query:"str1" form:"str1" json:"str1" xml:"str1" validate:"required"`
	Str2  string `param:"str2" query:"str2" form:"str2" json:"str2" xml:"str2" validate:"required"`
}

// @Produce      json
// @Param        limit			query		int  		true  		"generate fizzbuzz from 1 to this number (inclusive)"
// @Param        int1			query   	int  		true  		"str2 will be append if n is divisible by this number"
// @Param        str1			query   	string  	true  		"value to be append if n is divisible by int1"
// @Param        int2			query   	int  		true  		"str2 will be append if n is divisible by this number"
// @Param        str2			query   	string  	true  		"value to be append if n is divisible by int2"
// @Success      200			{object}  	[]string
// @Failure      400
// @Failure      500
// @Router       /fizzbuzz [get]
func (h *FizzbuzzHandler) Fizzbuzz(c echo.Context) (err error) {
	dto := new(GetFizzbuzzDTO)
	if err := c.Bind(dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(dto); err != nil {
		return err
	}

	strings := make([]string, dto.Limit)

	for i := 1; i <= dto.Limit; i++ {
		s := ""

		if i%dto.Int1 == 0 {
			s += dto.Str1
		}

		if i%dto.Int2 == 0 {
			s += dto.Str2
		}

		if s == "" {
			s = strconv.Itoa(i)
		}

		strings[i-1] = s
	}

	cacheKey, err := h.getFizzbuzzCacheKey(dto)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	count := 0

	if cachedCount, err := h.Cache.Get(string(cacheKey)); err == nil {
		count, err = strconv.Atoi(cachedCount)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	count += 1
	err = h.Cache.Set(string(cacheKey), strconv.Itoa(count))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if _, hits, err := getBestFizzbuzz(h.Cache); err != nil || count+1 > hits {
		err = h.Cache.Set("fizzbuzz-best", string(cacheKey))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(200, strings)
}

func (h *FizzbuzzHandler) getFizzbuzzCacheKey(dto *GetFizzbuzzDTO) (string, error) {
	bytes, err := json.Marshal(dto)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func getBestFizzbuzz(cache services.Cache) (key string, hits int, err error) {
	key, err = cache.Get("fizzbuzz-best")
	if err != nil {
		return "", 0, fmt.Errorf("There is no best fizzbuzz yet")
	}

	hitsString, err := cache.Get(key)
	if err != nil {
		return "", 0, err
	}

	hits, err = strconv.Atoi(hitsString)
	if err != nil {
		return "", 0, err
	}

	return key, hits, nil
}
