package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/posteris/client-service/models"
	"github.com/posteris/client-service/services"
	"github.com/posteris/commons/errors"
	"github.com/posteris/commons/validation"
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
	return c.Status(fiber.StatusNotImplemented).JSON("{\"message\": \"Not Implemented Yet\"}")
}

func FindClientById(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).JSON("{\"message\": \"Not Implemented Yet\"}")
}

// CreateClient controller to create new client
// @Summary      Create new Client
// @Description  Create a new client and persist it at the database.
// @Tags         Client
// @Accept       json
// @Produce      json
// @Success      201 {object} models.Client
// @Failure 	 415 {object} errors.DefaultError
// @Failure 	 400 {object} []errors.ValidationError
// @Failure 	 500 {object} errors.DefaultError
// @Router       /api/v1/clientes [post]
func CreateClient(cli *fiber.Ctx) error {
	client := new(models.Client)

	fiber.SetParserDecoder(fiber.ParserConfig{
		IgnoreUnknownKeys: false,
		ZeroEmpty:         true,
	})

	if err := cli.BodyParser(client); err != nil {
		return cli.Status(fiber.StatusUnsupportedMediaType).JSON(
			errors.CreateDefaultError(err.Error()),
		)
	}

	if err := validation.ValidateModel(client); err != nil {
		return cli.Status(fiber.StatusBadRequest).JSON(err)
	}

	if err := services.Create(client); err != nil {
		return cli.Status(fiber.StatusInternalServerError).JSON(
			errors.CreateDefaultError(err.Error()),
		)
	}

	return cli.Status(fiber.StatusCreated).JSON(client)
}

func UpdateClientById(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).JSON("{\"message\": \"Not Implemented Yet\"}")
}

func DeleteClientById(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).JSON("{\"message\": \"Not Implemented Yet\"}")
}
