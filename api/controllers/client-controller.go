package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/posteris/client-service/models"
	"github.com/posteris/client-service/services"
	validate "github.com/posteris/custom-validate"
)

func ListAllClients(c *fiber.Ctx) error {
	c.Send([]byte("List All Users"))

	return nil
}

func FindClientById(c *fiber.Ctx) error {
	c.Send([]byte("List User by Id"))

	return nil
}

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
