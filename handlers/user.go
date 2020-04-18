package handlers

import (
	"net/http"
	"jwt-crud/models"
	"jwt-crud/helpers"
	"jwt-crud/configs"
	"github.com/labstack/echo"
)

type H map[string]interface{}

func Register() echo.HandlerFunc  {
	return func (c echo.Context) error {
		u := new(models.User)
		c.Bind(u)
		if err := c.Validate(u); err != nil {
			return err 
		}

		hash, _ := helpers.HashPassword(u.Password)
		u.Password = hash

		db := configs.InitGormDB()
		if err := db.Create(&u).Error; err != nil {
			return c.JSON(400, H{
				"message": "Email is already in use",
			})
		}

		return c.JSON(http.StatusOK, H{
			"message": "Register success",
		})
	}
}