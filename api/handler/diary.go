package handler

import "github.com/labstack/echo/v4"

type Diary interface {
	Write()
	Read()
	Edit()
	Delete()
}

func DiaryWrite(c echo.Context) error {
	return nil
}

func DiaryRead(c echo.Context) error {
	return nil
}

func DiaryEdit(c echo.Context) error {
	return nil
}

func DiaryDelete(c echo.Context) error {
	return nil
}
