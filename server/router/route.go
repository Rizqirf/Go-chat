package router

import (
	"net/http"

	"github.com/Rizqirf/go-chat/server/internal/user"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var e = echo.New()

func InitRouter(userHandler *user.Handler) {
	

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPost},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
		MaxAge: 24*60*60,
	}))
	

	e.POST("/signup", userHandler.CreateUser)
	e.POST("/login", userHandler.Login)
	e.GET("/logout", userHandler.Logout)

	// e.POST("/ws/createRoom", wsHandler.CreateRoom)
	// e.GET("/ws/joinRoom/:roomId", wsHandler.JoinRoom)
	// e.GET("/ws/getRooms", wsHandler.GetRooms)
	// e.GET("/ws/getClients/:roomId", wsHandler.GetClients)
}

func Start(addr string) error {
	return e.Start(addr)
}

