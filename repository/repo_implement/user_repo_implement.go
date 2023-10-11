package repoimplement

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/ddduc02/gh-trending/db"
	"github.com/ddduc02/gh-trending/models"
	"github.com/ddduc02/gh-trending/models/request"
	"github.com/ddduc02/gh-trending/repository"
	"github.com/labstack/gommon/log"
)

type UserRepoImplement struct {
	sql *db.Sql
}

func NewUserRepo(sql *db.Sql) repository.UserRepo {
	return &UserRepoImplement{sql: sql}
}

func (u *UserRepoImplement) CheckLogin(context context.Context, loginReq request.RequestSignIp) (models.User, error) {
	var user models.User

	err := u.sql.Db.GetContext(context, &user, "SELECT * FROM users WHERE email=$1", loginReq.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New("User is not found")
		}
		log.Error(err.Error())
		return user, err
	}
	return user, nil
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

func (u *UserRepoImplement) SelectUserById(context context.Context, userId string) (models.User, error) {
	var user models.User
	err := u.sql.Db.GetContext(context, &user, "SELECT * FROM users WHERE user_id=$1", userId)
	if err != nil {
		log.Error(err.Error())
		return user, errors.New("Select user Failed")
	}
	return user, nil
}

func (u *UserRepoImplement) UpdateUser(context context.Context, user models.User) (models.User, error) {
	sqlStatement := `UPDATE users SET full_name = :full_name,
										email = :email,
										updated_at = COALESCE (:updated_at, updated_at)
							WHERE user_id = :user_id
	
	`
	user.UpdatedAt = time.Now()
	_, err := u.sql.Db.NamedExecContext(context, sqlStatement, user)
	if err != nil {
		log.Error(err)
		return user, err
	}
	return user, nil
}
