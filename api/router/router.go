package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/posteris/client-service/api/router/routes"
)

func SetupRoutes(app *fiber.App) {
	router := app.Group("/api/v1", logger.New())

	routes.ClientRoutes(router)
}
