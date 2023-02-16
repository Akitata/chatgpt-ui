package main

import (
	"embed"
	"flag"
	"github.com/akitata/chatgpt-ui/chat"
	"github.com/akitata/chatgpt-ui/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	"log"
)

var (
	addr  = flag.String("adrr", "0.0.0.0:3000", "ip and port to listen")
	token = flag.String("token", "sk-xxx", "chatgpt api token.")
	prod  = flag.Bool("prod", false, "Enable prefork in Production")
)

//go:embed  public/*
var publicFS embed.FS

func main() {
	flag.Parse()

	chat.InitChatClient(*token)

	app := fiber.New(fiber.Config{
		Prefork: *prod,
	})
	app.Use(recover2.New(recover2.Config{
		EnableStackTrace: true,
	}))
	app.Use(logger.New())
	app.Use(handler.SessionHandler)

	handler.InitRouter(app, publicFS)

	log.Fatalln(app.Listen(*addr))
}
