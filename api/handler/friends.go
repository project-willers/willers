package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"

	"willers-api/model"
)

func GetFriends(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	// debug
	fmt.Fprintln(os.Stdout, name)

	friends, err := model.FindFriends(name)
	if err != nil {
		log.Println(err)
		return err
	}

	f := model.F{
		Friends: friends,
	}

	return c.JSON(http.StatusOK, f)
}

func GetMyFriendRequests(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	// debug
	fmt.Fprintln(os.Stdout, name)

	friends, err := model.GetMyFriendRequests(name)
	if err != nil {
		return err
	}

	f := model.F{
		Friends: friends,
	}

	return c.JSON(http.StatusOK, f)
}

func GetOtherFriendRequests(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	// debug
	fmt.Fprintln(os.Stdout, name)

	friends, err := model.GetOtherFriendRequests(name)
	if err != nil {
		return err
	}

	f := model.F{
		Friends: friends,
	}

	return c.JSON(http.StatusOK, f)
}

func RequestFriend(c echo.Context) error {
	req := new(model.Friend)
	if err := c.Bind(req); err != nil {
		log.Println(err)
		return echo.ErrBadRequest
	}
	if err := validate.Struct(req); err != nil {
		log.Println(err)
		return echo.ErrBadRequest
	}
	// debug
	fmt.Fprintln(os.Stdout, req)

	// Authorize
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	if req.MyName != name {
		return echo.ErrBadRequest
	}

	if err := model.FriendRequest(req); err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, "{}")
}

func AddFriend(c echo.Context) error {
	req := new(model.FriendResponse)
	if err := c.Bind(req); err != nil {
		log.Println(err)
		return echo.ErrBadRequest
	}
	if err := validate.Struct(req); err != nil {
		log.Println(err)
		return echo.ErrBadRequest
	}
	// debug
	fmt.Fprintln(os.Stdout, req)

	// Authorize
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	if req.MyName != name {
		return echo.ErrBadRequest
	}

	if err := model.AddFriend(req); err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, "{}")
}

func DeleteFriend(c echo.Context) error {
	req := new(model.Friend)
	if err := c.Bind(req); err != nil {
		return echo.ErrBadRequest
	}
	if err := validate.Struct(req); err != nil {
		return echo.ErrBadRequest
	}
	// debug
	fmt.Fprintln(os.Stdout, req)

	// Authorize
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	if req.MyName != name {
		return echo.ErrBadRequest
	}

	if err := model.DeleteFriend(req); err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, "{}")
}
