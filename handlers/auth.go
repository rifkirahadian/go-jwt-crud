package handlers

import (
	"net/http"
	"jwt-crud/models"
	"jwt-crud/forms"
	"jwt-crud/helpers"
	"jwt-crud/configs"
	"github.com/labstack/echo"
	"github.com/dgrijalva/jwt-go"
	
)
type M map[string]interface{}

func Login() echo.HandlerFunc  {
	return func (c echo.Context) error {
		form := new(forms.Login)
		c.Bind(form)
		if err := c.Validate(form); err != nil {
			return err 
		}

		//user exist check
		db := configs.InitGormDB()
		var user models.User
		userQuery := db.First(&user, "email=?", form.Email)
		if userQuery.RowsAffected == 0 {
			return c.JSON(404, M{"message": "User Not Found"})
		}

		//password check
		passwordCheck := helpers.CheckPasswordHash(form.Password, user.Password)
		if passwordCheck == false {
			return c.JSON(400, M{"message": "Password doesn't match"})
		}

		//generate jwt token
		token, err := helpers.GenerateToken(user.ID, user.Name, user.Email)
		if err != nil {
			panic(err)
		}

		return c.JSON(http.StatusOK, M{
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