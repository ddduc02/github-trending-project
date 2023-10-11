package handler

import (
	"net/http"

	"github.com/ddduc02/gh-trending/models"
	"github.com/ddduc02/gh-trending/models/request"
	"github.com/ddduc02/gh-trending/repository"
	"github.com/ddduc02/gh-trending/security"
	validator "github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	uuid "github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type UserHandler struct {
	UserRepo repository.UserRepo
}

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
	token, err := security.GenerateToken(user)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, models.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	user.Token = token
	return c.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		Message:    "Sign up successfully",
		Data:       user,
	})
}
func (u *UserHandler) HandleSignIn(c echo.Context) error {
	req := request.RequestSignIp{}
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
	user, err := u.UserRepo.CheckLogin(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, models.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	check := security.CompareHashAndPassword(req.Password, user.Password)
	if !check {
		return c.JSON(http.StatusUnauthorized, models.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    "Login failed",
			Data:       nil,
		})
	}
	token, err := security.GenerateToken(user)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, models.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	user.Token = token
	// set password  = "" trước khi trả về cho người dùng
	user.Password = ""
	return c.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		Message:    "Login successfully",
		Data:       user,
	})
}

func (u *UserHandler) Profile(c echo.Context) error {
	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*models.JwtCustomClaims)
	user, err := u.UserRepo.SelectUserById(c.Request().Context(), claims.UserId)
	if err != nil {
		return c.JSON(http.StatusNotFound, models.Response{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	return c.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		Message:    "Done",
		Data:       user,
	})
}

func (u *UserHandler) UpdateProfile(c echo.Context) error {
	req := request.RequestUpdate{}
	if err := c.Bind(&req); err != nil {
		return err
	}

	// err := c.Validate(req)
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, models.Response{
	// 		StatusCode: http.StatusBadRequest,
	// 		Message:    "loi " + err.Error(),
	// 	})
	// }
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*models.JwtCustomClaims)
	user := models.User{
		UserId:   claims.UserId,
		FullName: req.FullName,
		Email:    req.Email,
	}

	user, err := u.UserRepo.UpdateUser(c.Request().Context(), user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, models.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "Loi",
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		Message:    "Update user successfully",
		Data:       user,
	})
}
