package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"

	echo "github.com/labstack/echo/v4"

	"willers-api/auth"
	"willers-api/model"
)

func Register(c echo.Context) error {
	// JSON request
	user := new(model.Account)
	if err := c.Bind(user); err != nil {
		return echo.ErrBadRequest
	}
	if err := validate.Struct(user); err != nil {
		return echo.ErrBadRequest
	}
	// debug
	fmt.Fprintln(os.Stdout, user)
	user.Password = auth.HashStr(user.Password)

	// new User Create
	u, err := model.CreateUser(user)
	if err != nil {
		log.Println(err)
		return err
	}
	// debug
	fmt.Fprintln(os.Stdout, user)
	// token create
	token, err := auth.CreateToken(u)
	if err != nil && token == "" {
		log.Println(err)
		return err
		// return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, auth.Token{Jwt: token})
}
