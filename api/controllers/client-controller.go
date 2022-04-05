package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/posteris/client-service/models"
	"github.com/posteris/client-service/services"
	"github.com/posteris/commons/errors"
	"github.com/posteris/commons/parameters"
	"github.com/posteris/commons/parser"
)

// ListAllClients godoc
// @Summary      List all clients
// @Description  Perform a paginated search through the client repository geting all active users
// @Tags         Client
// @Accept       json
// @Produce      json
// @Success      200  {object}  []models.Client
// @Failure 400 {object} errors.DefaultError
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
// @Description  Create a new client and persist it at the database. This action can be done by sync and async way. By default, the sync is selected, but, when the parameter async is set as true, the system will assume that the implementation will perform the request asynchronously, and the requester should wait by the response at the provided callback.
// @Tags         Client
// @Accept       json
// @Produce      json
// @Success      201 {object} models.Client
// @Failure 	 415 {object} errors.DefaultError
// @Failure 	 400 {object} []errors.ValidationError
// @Failure 	 500 {object} errors.DefaultError
// @Router       /api/v1/clientes [post]
func CreateClient(c *fiber.Ctx) error {
	client := new(models.Client)

	if err := parser.BodyParser(client, c); err != nil {

	}

	// if err := c.BodyParser(client); err != nil {
	// 	return c.Status(fiber.StatusUnsupportedMediaType).JSON(
	// 		errors.CreateDefaultError(err.Error()),
	// 	)
	// }

	// if err := validate.ValidateModel(client); err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(err)
	// }

	if parameters.IsAsyncRequest(c) {

		return c.Status(fiber.StatusCreated).JSON(client)
	}

	if err := services.CreateClient(client); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			errors.CreateDefaultError(err.Error()),
		)
	}

	return c.Status(fiber.StatusCreated).JSON(client)
}

func UpdateClientById(c *fiber.Ctx) error {
	c.Send([]byte("Update User by Id"))
	1
	return nil
}

func DeleteClientById(c *fiber.Ctx) error {
	c.Send([]byte("Delete User by Id"))

	return nil
}
