package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"

	"willers-api/model"
)

func GetFriend(c echo.Context) error {
	req := new(model.Friend)
	if err := c.Bind(req); err != nil {
		return echo.ErrBadRequest
	}
	if err := validate.Struct(req); err != nil {
		return echo.ErrBadRequest
	}
	// debug
	fmt.Fprintln(os.Stdout, req)

	friends, err := model.FindFriend(req)
	if err != nil {
		return err
	}

	json, err := json.Marshal(friends)
	return c.JSON(http.StatusOK, json)
}

func GetFriendRequest(c echo.Context) error {
	req := new(model.Friend)
	if err := c.Bind(req); err != nil {
		return echo.ErrBadRequest
	}
	if err := validate.Struct(req); err != nil {
		return echo.ErrBadRequest
	}
	// debug
	fmt.Fprintln(os.Stdout, req)

	friends, err := model.FindFriendRequest(req)
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
