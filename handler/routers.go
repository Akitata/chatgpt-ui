package handler

import (
	"embed"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"net/http"
)

func InitRouter(app *fiber.App, publicFS embed.FS) {

	app.Get("/chat", Chat)

	app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(publicFS),
		Browse:     true,
		PathPrefix: "/public",
	}))
}
