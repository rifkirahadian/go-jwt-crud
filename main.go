package main

import (
	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
	"jwt-crud/handlers"
	"jwt-crud/configs"
	"jwt-crud/helpers"
	"jwt-crud/models"
	"gopkg.in/go-playground/validator.v9"
	"github.com/labstack/echo/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main()  {
	e := echo.New()

	db := configs.InitGormDB()
	db.AutoMigrate(&models.Task{})
	db.AutoMigrate(&models.User{})

	e.Validator = &CustomValidator{validator: validator.New()}
	e.HTTPErrorHandler = helpers.ValidationResponse

	//no auth routes
	e.POST("/register", handlers.Register())
	// e.POST("/login", handlers.Login(db))
	
	//auth middleware
	r := e.Group("/auth")
	r.Use(middleware.JWT([]byte("secret")))

	//auth routes
	r.GET("/user", handlers.UserAuth())

	e.Logger.Fatal(e.Start(":8000"))
}