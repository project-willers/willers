package handler

import (
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	echo "github.com/labstack/echo/v4"

	"willers-api/auth"
	"willers-api/model"
)

func Login(c echo.Context) error {
	// JSON request
	u := new(model.LoginInfo)
	if err := c.Bind(u); err != nil {
		return err
	}
	if err := validator.New().Struct(u); err != nil {
		return err
	}

	user := &model.Account{}
	user, err := model.FindUser(&model.LoginInfo{Name: u.Name})
	if err != nil {
		return err
	}

	if !((u.Name == user.Name || u.Name == user.Email) && u.Password == user.Password) {
		return &echo.HTTPError{
			Code:    http.StatusUnauthorized,
			Message: "Invalid name or password",
		}
	}

	token, err := auth.CreateToken(user)
	if err != nil && token == "" {
		return echo.ErrInternalServerError
	}

	log.Println("Login Success")
	return c.JSON(http.StatusOK, auth.Token{Jwt: token})
}
