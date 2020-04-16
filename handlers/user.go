package handlers

import (
	"database/sql"
	"net/http"
	"jwt-crud/models"
	"jwt-crud/helpers"
	"github.com/labstack/echo"
)

type H map[string]interface{}

func Register(db *sql.DB) echo.HandlerFunc  {
	return func (c echo.Context) error {
		u := new(models.User)

		if err := c.Bind(u); err != nil {
			return err
		}

		if err := c.Validate(u); err != nil {
			return err 
		}

		hash, _ := helpers.HashPassword(u.Password)
		
		id, err := models.CreateUser(db, u.Name, u.Email, hash)
		if err == nil {
			return c.JSON(http.StatusOK, H{
				"created": id,
			})
		}else{
			return err
		}

		return c.String(http.StatusOK, u.Name)
	}
}