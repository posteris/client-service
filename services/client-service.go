package services

import (
	"errors"

	"github.com/posteris/client-service/db"
	"github.com/posteris/client-service/models"
)

func GetById(id string) *models.Client {
	dbInstance := db.GetInstance()

	var client models.Client
	result := dbInstance.First(&client, "id = ?", id)

	if result.RowsAffected == 0 {
		return nil
	}

	return &client
}

func List(search map[string]interface{}) []models.Client {
	var clientList []models.Client

	dbInstance := db.GetInstance()

	dbInstance.Where(search).Find(&clientList)

	return clientList
}

func Create(client *models.Client) error {
	if client.ID != "" {
		return errors.New("you cannot send the Client ID in the creation phase")
	}

	if client.Active {
		return errors.New("you cannot create the Client with active status. Create it, then activate")
	}

	dbInstance := db.GetInstance()

	return dbInstance.Create(client).Error
}

func Update(client *models.Client) error {
	dbInstance := db.GetInstance()

	return dbInstance.Model(client).Omit("id").Omit("active").Updates(client).Error
}

func Delete(id string) error {
	dbInstance := db.GetInstance()

	return dbInstance.Delete(&models.Client{}, id).Error
}
