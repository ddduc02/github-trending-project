package middleware

import (
	"github.com/ddduc02/gh-trending/models"
	"github.com/ddduc02/gh-trending/security"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims:     &models.JwtCustomClaims{},
		SigningKey: []byte(security.SECRET_KEY),
	}

	return middleware.JWTWithConfig(config)
}
