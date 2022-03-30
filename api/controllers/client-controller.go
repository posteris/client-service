package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/posteris/client-service/models"
	"github.com/posteris/client-service/services"
	validate "github.com/posteris/custom-validate"
)

// ListAllClients godoc
// @Summary      List all clients
// @Description  Perform a paginated search through the client repository geting all active users
// @Tags         Client
// @Accept       json
// @Produce      json
// @Success      200  {object}  []models.Client
// @Failure 400 {object} fiber.Map
// @Failure 404 {object} fiber.Map
// @Failure 500 {object} fiber.Map
// @Router       /api/v1/clientes/{id} [get]
func ListAllClients(c *fiber.Ctx) error {
	c.Send([]byte("List All Users"))

	return nil
}

func FindClientById(c *fiber.Ctx) error {
	c.Send([]byte("List User by Id"))

	return nil
}

// CreateClient controller to create new client
// @Summary      Create new Client
// @Description  Create a new client and persist it at the database. This action can be done by sync and async way. By default, the sync is selected, but, when the parameter async is set as true, the system will assume that the implementation will perform the request asynchronously, and the requester should wait by the response at a callback URL.
// @Tags         Client
// @Accept       json
// @Produce      json
// @Success      201 {object}  models.Client
// @Failure 	 415 {object} fiber.Map the content is not a valid JSON
// @Failure 	 400 {object} []validate.CustomError
// @Failure 	 500 {object} fiber.Map error when try to save object
// @Router       /api/v1/clientes [post]
func CreateClient(c *fiber.Ctx) error {
	client := new(models.Client)

	if err := c.BodyParser(client); err != nil {
		return c.Status(fiber.StatusUnsupportedMediaType).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := validate.Run(client); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	if err := services.CreateClient(client); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// c.Send([]byte("Create a new User"))
	return c.Status(fiber.StatusCreated).JSON(client)
}

func UpdateClientById(c *fiber.Ctx) error {
	c.Send([]byte("Update User by Id"))

	return nil
}

func DeleteClientById(c *fiber.Ctx) error {
	c.Send([]byte("Delete User by Id"))

	return nil
}
