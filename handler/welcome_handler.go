package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Welcome(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
