package main

import (
	"log"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/posteris/client-service/api/middleware"
	"github.com/posteris/client-service/api/router"
	"github.com/posteris/client-service/database"
	_ "github.com/posteris/client-service/docs"
)

// @title Client Registration service
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	database.InitDatabase()

	app := fiber.New()

	app.Get("/dashboard", monitor.New())
	app.Get("/swagger/*", swagger.HandlerDefault) // default

	middleware.FiberMiddleware(app)

	//register routes
	router.SetupRoutes(app)

	err := app.Listen(":8000")
	if err != nil {
		log.Fatalf("Startup error: %v", err)
	}
}
