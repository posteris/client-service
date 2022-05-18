package services

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/posteris/client-service/models"
)

func createNewClientToUpdate() *models.Client {
	clientToUpdate := &models.Client{
		Name:    models.GetStringPointer("john"),
		Surname: models.GetStringPointer("smith"),
		Email: models.GetStringPointer(
			fmt.Sprintf("%s@test.go", uuid.NewString()),
		),
	}

	Create(clientToUpdate)

	return clientToUpdate
}

func createClient(id, name, surname, email string, active bool) *models.Client {
	client := &models.Client{
		ID:      id,
		Name:    &name,
		Surname: &surname,
		Email:   &email,
	}

	if *client.Name == "" {
		client.Name = nil
	}

	if *client.Surname == "" {
		client.Surname = nil
	}

	if *client.Email == "" {
		client.Email = nil
	}

	return client
}
