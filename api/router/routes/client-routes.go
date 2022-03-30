package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/posteris/client-service/api/controllers"
)

//ClientRoutes function to create the client routes mapping the client controllers
func ClientRoutes(router fiber.Router) {
	client := router.Group("/clients")

	client.Get("/", controllers.ListAllClients)

	client.Get("/:id", controllers.FindClientById)

	client.Post("/", controllers.CreateClient)

	client.Put("/:id", controllers.UpdateClientById)

	client.Delete("/:id", controllers.DeleteClientById)
}
