package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"willers-api/model"

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
	if err := model.AddDiary(diary); err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, "")
}

func DiaryRead(c echo.Context) error {
	// TODO
	// jwt tokenからusernameを取得
	userName := ""
	// debug
	fmt.Fprintln(os.Stdout, userName)

	diaries, err := model.GetDiaries(userName)
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
	if err := model.DeleteDiary(diary); err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, "")
}
