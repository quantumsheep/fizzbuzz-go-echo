package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type getFizzbuzzDTO struct {
	Limit int    `param:"limit" query:"limit" form:"limit" json:"limit" xml:"limit" validate:"required"`
	Int1  int    `param:"int1" query:"int1" form:"int1" json:"int1" xml:"int1" validate:"required"`
	Int2  int    `param:"int2" query:"int2" form:"int2" json:"int2" xml:"int2" validate:"required"`
	Str1  string `param:"str1" query:"str1" form:"str1" json:"str1" xml:"str1" validate:"required"`
	Str2  string `param:"str2" query:"str2" form:"str2" json:"str2" xml:"str2" validate:"required"`
}

func Fizzbuzz(c echo.Context) (err error) {
	u := new(getFizzbuzzDTO)
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(u); err != nil {
		return err
	}

	strings := make([]string, u.Limit)

	for i := 1; i <= u.Limit; i++ {
		s := ""

		if i%u.Int1 == 0 {
			s += u.Str1
		}

		if i%u.Int2 == 0 {
			s += u.Str2
		}

		if s == "" {
			s = strconv.Itoa(i)
		}

		strings[i-1] = s
	}

	return c.JSON(200, strings)
}
