package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"willers-api/model"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func DiaryWrite(c echo.Context) error {
	diary := new(model.Diary)
	if err := c.Bind(diary); err != nil {
		return echo.ErrBadRequest
	}
	if err := validate.Struct(diary); err != nil {
		return echo.ErrBadRequest
	}
	// debug
	fmt.Fprintln(os.Stdout, diary)

	// Authorize
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	if diary.UserName != name {
		return echo.ErrBadRequest
	}

	if err := model.AddDiary(diary); err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, "")
}

func DiaryRead(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	// debug
	fmt.Fprintln(os.Stdout, name)

	diaries, err := model.GetDiaries(name)
	if err != nil {
		return err
	}

	json, err := json.Marshal(diaries)
	return c.JSON(http.StatusOK, json)
}

func DiaryEdit(c echo.Context) error {
	diary := new(model.Diary)
	if err := c.Bind(diary); err != nil {
		return echo.ErrBadRequest
	}
	if err := validate.Struct(diary); err != nil {
		return echo.ErrBadRequest
	}
	// debug
	fmt.Fprintln(os.Stdout, diary)

	// Authorize
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	if diary.UserName != name {
		return echo.ErrBadRequest
	}

	if err := model.UpdateDiary(diary); err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, "")
}

func DiaryDelete(c echo.Context) error {
	diary := new(model.Diary)
	if err := c.Bind(diary); err != nil {
		return echo.ErrBadRequest
	}
	if err := validate.Struct(diary); err != nil {
		return echo.ErrBadRequest
	}
	// debug
	fmt.Fprintln(os.Stdout, diary)

	// Authorize
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	if diary.UserName != name {
		return echo.ErrBadRequest
	}

	if err := model.DeleteDiary(diary); err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, "")
}
