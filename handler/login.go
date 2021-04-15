package handler

import (
	"echo/mdw"
	"echo/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func Login(c echo.Context)error  {

	username := c.Get("name").(string)
	//password := c.Get("password").(string)
	role :=c.Get("role").(bool)

	//create token

	token := jwt.New(jwt.SigningMethodHS256)

	//set claim
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = username
	claims["role"] = role
	claims["exp"] = time.Now().Add(4*time.Minute).Unix()

	//gen token

	t, err := token.SignedString(mdw.GetSecretKey())

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK,&models.LoginResponse{
		Token: t,
	})

}

