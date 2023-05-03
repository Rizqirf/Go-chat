package router

import (
	"net/http"

	"github.com/Rizqirf/go-chat/server/internal/user"
	"github.com/Rizqirf/go-chat/server/internal/ws"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var e = echo.New()

// func allowOrigin(origin string) (bool,error) {
// 	regexp.MatchString(`^http:\/\/localhost\:3000`, origin)
// 	// return origin == "http://localhost:3000"
// }

func InitRouter(userHandler *user.Handler, wsHandler *ws.Handler) {
	

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPost},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		// AllowOriginFunc: allowOrigin,
		AllowCredentials: true,
		MaxAge: 24*60*60,
	}))
	

	e.POST("/signup", userHandler.CreateUser)
	e.POST("/login", userHandler.Login)
	e.GET("/logout", userHandler.Logout)

	e.POST("/ws/createRoom", wsHandler.CreateRoom)
	e.GET("/ws/joinRoom/:roomId", wsHandler.JoinRoom)
	e.GET("/ws/getRooms", wsHandler.GetRooms)
	e.GET("/ws/getClients/:roomId", wsHandler.GetClients)
}

func Start(addr string) error {
	return e.Start(addr)
}

