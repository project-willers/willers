package router

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"willers-api/auth"
	"willers-api/handler"
)

func Init() *echo.Echo {
	// create instance
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodPost},
	}))

	// create routing
	e.POST("/login", handler.Login)
	e.POST("/register", handler.Register)

	// affect jwt middleware routing
	api := e.Group("/api")
	api.Use(middleware.JWTWithConfig(auth.Jwtconfig()))

	// TODO friends
	api.GET("/friend", handler.GetFriend)
	api.POST("/friend/request", handler.RequestFriend)
	api.POST("/friend/add", handler.AddFriend)
	api.POST("/friend/delete", handler.DeleteFriend)

	// TODO diary
	api.GET("/diary/read", handler.DiaryRead)
	api.POST("/diary/write", handler.DiaryWrite)
	api.POST("/diary/edit", handler.DiaryEdit)
	api.POST("/diary/delete", handler.DiaryDelete)

	// TODO comment
	api.GET("/comment/read", handler.CommentRead)
	api.POST("/comment/write", handler.CommentWrite)
	api.POST("/comment/edit", handler.CommentEdit)
	api.POST("/comment/delete", handler.CommentDelete)

	return e
}
