package services

import (
	"github.com/posteris/client-service/db"
	"github.com/posteris/client-service/models"
)

func Create(client *models.Client) error {
	dbInstance := db.GetInstance()

	return dbInstance.Create(client).Error
}
