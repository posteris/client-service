package main

import (
	"log"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/posteris/client-service/api/middleware"
	"github.com/posteris/client-service/api/router"
	"github.com/posteris/client-service/db"
	_ "github.com/posteris/client-service/docs"
)

// @title Client Registration service
// @version 1.0
// @description Client registration service that enable to manage clients, their addresses, their contacts and their documents. In this service is also included the client registration
// @termsOfService http://posteris.io/terms/
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8000
// @BasePath /
func main() {
	db.InitDatabase()
	db.Automigrate()

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
