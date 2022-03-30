package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/posteris/client-service/api/router"
)

func Run(serverPort string) {
	app := fiber.New()

	app.Get("/dashboard", monitor.New())

	//register routes
	router.SetupRoutes(app)

	app.Listen(serverPort)
}
