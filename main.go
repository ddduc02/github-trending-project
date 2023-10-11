package main

import (
	"github.com/ddduc02/gh-trending/db"
	"github.com/ddduc02/gh-trending/handler"
	repoimplement "github.com/ddduc02/gh-trending/repository/repo_implement"
	"github.com/ddduc02/gh-trending/router"
	"github.com/labstack/echo/v4"
)

func main() {
	sql := &db.Sql{
		Host:     "localhost",
		Port:     5432,
		UserName: "postgres",
		Password: "abc123",
		DbName:   "github-trending",
	}

	sql.Connect()
	defer sql.Close()

	e := echo.New()
	userHandler := handler.UserHandler{
		UserRepo: repoimplement.NewUserRepo(sql),
	}
	api := router.API{
		Echo:        e,
		UserHandler: userHandler,
	}
	api.SetupRouter()
	e.Logger.Fatal(e.Start(":8080"))
}
