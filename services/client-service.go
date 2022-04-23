package services

import (
	"github.com/posteris/client-service/database"
	"github.com/posteris/client-service/models"
)

func Create(client *models.Client, async bool) error {
	dbInstance := database.GetInstance()

	dbInstance.Save(client).Commit()

	return nil
}
