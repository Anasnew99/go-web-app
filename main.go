package main

import (
	"anasnew99/server/chat_app/server"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	server.StartServer()
}
