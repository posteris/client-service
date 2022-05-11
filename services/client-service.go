package services

import (
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

func List(search map[string]interface{}) ([]models.Client, error) {
	var clientList []models.Client

	dbInstance := db.GetInstance()

	err := dbInstance.Where(search).Find(&clientList).Error

	return clientList, err
}

func Create(client *models.Client) error {
	dbInstance := db.GetInstance()

	return dbInstance.Create(client).Error
}

func Update(client *models.Client) error {
	dbInstance := db.GetInstance()

	return dbInstance.Model(client).Omit("id").Updates(client).Error
}
