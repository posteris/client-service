package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/posteris/client-service/models"
	"github.com/posteris/client-service/services"
	"github.com/posteris/commons/errors"
	"github.com/posteris/commons/validation"
)

// List Clientes
// @Summary      List Clients
// @Description  List all clients
// @Tags         Client
// @Accept       json
// @Produce      json
// @Param name 	  query string false "Client Name"
// @Param surname query string false "Client Name"
// @Param email   query string false "Client email" Format(email)
// @Param active  query bool   false "is cliente active?"
// @Success      200 {object} models.Client
// @Failure 	 500 {object} errors.DefaultError
// @Router       /api/v1/clientes [get]
func ListAllClients(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).JSON("{\"message\": \"Not Implemented Yet\"}")
}

// Get client By ID
// @Summary      Get client by ID
// @Description  Obtains the client object based in their ID
// @Tags         Client
// @Accept       json
// @Produce      json
// @Param id path string true "Client ID"
// @Success      200 {object} models.Client
// @Failure      404 {object} errors.DefaultError
// @Failure 	 500 {object} errors.DefaultError
// @Router       /api/v1/clientes/{id} [get]
func FindClientById(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(
			errors.CreateDefaultError("Client ID is required"),
		)
	}

	client := services.GetById(id)

	if client == nil {
		return c.Status(fiber.StatusNotFound).JSON(
			errors.CreateDefaultError("Client not found"),
		)
	}

	return c.Status(fiber.StatusOK).JSON(client)
}

// CreateClient controller to create new client
// @Summary      Create new Client
// @Description  Create a new client and persist it at the database.
// @Tags         Client
// @Accept       json
// @Produce      json
// @Param body body models.Client true "Body"
// @Success      201 {object} models.Client
// @Failure 	 415 {object} errors.DefaultError
// @Failure 	 400 {object} []errors.ValidationError
// @Failure 	 500 {object} errors.DefaultError
// @Router       /api/v1/clientes [post]
func CreateClient(c *fiber.Ctx) error {
	client := new(models.Client)

	if err := c.BodyParser(client); err != nil {
		return c.Status(fiber.StatusUnsupportedMediaType).JSON(
			errors.CreateDefaultError(err.Error()),
		)
	}

	if err := validation.ValidateModel(client); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	if err := services.Create(client); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			errors.CreateDefaultError(err.Error()),
		)
	}

	return c.Status(fiber.StatusCreated).JSON(client)
}

func UpdateClientById(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).JSON("{\"message\": \"Not Implemented Yet\"}")
}

func DeleteClientById(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).JSON("{\"message\": \"Not Implemented Yet\"}")
}
