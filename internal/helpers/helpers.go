package helpers

import (
	"strconv"

	"github.com/labstack/echo"
)

func URLParamInt(c echo.Context, key string) (int, error) {
	return strconv.Atoi(c.Param(key))
}
