package repoimplement

import (
	"context"
	"errors"
	"time"

	"github.com/ddduc02/gh-trending/db"
	"github.com/ddduc02/gh-trending/models"
	"github.com/ddduc02/gh-trending/repository"
	"github.com/labstack/gommon/log"
)

type UserRepoImplement struct {
	sql *db.Sql
}

func NewUserRepo(sql *db.Sql) repository.UserRepo {
	return &UserRepoImplement{sql: sql}
}

func (u UserRepoImplement) SaveUser(context context.Context, user models.User) (models.User, error) {
	statement := `INSERT INTO users(user_id, full_name, email, password, role, created_at, updated_at)
	VALUES(:user_id, :full_name, :email, :password, :role, :created_at, :updated_at)
	`
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	_, err := u.sql.Db.NamedExecContext(context, statement, user)
	if err != nil {
		log.Error(err.Error())
		return user, errors.New("Sign Up Failed")
	}
	return user, nil
}
