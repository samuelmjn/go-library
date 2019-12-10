package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/samuelmjn/go-library/config"
	"github.com/samuelmjn/go-library/db"
	"github.com/samuelmjn/go-library/httpsvc"
	"github.com/samuelmjn/go-library/repository"
)

func main() {
	config.InitConfig()
	e := echo.New()

	//Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS Handler
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	db, err := db.InitializeDatabase()
	if err != nil {
		panic(err)
	}
	bookRepo := repository.NewBookRepository(db)
	userRepo := repository.NewUserRepository(db)
	service := httpsvc.NewService()
	service.RegisterEcho(e)
	service.RegisterBookRepository(bookRepo)
	service.RegisterUserRepository(userRepo)

	service.InitializeRoutes()

	e.Logger.Fatal(e.Start(":" + config.Port()))
}
