package main

import (
	"log"

	"github.com/Rizqirf/go-chat/server/db"
	"github.com/Rizqirf/go-chat/server/internal/user"
	"github.com/Rizqirf/go-chat/server/internal/ws"
	"github.com/Rizqirf/go-chat/server/router"
)

func main() {

	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("could not initialize database connection: %s", err)
	}

	userRep := user.NewRepo(dbConn.GetDB())
	userSvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSvc)

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()

	router.InitRouter(userHandler, wsHandler)
	router.Start("0.0.0.0:8080")

}