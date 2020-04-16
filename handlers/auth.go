package handlers

import (
	"database/sql"
	"net/http"
	"jwt-crud/models"
	"jwt-crud/forms"
	"jwt-crud/helpers"
	"github.com/labstack/echo"
	"github.com/dgrijalva/jwt-go"
	
)
type M map[string]interface{}

func Login(db *sql.DB) echo.HandlerFunc  {
	return func (c echo.Context) error {
		form := new(forms.Login)
		c.Bind(form)
		if err := c.Validate(form); err != nil {
			return err 
		}

		//user exist check
		user := models.ShowUserByEmail(db, form.Email)
		if user.Name == "" {
			data := M{"message": "User Not Found"}
			return c.JSON(404, data)
		}

		//password check
		passwordCheck := helpers.CheckPasswordHash(form.Password, user.Password)
		if passwordCheck == false {
			data := M{"message": "Password doesn't match"}
			return c.JSON(400, data)
		}

		//generate jwt token
		token, err := helpers.GenerateToken(user.ID, user.Name, user.Email)
		if err != nil {
			panic(err)
		}

		return c.JSON(http.StatusOK, map[string]string{
			"token": token,
		})
	}
}

func UserAuth() echo.HandlerFunc {
	return func (c echo.Context) error {
		user := c.Get("user").(*jwt.Token)

		return c.JSON(http.StatusOK, user.Claims)
	}
}