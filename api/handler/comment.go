package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"willers-api/model"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func CommentWrite(c echo.Context) error {
	comment := new(model.Comment)
	if err := c.Bind(comment); err != nil {
		return echo.ErrBadRequest
	}
	if err := validate.Struct(comment); err != nil {
		return echo.ErrBadRequest
	}
	// debug
	fmt.Fprintln(os.Stdout, comment)

	// Authorize
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	if comment.CommentUser != name {
		log.Println("You are not user!")
		return echo.ErrBadRequest
	}

	if err := model.AddComment(comment); err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, "")
}

func CommentRead(c echo.Context) error {
	diary := new(model.Diary)
	if err := c.Bind(diary); err != nil {
		return echo.ErrBadRequest
	}
	if err := validate.Struct(diary); err != nil {
		return echo.ErrBadRequest
	}
	// debug
	fmt.Fprintln(os.Stdout, diary)

	comments, err := model.GetComments(diary)
	if err != nil {
		return err
	}

	json, err := json.Marshal(comments)
	return c.JSON(http.StatusOK, json)
}

func CommentEdit(c echo.Context) error {
	comment := new(model.Comment)
	if err := c.Bind(comment); err != nil {
		return echo.ErrBadRequest
	}
	if err := validate.Struct(comment); err != nil {
		return echo.ErrBadRequest
	}
	// debug
	fmt.Fprintln(os.Stdout, comment)

	// Authorize
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	if comment.CommentUser != name {
		log.Println("You are not user!")
		return echo.ErrBadRequest
	}

	if err := model.UpdateComment(comment); err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, "")
}

func CommentDelete(c echo.Context) error {
	comment := new(model.Comment)
	if err := c.Bind(comment); err != nil {
		return echo.ErrBadRequest
	}
	if err := validate.Struct(comment); err != nil {
		return echo.ErrBadRequest
	}
	// debug
	fmt.Fprintln(os.Stdout, comment)

	// Authorize
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	if comment.CommentUser != name {
		log.Println("You are not user!")
		return echo.ErrBadRequest
	}

	if err := model.DeleteComment(comment); err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, "")
}
