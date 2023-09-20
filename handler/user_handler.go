package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandleSignIn(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"user":  "DucDO",
		"email": "ducdo@gmail.com",
	})
}
func HandleSignUp(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
