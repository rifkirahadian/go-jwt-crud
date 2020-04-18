package helpers

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateToken(id uint, name string, email string) (string, error) {
	//create token
	token := jwt.New(jwt.SigningMethodHS256)

	//set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["name"] = name
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	//generate encode token
	t, err := token.SignedString([]byte("secret"))
	
	return string(t),err
}