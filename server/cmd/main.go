package main

import (
	"log"
	"server/db"
	"server/internal/chatbot"
	"server/internal/user"
	"server/internal/ws"
	"server/router"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("could not initialize database connection: %s", err)
	}

	userRepository := user.NewRepository(dbConn.GetDB())
	userService := user.NewService(userRepository)
	userHandler := user.NewHandler(userService)

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	chatbotHandler := chatbot.NewHandler()
	go hub.Run()

	router.InitRouter(userHandler, wsHandler, chatbotHandler)
	router.Start("0.0.0.0:8080")
}
