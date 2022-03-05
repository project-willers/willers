package router

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

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
	api.Use(middleware.JWT([]byte("SECRETKEY")))

	// TODO friends
	api.GET("/friend", handler.GetFriends)
	// api.GET("/friend/:id/request", handler.GetMyFriendRequests)
	api.GET("/friend/my/request", handler.GetMyFriendRequests)
	api.GET("/friend/other/request", handler.GetOtherFriendRequests)

	api.POST("/friend/request", handler.RequestFriend)
	api.POST("/friend", handler.AddFriend)
	api.DELETE("/friend", handler.DeleteFriend)

	// TODO diary
	api.GET("/diary", handler.DiaryRead)
	api.POST("/diary", handler.DiaryWrite)
	api.PATCH("/diary", handler.DiaryEdit)
	api.DELETE("/diary", handler.DiaryDelete)

	// TODO comment
	api.GET("/comment", handler.CommentRead)
	api.POST("/comment", handler.CommentWrite)
	api.PATCH("/comment", handler.CommentEdit)
	api.DELETE("/comment", handler.CommentDelete)

	return e
}
