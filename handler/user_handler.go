package handler

import (
	"net/http"

	"github.com/ddduc02/gh-trending/models"
	"github.com/ddduc02/gh-trending/models/request"
	"github.com/ddduc02/gh-trending/repository"
	"github.com/ddduc02/gh-trending/security"
	validator "github.com/go-playground/validator/v10"
	uuid "github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type UserHandler struct {
	UserRepo repository.UserRepo
}

var validate *validator.Validate

func (u *UserHandler) HandleSignUp(c echo.Context) error {
	req := request.RequestSignUp{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	var role string
	if req.Email == "ducdo@gmail.com" {
		role = "ADMIN"
	} else {
		role = "MEMBER"
	}
	hash := security.HashPassword(req.Password)
	userId, err := uuid.NewUUID()
	if err != nil {
		return c.JSON(http.StatusForbidden, models.Response{
			StatusCode: http.StatusForbidden,
			Message:    err.Error(),
		})
	}
	user := models.User{
		UserId:   userId.String(),
		FullName: req.FullName,
		Email:    req.Email,
		Password: hash,
		Role:     role,
		Token:    "",
	}
	user, err = u.UserRepo.SaveUser(c.Request().Context(), user)
	if err != nil {
		return c.JSON(http.StatusConflict, models.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		Message:    "Sign up successfully",
		Data:       user,
	})
}
func (u *UserHandler) HandleSignIn(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"user":  "DucDO",
		"email": "ducdo@gmail.com",
	})
}
