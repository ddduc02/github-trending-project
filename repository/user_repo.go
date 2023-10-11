package repository

import (
	"context"

	"github.com/ddduc02/gh-trending/models"
	"github.com/ddduc02/gh-trending/models/request"
)

type UserRepo interface {
	CheckLogin(context context.Context, loginReq request.RequestSignIp) (models.User, error)
	SaveUser(context context.Context, user models.User) (models.User, error)
	SelectUserById(context context.Context, userId string) (models.User, error)
	UpdateUser(context context.Context, user models.User) (models.User, error)
}
