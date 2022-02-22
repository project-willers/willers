package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"

	"willers-api/model"
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
	if name != diary.UserName {
		return c.JSON(http.StatusNotAcceptable, "")
	}

	if err := model.AddDiary(diary); err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, "{}")
}

func DiaryRead(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)

	diaries, err := model.GetDiaries(name)
	if err != nil {
		return err
	}

	d := model.D{
		Diaries: diaries,
	}

	return c.JSON(http.StatusOK, d)
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
	if name != diary.UserName {
		return c.JSON(http.StatusNotAcceptable, "")
	}

	if err := model.UpdateDiary(diary); err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, "{}")
}

// まだ実装できてないです。
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
	if name != diary.UserName {
		return c.JSON(http.StatusNotAcceptable, "")
	}

	if err := model.DeleteDiary(diary); err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, "{}")
}
