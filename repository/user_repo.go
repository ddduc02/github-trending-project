package repository

import (
	"context"

	"github.com/ddduc02/gh-trending/models"
)

type UserRepo interface {
	SaveUser(context context.Context, user models.User) (models.User, error)
}
