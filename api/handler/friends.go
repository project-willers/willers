package handler

import (
	"encoding/json"
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

	json, err := json.Marshal(friends)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(json)
	return c.JSON(http.StatusOK, json)
}

func GetFriendRequests(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	// debug
	fmt.Fprintln(os.Stdout, name)

	friends, err := model.FindFriendRequests(name)
	if err != nil {
		return err
	}

	json, err := json.Marshal(friends)
	return c.JSON(http.StatusOK, json)
}

func RequestFriend(c echo.Context) error {
	req := new(model.Friend)
	if err := c.Bind(req); err != nil {
		return echo.ErrBadRequest
	}
	if err := validate.Struct(req); err != nil {
		return echo.ErrBadRequest
	}
	// debug
	fmt.Fprintln(os.Stdout, req)
	if err := model.FriendRequest(req); err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, "")
}

func AddFriend(c echo.Context) error {
	req := new(model.FriendResponse)
	if err := c.Bind(req); err != nil {
		return echo.ErrBadRequest
	}
	if err := validate.Struct(req); err != nil {
		return echo.ErrBadRequest
	}
	// debug
	fmt.Fprintln(os.Stdout, req)

	if err := model.AddFriend(req); err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, "")
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
	if err := model.DeleteFriend(req); err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, "")
}
